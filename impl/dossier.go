package impl

import (
	"encoding/binary"
	"github.com/kukino/Amaru"
	"log"
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

// newDossier creates a new record at the end of the aMMap.
func (a *anthologyImpl) newDossier(tid Amaru.TokenID, capacity uint32) uint64 {
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
	return offset
}

// dossierAddDocID adds a DocID to an existing Dossier
func (a *anthologyImpl) dossierAddDocID(offset uint64, did Amaru.DocID) bool {
	if offset > uint64(len(a.aMMap)) {
		return false
	}
	capacity := binary.BigEndian.Uint32(a.aMMap[offset+Dossier_CapacityOfs:])
	count := binary.BigEndian.Uint32(a.aMMap[offset+Dossier_CountOfs:])
	if count >= capacity {
		return false
	}
	count++
	binary.BigEndian.PutUint32(a.aMMap[offset+Dossier_CountOfs:], count)
	docIdOfs := offset + Dossier_RecordSize + Dossier_DocIDSize*uint64(count)
	binary.BigEndian.PutUint32(a.aMMap[docIdOfs:], uint32(did))
	return true
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
