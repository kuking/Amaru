package Amaru

type TokenID uint32
type TokenType uint8
type DocID uint32

const (
	TextToken TokenType = 0
	TagToken  TokenType = 1

	InvalidTokenID TokenID = 0xffff_ffff
	InvalidDocID   DocID   = 0xfff_fffff
	InvalidOffset  uint64  = 0xffff_ffff_ffff_ffff
	MaxTokenLen    int     = 25
)

type Amaru interface {
	Tokens() Tokens
	Documents() Documents
	Anthology() Anthology
	Path() string
	Load() error
	Save() error
	Exist() bool
	Clear()
	Create() error
}

type Token struct {
	Type TokenType
	Text string
}

type Tokens interface {
	Get(tid TokenID) *Token
	GetId(tokenType TokenType, text string) TokenID
	GetIds(tokenType TokenType, texts []string) []TokenID
	Count() int
	Add(tokenType TokenType, text string) (TokenID, string)
	Load() error
	Save() error
	Exist() bool
	Clear()
	Create() error
}

type Document struct {
	Did     DocID
	URL     string
	Ranking float32
}

type Documents interface {
	Get(did DocID) *Document
	GetAll(docids []DocID) []*Document
	Count() int
	Add(url string, ranking float32) DocID
	Load() error
	Save() error
	Exist() bool
	Clear()
	Create() error
}

type Anthology interface {
	GetDossier(tid TokenID) Dossier // Readonly
	Add(did DocID, tid TokenID)
	Compact() error
	Load() error
	Save() error
	Exist() bool
	Clear()
	Close() error
	Create() error
	FindDocIDsWith(tids []TokenID, limit int) []DocID
}

type Dossier interface {
	Offset() uint64
	TokenID() TokenID
	Capacity() uint32
	Count() uint32
	Get(n uint32) DocID
	Set(n uint32, did DocID)
	Add(did DocID) (newCap uint32, err error)
	Sort()
	// Data Interface
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	// Internals but useful
	SetCapacity(uint32)
	SizeInBytes() uint64
}
