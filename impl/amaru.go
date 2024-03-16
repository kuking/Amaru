package impl

import (
	"errors"
	Amaru "github.com/kukino/Amaru"
	"os"
	"path"
)

const (
	newFileFlags int         = os.O_RDWR | os.O_CREATE
	newFilePerms os.FileMode = 0644
)

type amaruImpl struct {
	writable  bool
	path      string
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
	if err := a.anthology.Load(); err != nil {
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
	if err := a.anthology.Save(); err != nil {
		return err
	}
	return nil
}

func (a *amaruImpl) Exist() bool {
	return a.tokens.Exist() && a.documents.Exist() && a.anthology.Exist()
}

func (a *amaruImpl) Clear() {
	a.tokens.Clear()
	a.documents.Clear()
	a.anthology.Clear()
}

func (a *amaruImpl) Create() error {
	if err := a.tokens.Create(); err != nil {
		return err
	}
	if err := a.documents.Create(); err != nil {
		return err
	}
	if err := a.anthology.Create(); err != nil {
		return err
	}
	return nil
}

func (a *amaruImpl) Tokens() Amaru.Tokens {
	return a.tokens
}

func (a *amaruImpl) Documents() Amaru.Documents {
	return a.documents
}

func (a *amaruImpl) Anthology() Amaru.Anthology {
	return a.anthology
}

func (a *amaruImpl) Path() string {
	return a.path
}

func NewAmaru(storageFolder string, writable bool) (Amaru.Amaru, error) {
	tokensFile := path.Join(storageFolder, "tokens")
	documentsFile := path.Join(storageFolder, "documents")
	anthologyFile := path.Join(storageFolder, "anthology")
	anthology, err := NewAnthology(anthologyFile, writable)
	if err != nil {
		return nil, err
	}
	impl := amaruImpl{
		writable:  writable,
		path:      storageFolder,
		tokens:    NewTokens(tokensFile, writable),
		documents: NewDocuments(documentsFile, writable),
		anthology: anthology,
	}
	return &impl, nil
}
