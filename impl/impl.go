package impl

import (
	Amaru "github.com/kukino/Amaru"
	"path"
)

type amaruImpl struct {
	tokens    Amaru.Tokens
	documents Amaru.Documents
}

func (a amaruImpl) Load() error {
	if err := a.tokens.Load(); err != nil {
		return err
	}
	//TODO implement me
	panic("implement me")
}

func (a amaruImpl) Save() error {
	//TODO implement me
	panic("implement me")
}

func (a amaruImpl) Tokens() Amaru.Tokens {
	return a.tokens
}

func (a amaruImpl) Documents() Amaru.Documents {
	return a.documents
}

func NewAmaru(storagePath string, writable bool) Amaru.Amaru {
	tokensFile := path.Join(storagePath, "tokens")

	impl := amaruImpl{
		tokens:    NewTokens(tokensFile, writable),
		documents: nil,
	}
	return impl
}
