package impl

import (
	"github.com/edsrzf/mmap-go"
	"github.com/kukino/Amaru"
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
	var position []uint32
	var counts []uint32

	tids = removeDuplicateTokenIDs(tids)

	if len(tids) == 0 {
		return docids
	}

	for _, tid := range tids {
		dossier := a.GetDossier(tid)
		dossiers = append(dossiers, dossier)
		if dossier.Count() == 0 {
			return docids
		}
		position = append(position, 0)
		counts = append(counts, dossier.Count())
	}

	for {
		// Assume match found; verify against first dossier's current DID.
		did := dossiers[0].Get(position[0])
		smallestTid := Amaru.InvalidDocID
		smallestPos := -1
		match := true
		for i := 0; i < len(dossiers); i++ {
			currentDid := dossiers[i].Get(position[i])
			if i != 0 && currentDid != did { // If any dossier's current did don't match, not a match.
				match = false
			}
			if currentDid < smallestTid { // Determine if current did is smallest to decide next position increment.
				smallestPos = i
				smallestTid = currentDid
			}
			if position[i] == counts[i] { // end of any dossier? we are done
				return docids
			}
		}
		if match {
			docids = append(docids, did)
			if len(docids) >= limit {
				return docids
			}
		}
		position[smallestPos]++ // Increment position in dossier with smallest DID to move forward.
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
