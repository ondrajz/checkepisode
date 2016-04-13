package main

import (
	"flag"
	"log"

	"github.com/garfunkel/go-tvdb"
)

func main() {
	var (
		query = flag.String("q", "new girl", "search query")
	)
	flag.Parse()
	log.Printf("searching for '%s'", *query)
	list, err := tvdb.SearchSeries(*query, 5)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d results", len(list.Series))
	for i, s := range list.Series {
		log.Printf(" %d. %v", i+1, s.SeriesName)
		log.Printf("%+v", s)
	}
}
