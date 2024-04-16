package impl

import (
	"encoding/binary"
	"errors"
	"github.com/edsrzf/mmap-go"
	"github.com/kuking/Amaru"
	"log"
	"sort"
	"time"
)

const (
	Dossier_TokenIDSize  uint64 = 4 // size of TokenID in bytes
	Dossier_DocIDSize    uint64 = 4 // size of DocID in bytes
	Dossier_CapacitySize uint64 = 4 // size of capacity field in bytes
	Dossier_CountSize    uint64 = 4 // size of count field in bytes
	Dossier_RecordSize   uint64 = Dossier_TokenIDSize + Dossier_CapacitySize + Dossier_CountSize

	Dossier_TokenIDOfs  uint64 = 0
	Dossier_CapacityOfs uint64 = Dossier_TokenIDOfs + Dossier_TokenIDSize
	Dossier_CountOfs    uint64 = Dossier_CapacityOfs + Dossier_CapacitySize
)

type dossierImpl struct {
	aMMap  mmap.MMap
	iMMap  mmap.MMap
	offset uint64
	tid    Amaru.TokenID
}

func (d *dossierImpl) Offset() uint64 {
	return d.offset
}

func (d *dossierImpl) TokenID() Amaru.TokenID {
	return d.tid
}

func (d *dossierImpl) Capacity() uint32 {
	return binary.BigEndian.Uint32(d.aMMap[d.offset+Dossier_CapacityOfs:])
}

// SetCapacity will not allocate the underlying bytes, this is not a general purpose API but internal.
func (d *dossierImpl) SetCapacity(cap uint32) {
	binary.BigEndian.PutUint32(d.aMMap[d.offset+Dossier_CapacityOfs:], cap)
}

func (d *dossierImpl) Count() uint32 {
	return binary.BigEndian.Uint32(d.aMMap[d.offset+Dossier_CountOfs:])
}

func (d *dossierImpl) Get(n uint32) Amaru.DocID {
	if n >= d.Count() {
		return Amaru.InvalidDocID
	}
	return Amaru.DocID(binary.BigEndian.Uint32(d.aMMap[d.offset+Dossier_RecordSize+Dossier_DocIDSize*uint64(n):]))
}

func (d *dossierImpl) Set(n uint32, did Amaru.DocID) {
	if n >= d.Count() {
		panic("you should know where to set inside the dossier")
	}
	binary.BigEndian.PutUint32(d.aMMap[d.offset+Dossier_RecordSize+Dossier_DocIDSize*uint64(n):], uint32(did))
}

func (d *dossierImpl) SizeInBytes() uint64 {
	return Dossier_RecordSize + Dossier_DocIDSize*uint64(d.Capacity())
}

func (d *dossierImpl) Add(did Amaru.DocID) (newCount uint32, err error) {
	if d.Count() == d.Capacity() {
		return d.Capacity(), errors.New("cannot add doc_id, dossier is full")
	}
	n := d.Count()
	binary.BigEndian.PutUint32(d.aMMap[d.offset+Dossier_CountOfs:], n+1)
	d.Set(n, did)
	return n + 1, nil
}

func (d *dossierImpl) Sort() {
	sort.Sort(d)
}

func (d *dossierImpl) Len() int {
	return int(d.Count())
}

func (d *dossierImpl) Less(i, j int) bool {
	return d.Get(uint32(i)) < d.Get(uint32(j))
}

func (d *dossierImpl) Swap(i, j int) {
	did := d.Get(uint32(i))
	d.Set(uint32(i), d.Get(uint32(j)))
	d.Set(uint32(j), did)
}

