package impl

import (
	Amaru "github.com/kukino/Amaru"
	"path"
)

type amaruImpl struct {
	tokens    Amaru.Tokens
	documents Amaru.Documents
}

func (a *amaruImpl) Load() error {
	if err := a.tokens.Load(); err != nil {
		return err
	}
	if err := a.documents.Save(); err != nil {
		return err
	}
	return nil
}

func (a *amaruImpl) Save() error {
	if err := a.tokens.Save(); err != nil {
		return err
	}
	if err := a.documents.Save(); err != nil {
		return err
	}
	return nil
}

func (a *amaruImpl) Clear() {
	a.tokens.Clear()
	a.documents.Clear()
}

func (a *amaruImpl) Tokens() Amaru.Tokens {
	return a.tokens
}

func (a *amaruImpl) Documents() Amaru.Documents {
	return a.documents
}

func NewAmaru(storageFolder string, writable bool) Amaru.Amaru {
	tokensFile := path.Join(storageFolder, "tokens")
	documentsFile := path.Join(storageFolder, "documents")
	impl := amaruImpl{
		tokens:    NewTokens(tokensFile, writable),
		documents: NewDocuments(documentsFile, writable),
	}
	impl.Load()
	return &impl
}
