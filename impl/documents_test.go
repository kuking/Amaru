package impl

import (
	"fmt"
	"github.com/kuking/Amaru"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const EPSILON float64 = 0.0000001

func TestGetDocument(t *testing.T) {
	testFilePath := getTempFile(t, "test_documents_*")
	defer os.Remove(testFilePath)

	docs := NewDocuments(testFilePath, true)
	did := docs.Add("some://lala", 0.2)

	doc := docs.Get(did)
	assert.NotNil(t, doc)
	assert.Equal(t, doc.URL, "some://lala")
	assert.InEpsilon(t, doc.Ranking, 0.2, EPSILON)
}

func TestDocumentURLIsUnique(t *testing.T) {
	testFilePath := getTempFile(t, "test_documents_*")
	defer os.Remove(testFilePath)

	docs := NewDocuments(testFilePath, true)
	did1 := docs.Add("some://lala", 0.2)
	did2 := docs.Add("some://lala", 0.35)
	did3 := docs.Add("some://lala", 0.45)
	did4 := docs.Add("some://lele", 0.69)
	assert.Equal(t, did1, did2)
	assert.Equal(t, did2, did3)
	assert.NotEqual(t, did3, did4)

	doc123 := docs.Get(did1)
	doc4 := docs.Get(did4)

	assert.NotNil(t, doc123)
	assert.NotNil(t, doc4)
	assert.Equal(t, doc123.URL, "some://lala")
	assert.Equal(t, doc4.URL, "some://lele")
	assert.InEpsilon(t, doc123.Ranking, 0.45, EPSILON)
	assert.InEpsilon(t, doc4.Ranking, 0.69, EPSILON)
}

func TestReadOnlyDocuments(t *testing.T) {
	testFilePath := getTempFile(t, "test_documents_*")
	defer os.Remove(testFilePath)

	docs := NewDocuments(testFilePath, false)
	did := docs.Add("some://url", 1)
	assert.Equal(t, Amaru.InvalidDocID, did)

	assert.Error(t, docs.Save())
}

func TestDocumentsSaveLoad(t *testing.T) {
	testFilePath := getTempFile(t, "test_documents_*")
	defer os.Remove(testFilePath)

	docs := NewDocuments(testFilePath, true)

	for n := 0; n < 1_000; n++ {
		docs.Add(fmt.Sprintf("some://%d", n), float32(n/1_0000))
	}
	assert.Equal(t, docs.Count(), 1_000)
	assert.NoError(t, docs.Save())

	docs2 := NewDocuments(testFilePath, true)
	assert.Equal(t, docs.Count(), docs2.Count())
	for n := 0; n < docs.Count(); n++ {
		tl := docs.Get(Amaru.DocID(n))
		tr := docs2.Get(Amaru.DocID(n))
		assert.Equal(t, tl, tr)
	}
}
