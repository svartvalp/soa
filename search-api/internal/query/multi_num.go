package query

import (
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/numeric"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/searcher"
	"github.com/blevesearch/bleve_index_api"
)

type MultiNumQuery struct {
	Vals  []float64
	Boost float64
	Field string
	Limit bool
}

func (q *MultiNumQuery) Searcher(i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	terms := make([][]byte, 0, len(q.Vals))
	for _, v := range q.Vals {
		terms = append(terms, numeric.MustNewPrefixCodedInt64(numeric.Float64ToInt64(v), 0))
	}
	field := m.DefaultSearchField()
	if q.Field != "" {
		field = q.Field
	}
	return searcher.NewMultiTermSearcherBytes(i, terms, field, q.Boost, options, q.Limit)
}

func NewMultiNum(vals []float64) *MultiNumQuery {
	return &MultiNumQuery{Vals: vals}
}
