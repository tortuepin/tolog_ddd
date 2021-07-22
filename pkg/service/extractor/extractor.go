package extractor

import "github.com/tortuepin/tolog_ddd/pkg/domain/model"

type Extractor interface {
	Extract([]model.Log, Query) ([]model.Log, error)
}

func NewDefaultExtractor() (*DefaultExtractor, error) {
	return &DefaultExtractor{}, nil
}

type DefaultExtractor struct {
}

func (e *DefaultExtractor) Extract(logs []model.Log, query Query) ([]model.Log, error) {
	ret := []model.Log{}
	for _, l := range logs {
		if query.Satisfy(l) {
			ret = append(ret, l)
		}
	}
	return ret, nil
}
