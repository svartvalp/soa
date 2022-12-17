package main

import (
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/numeric"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/query"
	"github.com/blevesearch/bleve/v2/search/searcher"
	"github.com/blevesearch/bleve_index_api"
)

type Query struct {
	vals []float64
}

func (q *Query) Searcher(i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	terms := make([][]byte, 0, len(q.vals))
	for _, v := range q.vals {
		terms = append(terms, numeric.MustNewPrefixCodedInt64(numeric.Float64ToInt64(v), 0))
	}
	return searcher.NewMultiTermSearcherBytes(i, terms, m.DefaultSearchField(), 0, options, false)
}

func New(vals []float64) query.Query {
	return &Query{vals: vals}
}
