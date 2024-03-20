package impl

import (
	"fmt"
	"github.com/kuking/Amaru"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetToken(t *testing.T) {
	testFilePath := getTempFile(t, "test_tokens_*")
	defer os.Remove(testFilePath)

	tokens := NewTokens(testFilePath, true)
	tid, _ := tokens.Add(Amaru.TextToken, "example")

	token := tokens.Get(tid)
	assert.NotNil(t, token)
	assert.Equal(t, token.Text, "example")
	assert.Equal(t, token.Type, Amaru.TextToken)
}

func TestGetTokenMultipleTypesAndTimes(t *testing.T) {
	testFilePath := getTempFile(t, "test_tokens_*")
	defer os.Remove(testFilePath)

	tokens := NewTokens(testFilePath, true)
	tid1, _ := tokens.Add(Amaru.TextToken, "example")
	tid2, _ := tokens.Add(Amaru.TagToken, "example")
	assert.NotEqual(t, t, tid1, tid2, "Different types should hold different values")

	tid1Again, _ := tokens.Add(Amaru.TextToken, "example")
	tid2Again, _ := tokens.Add(Amaru.TagToken, "example")

	assert.Equal(t, tid1, tid1Again, "Adding an existing token should return the existing ID")
	assert.Equal(t, tid2, tid2Again, "Adding an existing token should return the existing ID")
}

func TestReadOnlyTokens(t *testing.T) {
	testFilePath := getTempFile(t, "test_tokens_*")
	defer os.Remove(testFilePath)

	tokens := NewTokens(testFilePath, false)
	res, _ := tokens.Add(Amaru.TextToken, "test")
	assert.Equal(t, Amaru.InvalidTokenID, res)

	assert.Error(t, tokens.Save())
}

func TestTokensSaveLoad(t *testing.T) {
	testFilePath := getTempFile(t, "test_tokens_*")
	defer os.Remove(testFilePath)

	tokens := NewTokens(testFilePath, true)
	for n := 0; n < 1_000; n++ {
		var tokenType Amaru.TokenType
		if n%2 == 0 {
			tokenType = Amaru.TextToken
		} else {
			tokenType = Amaru.TagToken
		}
		text := fmt.Sprintf("example:%d", n)
		tokens.Add(tokenType, text)
	}
	assert.Equal(t, tokens.Count(), 1_000)
	assert.NoError(t, tokens.Save())

	tokens2 := NewTokens(testFilePath, true)
	assert.Equal(t, tokens.Count(), tokens2.Count())
	for n := 0; n < tokens2.Count(); n++ {
		tl := tokens.Get(Amaru.TokenID(n))
		tr := tokens2.Get(Amaru.TokenID(n))
		assert.Equal(t, tl, tr)
	}
}

func TestAddLongText(t *testing.T) {
	testFilePath := getTempFile(t, "test_tokens_*")
	defer os.Remove(testFilePath)

	text := "Smiles😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀😀"
	assert.True(t, len([]byte(text)) > Amaru.MaxTokenLen)
	tokens := NewTokens(testFilePath, true)
	tid, _ := tokens.Add(Amaru.TextToken, text)

	token := tokens.Get(tid)
	assert.True(t, len([]byte(token.Text)) < Amaru.MaxTokenLen)
}
