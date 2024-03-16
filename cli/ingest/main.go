package main

import (
	"encoding/json"
	"fmt"
	"github.com/kukino/Amaru"
	"github.com/kukino/Amaru/impl"
	"github.com/kukino/Amaru/text"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

// TODO: Ingest Documents sorted by descending ranking, so the intersection can be stopped at X results, and they will be the top ones.
// did=1 is the highest document existing, did=2 is the second, and so fort ... therefore the intersection can stop as soon as it finds X elements
// and they will be the top elements.

type JSON map[string]interface{}
type Document struct {
	Json JSON
}

type StemmedDoc struct {
	Url     string
	Ranking float32
	Stems   []string
}

func ingest() {
	t0 := time.Now()
	log.Println("Indexing ...")

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

	tokens := amaru.Tokens()
	docs := amaru.Documents()
	anth := amaru.Anthology()

	docsChan := make(chan Document, 100)
	go readJsons(path.Join(basePath, "/KUKINO/GO/Amaru/tmp/dataset/profiles/"), docsChan)
	stemsChan := make(chan StemmedDoc, 100)
	var wg sync.WaitGroup
	for th := 0; th < 100; th++ { // 100 threads stemming, etc.
		wg.Add(1)
		go stemDocuments(docsChan, stemsChan, &wg)
	}
	go func() {
		wg.Wait()
		close(stemsChan)
	}()

	c := 0
	for stem := range stemsChan {

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

		if c%10_000 == 0 && c > 0 {
			elapsed := time.Since(t0)
			log.Printf("%d documents ingested; thoughput is %.2f docs/s\n", did, float64(c)/elapsed.Seconds())
		}

		c++
		if c%5_000_000_000 == 0 { // never, really
			println(c, "...")
			if err := amaru.Save(); err != nil {
				panic(err)
			}
			println("saved")
		}
		if c > 2_000_000 {
			break
		}
	}

	log.Println("Sorting all Dossiers, not really necessary but ... ")
	log.Println("Dossiers (one per Token):", tokens.Count())
	for t := 0; t < tokens.Count(); t++ {
		anth.GetDossier(Amaru.TokenID(t)).Sort()
		if t%1000 == 0 {
			print(".")
			_ = os.Stdout.Sync()
		}
		println()
	}

	log.Println("Compacting anthology ...")
	if err = anth.Compact(); err != nil {
		log.Fatal(err)
	}
	log.Println("done")

	if err := amaru.Save(); err != nil {
		panic(err)
	}

	elapsed := time.Since(t0)
	log.Printf("Total time was %v throughput was %.2f docs/s", elapsed.Truncate(time.Millisecond), float64(c)/elapsed.Seconds())
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
		likes := doc.Json["likes"].(float64)
		location := ""
		if doc.Json["loc"] != nil {
			location = doc.Json["loc"].(string)
		}

		doc := description + " " + location + " " + handle
		doc = text.RemoveBOM(doc)
		doc = text.RemoveEmojis(doc)
		doc = text.NormaliseFancyUnicodeToASCII(doc)
		doc = text.ReplaceStopWords(doc)

		sd := StemmedDoc{
			Url:     fmt.Sprintf("pof://%s", handle),
			Ranking: float32(likes),
			Stems:   text.Stems(doc),
		}
		stemChan <- sd
	}
}

func main() {
	ingest()
}
