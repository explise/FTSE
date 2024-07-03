package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	var dumpPath string
	flag.StringVar(&dumpPath, "p", "enwiki-20240301-abstract1.xml", "wiki abstract dump path")
	flag.Parse()

	log.Println("Starting simplefts")

	start := time.Now()
	docs, err := loadDocuments(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(index)
	idx.add(docs)
	save(idx)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	for true {
		start = time.Now()
		var qu string
		fmt.Scanln(&qu)
		if qu == "exit" {
			break
		}
		matchedIDs := idx.search(qu)
		log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

		for _, id := range matchedIDs {
			doc := docs[id]
			log.Printf("%d\t%s\n", id, doc.Text)
		}
	}
}
