package impl

import (
	"errors"
	Amaru "github.com/kukino/Amaru"
	"path"
)

type amaruImpl struct {
	writable  bool
	tokens    Amaru.Tokens
	documents Amaru.Documents
	anthology Amaru.Anthology
}

func (a *amaruImpl) Load() error {
	if err := a.tokens.Load(); err != nil {
		return err
	}
	if err := a.documents.Load(); err != nil {
		return err
	}
	return nil
}

func (a *amaruImpl) Save() error {
	if !a.writable {
		return errors.New("not writable")
	}
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

func NewAmaru(storageFolder string, writable bool) (Amaru.Amaru, error) {
	tokensFile := path.Join(storageFolder, "tokens")
	documentsFile := path.Join(storageFolder, "documents")
	anthologyFile := path.Join(storageFolder, "anthology")
	anthologyIndexFile := path.Join(storageFolder, "anthology.idx")
	anthology, err := NewAnthology(anthologyFile, anthologyIndexFile, writable)
	if err != nil {
		return nil, err
	}
	impl := amaruImpl{
		writable:  writable,
		tokens:    NewTokens(tokensFile, writable),
		documents: NewDocuments(documentsFile, writable),
		anthology: anthology,
	}
	if err = impl.Load(); err != nil {
		return nil, err
	}
	return &impl, nil
}
