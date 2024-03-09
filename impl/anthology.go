package impl

import (
	"github.com/edsrzf/mmap-go"
	"github.com/kukino/Amaru"
	"os"
)

type anthologyImpl struct {
	aFile *os.File
	aSize int64
	aMMap mmap.MMap
	iFile *os.File
	iSize int64
	iMMap mmap.MMap
}

func (a *anthologyImpl) Dossier(tid Amaru.TokenID) *[]Amaru.DocID {
	//TODO implement me
	panic("implement me")
}

func (a *anthologyImpl) Add(tid Amaru.TokenID, did Amaru.DocID) {
	//TODO implement me
	panic("implement me")
}

func (a *anthologyImpl) Load() error {
	//TODO implement me
	panic("implement me")
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

func NewAnthology(anthologyFile string, anthologyIndex string, writable bool) (Amaru.Anthology, error) {

	var err error
	anthology := anthologyImpl{}

	anthology.aFile, err = os.OpenFile(anthologyFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	if fi, err := anthology.aFile.Stat(); err == nil {
		if fi.Size() == 0 {
			anthology.aFile.Truncate(1_000_000_000)
			anthology.aSize = 1_000_000_000
		} else {
			anthology.aSize = fi.Size()
		}
	} else {
		return nil, err
	}
	anthology.aMMap, err = mmap.Map(anthology.aFile, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}

	anthology.iFile, err = os.OpenFile(anthologyIndex, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	empty := false
	if fi, err := anthology.iFile.Stat(); err == nil {
		if fi.Size() == 0 {
			anthology.iFile.Truncate(100_000)
			anthology.iSize = 1000_000
			empty = true
		} else {
			anthology.iSize = fi.Size()
		}
	} else {
		return nil, err
	}
	anthology.iMMap, err = mmap.Map(anthology.iFile, mmap.RDWR, 0)
	if err != nil {
		return nil, err
	}
	if empty {
		for i := 0; i < len(anthology.iMMap); i++ {
			anthology.iMMap[i] = 0xff
		}
	}

	return &anthology, nil

}
