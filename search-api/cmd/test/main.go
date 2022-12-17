package main

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
)

func main() {
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		panic(err)
	}

	// index some data
	err = index.Index("1", Data{
		A:   "1",
		B:   2,
		Arr: []float64{1, 2, 3},
	})
	if err != nil {
		panic(err)
	}
	err = index.Index("2", Data{
		A:   "text",
		B:   2,
		Arr: []float64{3, 4, 5},
	})
	_ = index.Index("3", Data{
		A:   "alala",
		B:   123,
		Arr: []float64{3, 555},
	})

	// // search for some text
	// a := 2.5
	// b := 3.6
	// d, _ := index.Fields()
	// fmt.Printf("%+v", d)
	query := New([]float64{555.0})
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"a", "b", "arr"}
	searchResults, err := index.Search(search)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", searchResults.Hits[0].Fields)
}