// newDossier creates a new record at the end of the aMMap.
func (a *anthologyImpl) newDossier(tid Amaru.TokenID, capacity uint32) Amaru.Dossier {
	var offset uint64 = a.lastKnownEoF

	// Scan all records to find the end of the file
	for offset < uint64(len(a.aMMap)) {
		capacity := binary.BigEndian.Uint32(a.aMMap[offset+Dossier_CapacityOfs:])
		if capacity == 0 {
			break
		}
		offset += Dossier_RecordSize + uint64(capacity)*Dossier_DocIDSize
	}
	if offset+Dossier_RecordSize > uint64(len(a.aMMap)) {
		panic("we ran out of space in the anthology")
	}
	tokenId := binary.BigEndian.Uint32(a.aMMap[offset+Dossier_TokenIDOfs:])
	if tokenId > uint32(tid) {
		panic("dossiers should be created in token order, for simplicity")
	}

	a.setTidOffset(tid, offset)
	a.lastKnownEoF = offset

	binary.BigEndian.PutUint32(a.aMMap[offset:], uint32(tid))
	binary.BigEndian.PutUint32(a.aMMap[offset+Dossier_CapacityOfs:], capacity)
	binary.BigEndian.PutUint32(a.aMMap[offset+Dossier_CountOfs:], 0)

	return a.GetDossier(tid)
}

func (a *anthologyImpl) GetDossier(tid Amaru.TokenID) Amaru.Dossier {
	return &dossierImpl{
		aMMap:  a.aMMap,
		iMMap:  a.iMMap,
		offset: a.tidOffset(tid),
		tid:    tid,
	}
}

func (a *anthologyImpl) tidOffset(tid Amaru.TokenID) uint64 {
	if uint32(len(a.iMMap)) < (uint32(tid)+1)*8 {
		log.Fatalf("iMMap index too small: TokenID %d won't fit", tid)
	}
	start := tid * 8
	data := a.iMMap[start:]
	return binary.BigEndian.Uint64(data)
}

func (a *anthologyImpl) setTidOffset(tid Amaru.TokenID, offset uint64) {
	if uint32(len(a.iMMap)) < (uint32(tid)+1)*8 {
		log.Fatalf("iMMap index too small: TokenID %d won't fit", tid)
	}
	start := uint32(tid) * 8
	binary.BigEndian.PutUint64(a.iMMap[start:], offset)
}

func (a *anthologyImpl) Compact() error {
	// this function makes every dossier of the size of current len, shrinking the aMMap and finally truncating it.
	t0 := time.Now()
	tid := Amaru.TokenID(0)
	newOffset := uint64(0) // offset for tid0 is 0
	var saved uint64
	c := 0
	for {
		d := a.GetDossier(tid)
		if d.Offset() == Amaru.InvalidOffset {
			break
		}

		// shrinks Dossier and calculated the saved space
		prevSize := d.SizeInBytes()
		d.SetCapacity(d.Count())
		delta := prevSize - d.SizeInBytes()
		saved += delta

		c++
		if c%100_000 == 99_999 {
			elapsed := time.Since(t0)
			log.Printf("%dk dossiers compacted; %.1fGiB saved; thoughput is %.1f dossiers/sec ...\n", c/1000, float64(saved)/1024.0/1024.0/1024.0, float64(c)/elapsed.Seconds())
		}

		// new Offset
		if newOffset == d.Offset() {
			// OK we are in the same place, we do nothing.
		} else {
			copy(a.aMMap[newOffset:], a.aMMap[d.Offset():d.Offset()+d.SizeInBytes()])
			a.setTidOffset(tid, newOffset)
		}

		newOffset += d.SizeInBytes()
		tid++
	}
	if fi, err := a.aFile.Stat(); err != nil {
		return err
	} else {
		saved += uint64(fi.Size() - int64(newOffset))
	}
	if err := a.aFile.Truncate(int64(newOffset)); err != nil {
		return err
	}

	iTrunc := int64(tid+1) * 8 // size of uint64
	if fi, err := a.iFile.Stat(); err != nil {
		return err
	} else {
		saved += uint64(fi.Size() - iTrunc)
	}
	if err := a.iFile.Truncate(iTrunc); err != nil {
		return err
	}

	log.Printf("Compacted Anthology is %.1fMiB (%.2fGiB saved)\n", float64(newOffset)/1024.0/1024.0, float64(saved)/1024.0/1024.0/1024.0)
	log.Println("Closing after Compacting ...")
	if err := a.Close(); err != nil {
		return err
	}
	log.Println("Reloading after Compacting ...")
	if err := a.Load(); err != nil {
		return err
	}
	elapsed := time.Since(t0)
	log.Printf("Compacting finished Successfully in %v!\n", elapsed.Truncate(time.Second))
	return nil
}
