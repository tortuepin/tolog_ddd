package service

import (
	"fmt"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/domain/repository"
)

type LogServiceInterface interface {
	NewLog([]model.Tag, model.LogContent) (model.Log, error)
	EditLog(model.Log, model.Log) error
	ReadLogs() ([]model.Log, error)
}

type LogService struct {
	reader  repository.Reader
	creater repository.Creater
	updater repository.Updater
}

func NewLogService(reader repository.Reader, creater repository.Creater, updater repository.Updater) (*LogService, error) {
	return &LogService{reader, creater, updater}, nil
}

func (s *LogService) NewLog(tags []model.Tag, content model.LogContent) (model.Log, error) {
	t, err := model.NewLogTimeNow()
	if err != nil {
		return model.Log{}, fmt.Errorf("failed in NewLogTimeNow(): %w", err)
	}

	log, err := model.NewLog(t, tags, content)
	if err != nil {
		return model.Log{}, fmt.Errorf("failed in NewLog(): %w", err)
	}

	if err := s.creater.Create(log); err != nil {
		return model.Log{}, fmt.Errorf("failed in Write(): %w", err)
	}

	return log, nil
}

func (s *LogService) EditLog(from model.Log, to model.Log) error {
	err := s.updater.Update(from, to)
	if err != nil {
		return fmt.Errorf("failed in Updater(): %w", err)
	}
	return nil
}

func (s *LogService) ReadLogs() ([]model.Log, error) {
	return s.reader.Read()
}
