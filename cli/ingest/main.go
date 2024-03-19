package main

import (
	"encoding/json"
	"fmt"
	"github.com/kukino/Amaru"
	"github.com/kukino/Amaru/impl"
	"github.com/kukino/Amaru/text"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type JSON map[string]interface{}
type Document struct {
	Json JSON
}

type StemmedDoc struct {
	Url     string
	Ranking float32
	Stems   []string
	Store   []byte
}

func ingest() {
	t0 := time.Now()
	log.Println("Ingesting ...")

	basePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	amaru, err := impl.NewAmaru(path.Join(basePath, "/KUKINO/GO/Amaru/tmp/idx1/"), true)
	if err != nil {
		panic(err)
	}
	if err := amaru.Create(); err != nil {
		panic(err)
	}
	store, err := impl.NewStore(path.Join(basePath, "/KUKINO/GO/Amaru/tmp/idx1/profiles"), true)
	if err != nil {
		panic(err)
	}
	if err := store.Create(); err != nil {
		panic(err)
	}

	tokens := amaru.Tokens()
	docs := amaru.Documents()
	anth := amaru.Anthology()

	docsChan := make(chan Document, 100)
	go readJsons(path.Join(basePath, "/KUKINO/GO/Amaru/tmp/dataset/profiles/"), docsChan)

	stemsChan := make(chan StemmedDoc, 100)
	var wg sync.WaitGroup
	for th := 0; th < 20; th++ { // 20 threads stemming, etc.
		wg.Add(1)
		go stemDocuments(docsChan, stemsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(stemsChan)
	}()

	tp0 := time.Now()
	log.Println("Preloading: loading, parsing, stemming then sorting ...")
	c := 0
	var stemsAll []StemmedDoc
	for stem := range stemsChan {
		stemsAll = append(stemsAll, stem)
		c++
		if c%100000 == 99999 {
			elapsed := time.Since(tp0)
			log.Printf("%vk documents prepared; thoughput is %.1f docs/s ...\n", c/1_000, float64(c)/elapsed.Seconds())
		}
	}

	log.Println("Sorting documents by ranking so dossiers do not need sorting during search")
	sort.Slice(stemsAll, func(i, j int) bool {
		return stemsAll[i].Ranking > stemsAll[j].Ranking
	})
	log.Println("Sorted done")

	ti0 := time.Now()
	c = 0
	for _, stem := range stemsAll { // stemsChan {

		var tids []Amaru.TokenID
		for _, oneStem := range stem.Stems {
			tid, _ := tokens.Add(Amaru.TextToken, oneStem)
			tids = append(tids, tid)
		}

		// tokenIds must be added in order so the anthology can be Compacted
		sort.Slice(tids, func(i, j int) bool {
			return tids[i] < tids[j]
		})

		did := docs.Add(stem.Url, stem.Ranking)
		for _, tid := range tids {
			anth.Add(did, tid)
		}
		store.Set(stem.Url, uint32(did), stem.Store)

		if c%100_000 == 0 && c > 0 {
			elapsed := time.Since(ti0)
			log.Printf("%dk documents ingested; thoughput is %.1f docs/s ...\n", did/1000, float64(c)/elapsed.Seconds())
		}

		c++
		if c%5_000_000_000 == 0 { // never, really
			if err := amaru.Save(); err != nil {
				panic(err)
			}
			println("saved")
		}
		if c > 2_000_000 {
			break
		}
	}

	log.Println("Compacting anthology ...")
	if err = anth.Compact(); err != nil {
		log.Fatal(err)
	}
	log.Println("Saving Amaru ... ")
	if err := amaru.Save(); err != nil {
		panic(err)
	}
	log.Println("Compacting and saving store ...")
	if err := store.Compact(); err != nil {
		panic(err)
	}
	if err := store.Save(); err != nil {
		panic(err)
	}
	elapsed := time.Since(t0)
	log.Printf("Ingestion took %v; Total throughput was %.1f docs/s", elapsed.Truncate(time.Millisecond), float64(c)/elapsed.Seconds())
}

func readJsons(basePath string, ch chan Document) {
	defer close(ch)

	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var jsonData JSON
			if err := json.Unmarshal(data, &jsonData); err != nil {
				return err
			}

			ch <- Document{Json: jsonData}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through files:", err)
	}
}

func stemDocuments(docsChan chan Document, stemChan chan StemmedDoc, wg *sync.WaitGroup) {
	defer wg.Done()
	for doc := range docsChan {

		handle := doc.Json["handle"].(string)
		description := doc.Json["desc.txt"].(string)
		rating := math.Max(doc.Json["likes"].(float64), 0)
		rating += math.Max(doc.Json["followers"].(float64), 0)
		rating += math.Max(doc.Json["lives.qty"].(float64), 0)
		rating += math.Max(doc.Json["posts.qty"].(float64), 0)
		rating += math.Max(doc.Json["media.qty"].(float64), 0)
		rating += math.Max(doc.Json["videos.qty"].(float64), 0)
		rating += math.Max(doc.Json["pics.qty"].(float64), 0)
		location := ""
		if doc.Json["loc"] != nil {
			location = doc.Json["loc"].(string)
		}

		// tags also should be added here

		doc := description + " " + location + " " + handle
		doc = text.RemoveBOM(doc)
		doc = text.RemoveEmojis(doc)
		doc = text.NormaliseFancyUnicodeToASCII(doc)
		doc = text.ReplaceStopWords(doc)

		sd := StemmedDoc{
			Url:     fmt.Sprintf("pof://%s", handle),
			Ranking: float32(rating),
			Stems:   text.Stems(doc),
			Store:   []byte(description),
		}
		stemChan <- sd
	}
}

func main() {
	ingest()
}
