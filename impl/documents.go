package impl

import (
	"bufio"
	"fmt"
	"github.com/kukino/Amaru"
	"os"
	"strconv"
	"strings"
)

type documentsImpl struct {
	path      string
	writable  bool
	documents []Amaru.Document
	cache     map[string]Amaru.DocID
}

func (d *documentsImpl) Clear() {
	d.documents = nil
	d.cache = map[string]Amaru.DocID{}
}

func (d *documentsImpl) Get(did Amaru.DocID) *Amaru.Document {
	if int(did) >= d.Count() {
		return nil
	}
	return &d.documents[int(did)]
}

func (d *documentsImpl) Count() int {
	return len(d.documents)
}

func (d *documentsImpl) Add(url string, ranking float32) Amaru.DocID {
	if !d.writable {
		return Amaru.InvalidDocID
	}
	if did, exist := d.cache[url]; exist {
		d.documents[did].Ranking = ranking
		return did
	}
	d.documents = append(d.documents, Amaru.Document{URL: url, Ranking: ranking})
	did := Amaru.DocID(d.Count() - 1)
	d.cache[url] = did
	return did
}

func (d *documentsImpl) Load() error {
	d.Clear()

	file, err := os.Open(d.path)
	if err != nil {
		return err
	}
	defer file.Close()

	d.Clear()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) != 2 {
			return fmt.Errorf("document file, invalid line format: %s", line)
		}

		ranking, err := strconv.ParseFloat(parts[0], 32)
		if err != nil {
			return err
		}

		d.documents = append(d.documents, Amaru.Document{
			URL:     parts[1],
			Ranking: float32(ranking),
		})
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *documentsImpl) Save() error {
	file, err := os.Create(d.path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, doc := range d.documents {
		_, err := fmt.Fprintf(file, "%.6f\t%s\n", doc.Ranking, doc.URL)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewDocuments(documentsFile string, writable bool) Amaru.Documents {
	documents := documentsImpl{
		path:     documentsFile,
		writable: writable,
	}
	documents.Load() // ignore error, it is OK if file does not exist
	return &documents
}
