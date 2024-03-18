package main

import (
	"bufio"
	"fmt"
	"github.com/kukino/Amaru"
	"github.com/kukino/Amaru/impl"
	"github.com/kukino/Amaru/text"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	amaru := openAmaru()

	fmt.Printf("Index: %s\n", amaru.Path())
	fmt.Printf("Stats:\n")
	fmt.Printf(" - %d Documents\n", amaru.Documents().Count())
	fmt.Printf(" - %d Tokens\n", amaru.Tokens().Count())
	fmt.Printf("\n")
	fmt.Printf("Commands:\n")
	fmt.Printf(" - a free text query\n")
	fmt.Printf(" - !tokens 	(list all the tokens)\n")
	fmt.Printf("\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Command: ")
		cmd := readLine(reader)
		if cmd == "!tokens" {
			for t := 0; t < amaru.Tokens().Count(); t++ {
				token := amaru.Tokens().Get(Amaru.TokenID(t))
				tt := "???"
				if token.Type == Amaru.TextToken {
					tt = "TXT"
				} else if token.Type == Amaru.TagToken {
					tt = "TAG"
				}
				fmt.Printf("[%05d] %s %s\n", t, tt, token.Text)

				if t%50 == 49 {
					fmt.Printf("Continue?")
					cmd := readLine(reader)
					if cmd == "n" || cmd == "N" {
						break
					}
				}
			}
		} else {
			t0 := time.Now()
			tids := amaru.Tokens().GetIds(Amaru.TextToken, text.Stems(cmd))
			docids := amaru.Anthology().FindDocIDsWith(tids, 5_000)
			docs := amaru.Documents().GetAll(docids)
			elapsed := time.Since(t0)
			log.Printf("Search took %v for %d results.\n", elapsed, len(docs))

			for n, doc := range docs {
				fmt.Printf("%v\t", doc.URL)
				if n > 100 {
					break
				}
			}
			fmt.Println()
		}

	}
}

func openAmaru() Amaru.Amaru {
	basePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	amaru, err := impl.NewAmaru(path.Join(basePath, "/KUKINO/GO/Amaru/tmp/idx1/"), false)
	if err != nil {
		panic(err)
	}
	if err := amaru.Load(); err != nil {
		panic(err)
	}
	return amaru
}

func readLine(reader *bufio.Reader) string {
	cmd, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Trim(cmd, "\t\n ")
}
