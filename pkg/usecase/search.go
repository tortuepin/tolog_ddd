package usecase

import (
	"fmt"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service"
	"github.com/tortuepin/tolog_ddd/pkg/service/search"
)

func SearchByTags(logservice service.LogServiceInterface, searchservice search.SearchServiceInterface, tagstrings []string) ([]model.Log, error) {
	logs, err := logservice.ReadLogs()
	if err != nil {
		return []model.Log{}, fmt.Errorf("error in usecase.SearchByTags(): %w", err)
	}
	query := search.NewTagQuery(tagstrings)
	searched, err := searchservice.Search(logs, query)
	if err != nil {
		return []model.Log{}, fmt.Errorf("error in usecase.SearchByTags(): %w", err)
	}
	return searched, nil
}
