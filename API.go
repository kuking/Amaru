package Amaru

type TokenID uint32
type TokenType uint8
type DocID uint32

const (
	TextToken TokenType = 0
	TagToken  TokenType = 1

	InvalidTokenID TokenID = 0xffffffff

	MaxTokenLen int = 25
)

type Amaru interface {
	Tokens() Tokens
	Documents() Documents
	Load() error
	Save() error
}

type Token struct {
	Type TokenType
	Text string
}

type Tokens interface {
	Get(tid TokenID) *Token
	GetId(tokenType TokenType, text string) TokenID
	Count() int
	Add(tokenType TokenType, text string) TokenID
	Load() error
	Save() error
}

type Document struct {
	URL     string
	Ranking float32
}
type Documents interface {
	Get(did DocID) *Document
	Count() int
	Add(url string, ranking float32) DocID
	Load()
	Save()
}
