package impl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/kukino/Amaru"
	"io"
	"os"
	"unicode/utf8"
)

type tokensImpl struct {
	path     string
	writable bool
	cache    map[Amaru.TokenType]map[string]Amaru.TokenID
	tokens   []Amaru.Token
}

func (t *tokensImpl) Get(tid Amaru.TokenID) *Amaru.Token {
	if int(tid) < len(t.tokens) {
		return &t.tokens[int(tid)]
	}
	return nil
}

func (t *tokensImpl) GetId(tokenType Amaru.TokenType, text string) Amaru.TokenID {
	if m, exist := t.cache[tokenType]; exist {
		if tid, exist := m[text]; exist {
			return tid
		}
	}
	return Amaru.InvalidTokenID
}

func (t *tokensImpl) Count() int {
	return len(t.tokens)
}

func (t *tokensImpl) Add(tokenType Amaru.TokenType, text string) Amaru.TokenID {
	if !t.writable {
		return Amaru.InvalidTokenID
	}

	tid := t.GetId(tokenType, text)
	if tid != Amaru.InvalidTokenID {
		return tid
	}

	tid = Amaru.TokenID(t.Count())

	text = sanitiseTokenText(text)

	token := Amaru.Token{
		Type: tokenType,
		Text: text,
	}
	t.tokens = append(t.tokens, token)
	t.cache[tokenType][text] = tid

	return tid
}

func sanitiseTokenText(text string) string {
	for len([]byte(text)) > Amaru.MaxTokenLen {
		r, size := utf8.DecodeLastRuneInString(text)
		if r == utf8.RuneError && size == 0 {
			// error, or not valid UTF-8
			text = text[:len(text)-1]
		} else {
			text = text[:len(text)-size]
		}
	}
	return text
}

func (t *tokensImpl) Load() error {
	file, err := os.Open(t.path)
	if err != nil {
		return err
	}
	defer file.Close()
	t.Clear()
	for {
		var tType uint8
		if err := binary.Read(file, binary.LittleEndian, &tType); err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			return nil
		}
		textBytes := make([]byte, Amaru.MaxTokenLen)
		if _, err := io.ReadFull(file, textBytes); err != nil {
			if err == io.EOF {
				break // End of file reached
			}
			return errors.New("tokens: incomplete record in file")
		}
		text := string(bytes.TrimRight(textBytes, "\x00"))

		t.Add(Amaru.TokenType(tType), text)
	}
	return nil
}

func (t *tokensImpl) Save() error {
	file, err := os.Create(t.path)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, token := range t.tokens {
		if err := binary.Write(file, binary.LittleEndian, uint8(token.Type)); err != nil {
			return err
		}
		text := []byte(token.Text)
		if len(text) < Amaru.MaxTokenLen {
			padding := make([]byte, Amaru.MaxTokenLen-len(text))
			text = append(text, padding...)
		} else if len(text) > Amaru.MaxTokenLen {
			text = text[:Amaru.MaxTokenLen]
		}
		if _, err := file.Write(text); err != nil {
			return err
		}
	}
	return nil
}

func (t *tokensImpl) Clear() {
	cache := make(map[Amaru.TokenType]map[string]Amaru.TokenID)
	cache[Amaru.TextToken] = make(map[string]Amaru.TokenID)
	cache[Amaru.TagToken] = make(map[string]Amaru.TokenID)
	t.cache = cache
	t.tokens = []Amaru.Token{}
}

func NewTokens(tokensFile string, writable bool) Amaru.Tokens {

	cache := make(map[Amaru.TokenType]map[string]Amaru.TokenID)
	cache[Amaru.TextToken] = make(map[string]Amaru.TokenID)
	cache[Amaru.TagToken] = make(map[string]Amaru.TokenID)
	tokens := tokensImpl{
		path:     tokensFile,
		writable: writable,
	}
	tokens.Load() // ignore error, it is OK if file does not exist
	return &tokens
}
