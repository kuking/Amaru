package impl

import (
	"encoding/binary"
	"errors"
	"github.com/edsrzf/mmap-go"
	"github.com/kukino/Amaru"
	"log"
	"sort"
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
	var offset uint64 = 0

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
