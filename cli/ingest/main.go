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
)

type JSON map[string]interface{}
type Record struct {
	Path string
	Json JSON
}

func main() {
	println("Indexing ...")

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

	feed := make(chan Record)
	go readJsons(path.Join(basePath, "/KUKINO/GO/Amaru/tmp/dataset/profiles/"), feed)

	c := 0
	for jsonData := range feed {

		//fmt.Println("-----")
		//fmt.Println(jsonData.Path)

		handle := jsonData.Json["handle"].(string)
		description := jsonData.Json["desc.txt"].(string)
		likes := jsonData.Json["likes"].(float64)
		location := ""
		if jsonData.Json["loc"] != nil {
			location = jsonData.Json["loc"].(string)
		}

		did := docs.Add(fmt.Sprintf("pof://%s", handle), float32(likes))

		doc := description + " " + location + " " + handle
		doc = text.RemoveBOM(doc)
		doc = text.RemoveEmojis(doc)
		doc = text.NormaliseFancyUnicodeToASCII(doc)
		doc = text.ReplaceStopWords(doc)

		stems := map[string]Amaru.TokenID{} // for displaying only
		var tids []Amaru.TokenID
		for _, stem := range text.Stems(doc) {
			tid, stemmed := tokens.Add(Amaru.TextToken, stem)
			stems[stemmed] = tid // for displaying only
			tids = append(tids, tid)
		}

		// tokenIds must be added in order so the anthology can be Compacted
		sort.Slice(tids, func(i, j int) bool {
			return tids[i] < tids[j]
		})

		for _, tid := range tids {
			anth.Add(did, tid)
		}

		if c%1000 == 0 {
			fmt.Printf("D%d\t", did)
			for token, tid := range stems {
				fmt.Printf("%d:%s\t", tid, token)
			}
			fmt.Println(len(stems))
		}

		c++
		if c%10000 == 0 {
			println(c, "...")
			if err := amaru.Save(); err != nil {
				panic(err)
			}
			println("saved")
		}
		if c > 10_000 {
			break
		}
	}

	println("Sorting all Dossiers, not really necessary but ... ")
	println("Dossiers (one per Token):", tokens.Count())
	for t := 0; t < tokens.Count(); t++ {
		anth.GetDossier(Amaru.TokenID(t)).Sort()
		if t%1000 == 0 {
			print(".")
			_ = os.Stdout.Sync()
		}
	}
	println()

	println("Compacting anthology ...")
	if err = anth.Compact(); err != nil {
		log.Fatal(err)
	}
	println("done")

	if err := amaru.Save(); err != nil {
		panic(err)
	}

}

func readJsons(basePath string, ch chan Record) {
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

			ch <- Record{Path: path, Json: jsonData}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through files:", err)
	}
}
