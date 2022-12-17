package query

import (
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/blevesearch/bleve/v2/numeric"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/blevesearch/bleve/v2/search/searcher"
	index "github.com/blevesearch/bleve_index_api"
)

type NumQuery struct {
	Val   float64
	Boost float64
	Field string
	Limit bool
}

func (q *NumQuery) Searcher(i index.IndexReader, m mapping.IndexMapping, options search.SearcherOptions) (search.Searcher, error) {
	terms := make([][]byte, 0, 1)
	terms = append(terms, numeric.MustNewPrefixCodedInt64(numeric.Float64ToInt64(q.Val), 0))
	field := m.DefaultSearchField()
	if q.Field != "" {
		field = q.Field
	}
	return searcher.NewMultiTermSearcherBytes(i, terms, field, q.Boost, options, q.Limit)
}

func NewNum(val float64) *NumQuery {
	return &NumQuery{Val: val}
}
