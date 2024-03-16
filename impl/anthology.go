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

func (a *anthologyImpl) FindDocIDsWith(tids []Amaru.TokenID) []Amaru.DocID {

	var docids []Amaru.DocID
	var dossiers []Amaru.Dossier
	var position []uint32
	for _, tid := range tids {
		dossiers = append(dossiers, a.GetDossier(tid))
		position = append(position, 0)
	}

	if len(tids) == 0 {
		return docids
	}

	finished := false
	for {
		// checks if the end of a dossier has been reach
		for i := 0; i < len(dossiers); i++ {
			if dossiers[i].Count() == position[i] {
				finished = true
				break
			}
		}
		if finished {
			break
		}

		// checks if the current index is the same for all dossiers => match
		found := true
		did := dossiers[0].Get(position[0])
		for i := 1; i < len(dossiers); i++ {
			if did != dossiers[i].Get(position[i]) {
				found = false
			}
			if !found {
				break
			}
		}
		if found {
			docids = append(docids, did)
		}

		// increases the dossier with the smallest number
		smallestTid := Amaru.InvalidDocID // highest possible token id
		smallestPos := -1
		for i := 0; i < len(dossiers); i++ {
			if dossiers[i].Get(position[i]) < smallestTid {
				smallestPos = i
				smallestTid = dossiers[i].Get(position[i])
			}
		}
		position[smallestPos]++

	}

	return docids
}

func NewAnthology(anthologyBasePath string, writable bool) (Amaru.Anthology, error) {
	anthology := anthologyImpl{
		aPath:                   anthologyBasePath,
		iPath:                   anthologyBasePath + ".idx",
		defaultDossierCapacity:  100_000,
		defaultAnthologySizeMiB: 100_000,
		defaultIndexSizeTks:     100_000,
	}
	return &anthology, nil
}
