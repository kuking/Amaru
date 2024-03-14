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

func (a *anthologyImpl) Dossier(tid Amaru.TokenID) *[]Amaru.DocID {
	//TODO implement me
	panic("implement me")
}

func (a *anthologyImpl) Add(did Amaru.DocID, tid Amaru.TokenID) {
	offset := a.tidOffset(tid)
	if offset == uint64(0xffff_ffff_ffff_ffff) {
		offset = a.newDossier(tid, a.defaultDossierCapacity)
	}

	if !a.dossierAddDocID(offset, did) {
		panic("could not add the did into this tid ...have you run out of initial space?")
	}
}

func (a *anthologyImpl) Compact() error {
	//TODO implement me
	panic("implement me")
}

func (a *anthologyImpl) Load() error {
	var err error
	a.aFile, err = os.OpenFile(a.aPath, os.O_RDWR|os.O_CREATE, 0755)
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

	a.iFile, err = os.OpenFile(a.iPath, os.O_RDWR|os.O_CREATE, 0755)
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

func NewAnthology(anthologyBasePath string, writable bool) (Amaru.Anthology, error) {
	anthology := anthologyImpl{
		aPath:                   anthologyBasePath,
		iPath:                   anthologyBasePath + ".idx",
		defaultDossierCapacity:  50_000,
		defaultAnthologySizeMiB: 10_000,
		defaultIndexSizeTks:     100_000,
	}
	return &anthology, nil
}
