package usecase

import "fmt"
import "github.com/tortuepin/tolog_ddd/pkg/service"
import "github.com/tortuepin/tolog_ddd/pkg/domain/model"

func LogNew(s service.LogServiceInterface, tagstrings []string) (model.Log, error) {
	tags, err := model.NewTags(tagstrings)
	if err != nil {
		return model.Log{}, fmt.Errorf("error in usecase.LogNew(): %w", err)
	}

	content := model.LogContent{}

	log, err := s.NewLog(tags, content)
	if err != nil {
		return model.Log{}, fmt.Errorf("error in usecase.LogNew(): %w", err)
	}
	return log, nil
}
