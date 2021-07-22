package search

import (
	"fmt"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

type SearchServiceInterface interface {
	Search([]model.Log, Query) ([]model.Log, error)
}

type SearchService struct {
	extractor Extractor
}

func NewSearchService(extractor Extractor) *SearchService {
	return &SearchService{extractor: extractor}
}

func (ss *SearchService) Search(logs []model.Log, query Query) ([]model.Log, error) {
	ret, err := ss.extractor.Extract(logs, query)
	if err != nil {
		return []model.Log{}, fmt.Errorf("error in SearchService.Search(): %w", err)
	}
	return ret, nil
}
