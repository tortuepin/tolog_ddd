package usecase

import "fmt"
import "github.com/tortuepin/tolog_ddd/pkg/service"
import "github.com/tortuepin/tolog_ddd/pkg/domain/model"

func LogNew(service service.LogService, tagstrings []string) error {
	tags, err := model.NewTags(tagstrings)
	if err != nil {
		return fmt.Errorf("error in usecase.LogNew(): %w", err)
	}

	content := model.LogContent{}

	if err := service.NewLog(tags, content); err != nil {
		return fmt.Errorf("error in usecase.LogNew(): %w", err)
	}
	return nil
}
