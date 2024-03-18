package impl

import (
	"github.com/edsrzf/mmap-go"
	"github.com/kukino/Amaru"
	"math"
	"os"
)

type anthologyImpl struct {
	aPath string
	aFile *os.File
	aMMap mmap.MMap
	iPath string
	iFile *os.File
	iMMap mmap.MMap

	lastKnownEoF            uint64
	defaultDossierCapacity  uint32
	defaultAnthologySizeMiB int64
	defaultIndexSizeTks     int64
}

func (a *anthologyImpl) Add(did Amaru.DocID, tid Amaru.TokenID) {
	offset := a.tidOffset(tid)
	var dossier Amaru.Dossier
	if offset == Amaru.InvalidOffset {
		dossier = a.newDossier(tid, a.defaultDossierCapacity)
	} else {
		dossier = a.GetDossier(tid)
	}

	if _, err := dossier.Add(did); err != nil {
		panic("could not add the did into this tid ...have you run out of initial space?")
	}
}

func (a *anthologyImpl) Load() error {
	var err error
	a.aFile, err = os.OpenFile(a.aPath, newFileFlags, newFilePerms)
	if err != nil {
		return err
	}
	if fi, err := a.aFile.Stat(); err == nil {
		if fi.Size() == 0 {
			a.aFile.Truncate(a.defaultAnthologySizeMiB * 1024 * 1024)
		}
	} else {
		return err
	}
	a.aMMap, err = mmap.Map(a.aFile, mmap.RDWR, 0)
	if err != nil {
		return err
	}

	a.iFile, err = os.OpenFile(a.iPath, newFileFlags, newFilePerms)
	if err != nil {
		return err
	}
	empty := false
	if fi, err := a.iFile.Stat(); err == nil {
		if fi.Size() == 0 {
			a.iFile.Truncate(a.defaultIndexSizeTks * 8 /*uint64*/)
			empty = true
		}
	} else {
		return err
	}
	a.iMMap, err = mmap.Map(a.iFile, mmap.RDWR, 0)
	if err != nil {
		return err
	}
	if empty {
		for i := 0; i < len(a.iMMap); i++ {
			a.iMMap[i] = 0xff
		}
	}
	return nil
}

func (a *anthologyImpl) Save() error {
	if err := a.iMMap.Flush(); err != nil {
		return err
	}
	if err := a.aMMap.Flush(); err != nil {
		return err
	}
	return nil
}

func (a *anthologyImpl) Exist() bool {
	return true
}

func (a *anthologyImpl) Close() error {
	if err := a.aMMap.Unmap(); err != nil {
		return err
	}
	if err := a.aFile.Close(); err != nil {
		return err
	}
	if err := a.iMMap.Unmap(); err != nil {
		return err
	}
	if err := a.iFile.Close(); err != nil {
		return err
	}
	return nil
}

func (a *anthologyImpl) Clear() {
	_ = a.Close()
	_ = os.Remove(a.aPath)
	_ = os.Remove(a.iPath)
	_ = a.Load()
}

func (a *anthologyImpl) Create() error {
	a.Clear()
	return nil
}

func (a *anthologyImpl) FindDocIDsWith(tids []Amaru.TokenID, limit int) []Amaru.DocID {

	var docids []Amaru.DocID
	var dossiers []Amaru.Dossier
	var pointers []uint32
	var counts []uint32

	tids = removeDuplicateTokenIDs(tids)

	if len(tids) == 0 {
		return docids
	}

	for _, tid := range tids {
		dossier := a.GetDossier(tid)
		dossiers = append(dossiers, dossier)
		pointers = append(pointers, 0)
		count := dossier.Count()
		counts = append(counts, count)
		if count == 0 {
			return docids
		}
	}

	for {
		// Assume match found; verify against first dossier's current DID.
		did := dossiers[0].Get(pointers[0])
		biggestDid := did
		match := true
		for i := 1; i < len(dossiers); i++ {
			currentDid := dossiers[i].Get(pointers[i])
			if currentDid != did { // If any dossier's current did don't match, not a match.
				match = false
			}
			if currentDid > biggestDid {
				biggestDid = currentDid
			}
		}
		if match {
			docids = append(docids, did)
			if len(docids) >= limit {
				return docids
			}
			for i := 0; i < len(dossiers); i++ {
				pointers[i]++
				if pointers[i] == counts[i] {
					return docids
				}
			}
		} else {
			for i := 0; i < len(dossiers); i++ {
				if dossiers[i].Get(pointers[i]) < biggestDid {
					left, right := pointers[i], counts[i]-1
					lastValidMid := uint32(math.MaxUint32) // last valid position.
					for left <= right {
						mid := left + (right-left)/2
						if dossiers[i].Get(mid) < biggestDid {
							left = mid + 1
						} else {
							lastValidMid = mid
							right = mid - 1
						}
					}
					// Move pointer past the last position less than biggestDid.
					if lastValidMid != math.MaxUint32 {
						pointers[i] = lastValidMid
					} else {
						pointers[i] = left
					}
					// If we reach the end, return the collected docids.
					if pointers[i] >= counts[i] {
						return docids
					}
				}
			}
		}
	}
}

func removeDuplicateTokenIDs(tids []Amaru.TokenID) []Amaru.TokenID {
	seen := make(map[Amaru.TokenID]struct{})
	j := 0
	for _, tid := range tids {
		if _, exists := seen[tid]; !exists {
			seen[tid] = struct{}{}
			tids[j] = tid
			j++
		}
	}
	return tids[:j]
}

func NewAnthology(anthologyBasePath string, writable bool) (Amaru.Anthology, error) {
	anthology := anthologyImpl{
		aPath:                   anthologyBasePath,
		iPath:                   anthologyBasePath + ".idx",
		lastKnownEoF:            0,
		defaultDossierCapacity:  2_000_000,  // 2M Docs
		defaultAnthologySizeMiB: 10_000_000, // 10TB
		defaultIndexSizeTks:     2_000_000,  // 2M Tokens
	}
	return &anthology, nil
}
