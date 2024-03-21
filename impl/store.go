package impl

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/edsrzf/mmap-go"
	"github.com/kuking/Amaru"
	"os"
	"sort"
	"strconv"
	"strings"
)

type storeImpl struct {
	iPath               string
	mPath               string
	mFile               *os.File
	firstFreeByte       uint64
	mmap                mmap.MMap
	writable            bool
	keyCache            map[string]uint64
	idCache             map[uint32]uint64
	idKeyCache          map[uint32]string
	defaultStoreSizeMiB int64
}

func (s *storeImpl) GetId(key string) uint32 {
	//if exist, id := s.keyCache[key]; exist {
	//	return id
	//} else {
	//	return
	//}
	//TODO implement me
	panic("implement me")
}

func (s *storeImpl) getByOffset(offset uint64) []byte {
	length := binary.BigEndian.Uint32(s.mmap[offset:])
	return s.mmap[offset+4 : offset+4+uint64(length)]
}
func (s *storeImpl) GetByKey(key string) []byte {
	if offset, exist := s.keyCache[key]; exist {
		return s.getByOffset(offset)
	}
	return nil
}

func (s *storeImpl) GetById(id uint32) []byte {
	if offset, exist := s.idCache[id]; exist {
		return s.getByOffset(offset)
	}
	return nil
}

func (s *storeImpl) Set(key string, id uint32, data []byte) bool {
	if _, exist := s.keyCache[key]; exist {
		return false
	}
	if _, exist := s.idCache[id]; exist {
		return false
	}

	length := uint32(len(data))
	offset := s.firstFreeByte

	binary.BigEndian.PutUint32(s.mmap[offset:], length)
	copy(s.mmap[offset+4:], data) // compression here?
	s.firstFreeByte += 4 + uint64(length)

	s.keyCache[key] = offset
	s.idCache[id] = offset
	s.idKeyCache[id] = key

	return true
}

func (s *storeImpl) Clear() {
	if err := s.Close(); err != nil {
		// nothing
	}
	s.idCache = make(map[uint32]uint64)
	s.keyCache = make(map[string]uint64)
	s.idKeyCache = make(map[uint32]string)
	s.firstFreeByte = 0
	if err := os.Remove(s.iPath); err != nil {
		// pass
	}
	if err := os.Remove(s.mPath); err != nil {
		// pass
	}
}

func (s *storeImpl) Close() error {
	if err := s.mmap.Unmap(); err != nil {
		return err
	}
	if err := s.mFile.Close(); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) Create() error {
	s.Clear()
	return s.Load()
}

func (s *storeImpl) Load() error {
	var err error
	s.mFile, err = os.OpenFile(s.mPath, newFileFlags, newFilePerms)
	if err != nil {
		return err
	}
	if fi, err := s.mFile.Stat(); err == nil {
		if fi.Size() == 0 {
			if err := s.mFile.Truncate(s.defaultStoreSizeMiB * 1024 * 1024); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	s.mmap, err = mmap.Map(s.mFile, mmap.RDWR, 0)
	if err != nil {
		return err
	}

	// now reads the index, index format is text:
	// 1st line: first free byte (0 by default)
	// 2nd line and rest:
	// [offset in hex] \t [id as in hex] \t [path]
	s.idCache = make(map[uint32]uint64)
	s.keyCache = make(map[string]uint64)
	s.idKeyCache = make(map[uint32]string)
	s.firstFreeByte = 0

	iFile, err := os.Open(s.iPath)
	defer iFile.Close()
	if err != nil {
		// assumes nothing is there
		return nil
	}

	scanner := bufio.NewScanner(iFile)
	if scanner.Scan() {
		s.firstFreeByte, err = strconv.ParseUint(scanner.Text(), 16, 64)
		if err != nil {
			return err
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) != 3 {
			return fmt.Errorf("store index, invalid format in line: %s", line)
		}

		offset, err := strconv.ParseUint(parts[0], 16, 64)
		if err != nil {
			return err
		}
		id, err := strconv.ParseUint(parts[1], 16, 32)
		if err != nil {
			return err
		}
		key := parts[2]

		s.keyCache[key] = offset
		s.idCache[uint32(id)] = offset
		s.idKeyCache[uint32(id)] = key
	}
	return nil
}

func (s *storeImpl) Save() error {
	if !s.writable {
		return errors.New("not writable")
	}

	if err := s.mmap.Flush(); err != nil {
		return err
	}
	if err := s.mFile.Sync(); err != nil {
		return err
	}
	iFile, err := os.OpenFile(s.iPath, newFileFlags, newFilePerms)
	if err != nil {
		return err
	}
	defer iFile.Close()

	var ids []uint32
	for id, _ := range s.idKeyCache {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { // sort by offset, neather
		return s.idCache[ids[i]] < s.idCache[ids[j]]
	})

	if _, err := fmt.Fprintf(iFile, "%016x\n", s.firstFreeByte); err != nil {
		return err
	}
	for _, id := range ids {
		if _, err := fmt.Fprintf(iFile, "%016x\t%08x\t%s\n", s.idCache[id], id, s.idKeyCache[id]); err != nil {
			return err
		}
	}

	return nil
}

func (s *storeImpl) Exist() bool {
	return true
}

func (s *storeImpl) Compact() error {
	if err := s.mFile.Truncate(int64(s.firstFreeByte)); err != nil {
		return err
	}
	return nil
}

func NewStore(storeBasePath string, writable bool) (Amaru.Store, error) {
	store := storeImpl{
		mPath:               storeBasePath,
		iPath:               storeBasePath + ".idx",
		firstFreeByte:       0,
		writable:            writable,
		defaultStoreSizeMiB: 10_000_000, // 10TB
	}

	return &store, nil
}
