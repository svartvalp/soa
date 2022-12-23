package search_service

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/blevesearch/bleve/v2"
	"github.com/soa/search-api/internal/models"
	"github.com/soa/search-api/internal/pkg/search-service/dto"
	q "github.com/soa/search-api/internal/query"
)

type Service struct {
	main bleve.IndexAlias
	old  bleve.Index
	new  bleve.Index
}

func NewService() (*Service, error) {
	meta, err := getMeta()
	if err != nil {
		return nil, err
	}
	index := meta.Index
	if index == "" {
		index = "current"
	}
	old, err := bleve.Open(index)
	if err != nil {
		if err == bleve.ErrorIndexPathDoesNotExist {
			old, err = bleve.New(index, bleve.NewIndexMapping())
			if err != nil {
				return nil, err
			}
		}
	}
	if index != meta.Index {
		err = setMeta(&dto.Meta{Index: index})
		if err != nil {
			return nil, err
		}
	}
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	n, err := bleve.New(id, bleve.NewIndexMapping())
	alias := bleve.NewIndexAlias(old)
	return &Service{
		main: alias,
		old:  old,
		new:  n,
	}, nil
}

func getMeta() (*dto.Meta, error) {
	f, err := os.Open("meta.json")
	if err != nil {
		m := dto.Meta{}
		var dat []byte
		dat, err = json.Marshal(&m)
		if err != nil {
			return nil, err
		}
		err = os.WriteFile("meta.json", dat, os.ModePerm)
		if err != nil {
			return nil, err
		}
		return &m, nil
	}
	dat, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var m dto.Meta
	err = json.Unmarshal(dat, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func setMeta(m *dto.Meta) error {
	var dat []byte
	var err error
	dat, err = json.Marshal(&m)
	if err != nil {
		return err
	}
	err = os.WriteFile("meta.json", dat, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) List(ctx context.Context, req *models.Filter) ([]int64, error) {
	query := bleve.NewConjunctionQuery()
	if req.Query != "" {
		query.AddQuery(bleve.NewFuzzyQuery(req.Query))
	}
	if req.CatID != 0 {
		catQ := q.NewNum(float64(req.CatID))
		catQ.Field = "category_ids"
		query.AddQuery(catQ)
	}
	if req.Brand != "" {
		brandQ := bleve.NewTermQuery(req.Brand)
		brandQ.FieldVal = "brand"
		query.AddQuery(brandQ)
	}
	if req.PriceFrom != 0 || req.PriceTo != 0 {
		var from, to *float64
		if req.PriceFrom != 0 {
			from = new(float64)
			*from = float64(req.PriceFrom)
		}
		if req.PriceTo != 0 {
			to = new(float64)
			*to = float64(req.PriceTo)
		}
		var minIn, maxIn *bool
		minIn = new(bool)
		maxIn = new(bool)
		*minIn = true
		*maxIn = true
		priceQ := bleve.NewNumericRangeInclusiveQuery(from, to, minIn, maxIn)
		priceQ.FieldVal = "price"
		query.AddQuery(priceQ)
	}
	searchReq := bleve.NewSearchRequest(query)
	if len(query.Conjuncts) == 0 {
		searchReq = bleve.NewSearchRequest(bleve.NewMatchAllQuery())
	}
	searchReq.Fields = []string{"*"}
	res, err := s.main.SearchInContext(ctx, searchReq)
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, res.Hits.Len())
	for _, h := range res.Hits {
		id, err := strconv.ParseInt(h.ID, 10, 64)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (s *Service) Regenerate(ctx context.Context, products []models.ProductInfo) error {
	var err error
	dtos := mapPrInfosToDtos(products)

	for _, d := range dtos {
		err = s.new.Index(d.ID, d)
		if err != nil {
			return err
		}
	}
	return nil
	// return s.SwapIndex(ctx)
}

func (s *Service) SwapIndex(ctx context.Context) error {
	var err error
	s.main.Swap([]bleve.Index{s.new}, []bleve.Index{s.old})
	s.old = s.new
	indexName := s.old.Name()
	err = setMeta(&dto.Meta{Index: indexName})
	if err != nil {
		return err
	}
	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	err = s.old.(bleve.IndexCopyable).CopyTo(bleve.FileSystemDirectory(id))
	if err != nil {
		return err
	}
	s.new, err = bleve.Open(id)
	if err != nil {
		return err
	}
	return nil
}

func mapPrInfosToDtos(products []models.ProductInfo) []dto.ProductInfo {
	res := make([]dto.ProductInfo, 0, len(products))
	for _, pr := range products {
		catIDs := make([]float64, 0, len(pr.Categorys))
		cats := make([]dto.Category, 0, len(pr.Categorys))
		for _, cat := range pr.Categorys {
			catIDs = append(catIDs, float64(cat.ID))
			cats = append(cats, dto.Category{
				ID:          float64(cat.ID),
				Name:        cat.Name,
				Description: cat.Description,
				ParentID:    float64(cat.ParentID),
				Level:       float64(cat.Level),
			})
		}
		res = append(res, dto.ProductInfo{
			ID:              strconv.FormatInt(pr.ID, 10),
			CategoryID:      float64(pr.CategoryID),
			Price:           float64(pr.Price),
			Name:            pr.Name,
			Description:     pr.Description,
			Brand:           pr.Brand,
			Image:           pr.Image,
			CategoryIDs:     catIDs,
			Characteristics: mapCharsFromPrInfo(pr.Characteristics),
			Categorys:       cats,
		})
	}
	return res
}

func mapCharsFromPrInfo(characteristics []models.ProductCharacteristic) []dto.ProductCharacteristic {
	res := make([]dto.ProductCharacteristic, 0, len(characteristics))
	for _, ch := range characteristics {
		res = append(res, dto.ProductCharacteristic{
			ProductId:   float64(ch.ProductId),
			ID:          float64(ch.ID),
			Name:        ch.Name,
			ChType:      ch.ChType,
			Description: ch.Description,
		})
	}
	return res
}
