package task

import (
	"fmt"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service"
	"github.com/tortuepin/tolog_ddd/pkg/service/search"
)

const TaskTagString = "#Task/list"

type TaskGenerator struct {
	logservice    service.LogServiceInterface
	searchservice search.SearchServiceInterface
}

func NewTaskGenerator(logservice service.LogServiceInterface, searchservice search.SearchServiceInterface) (*TaskGenerator, error) {
	return &TaskGenerator{logservice, searchservice}, nil
}

func (g *TaskGenerator) Generate(log model.Log) (model.Log, error) {
	// タグが`#Task/list`かどうかを判定する
	if !g.hasTaskListTag(log.Tags()) {
		return log, nil
	}
	// 直近の`#Task/list`をとってくる
	recent, err := g.fetchRecentTaskListLog()
	if err != nil {
		return model.Log{}, fmt.Errorf("error in TaskGenerator.Generate(): %w", err)
	}
	// contentをlogに代入する
	return model.NewLog(log.Time(), log.Tags(), recent.Content())
}

func (g *TaskGenerator) hasTaskListTag(tags []model.Tag) bool {
	for _, t := range tags {
		if t.Tag() == TaskTagString {
			return true
		}
	}
	return false
}

func (g *TaskGenerator) fetchRecentTaskListLog() (model.Log, error) {
	// TODO SearchByTagsのreturnの順序に依存しない形にする
	logs, err := g.logservice.ReadLogs()
	if err != nil {
		return model.Log{}, fmt.Errorf("error in TaskGenerator.fetchRecentTaskListLog(): %w", err)
	}
	query := search.NewTagQuery([]string{TaskTagString})
	searched, err := g.searchservice.Search(logs, query)
	if err != nil {
		return model.Log{}, fmt.Errorf("error in TaskGenerator.fetchRecentTaskListLog(): %w", err)
	}

	if len(searched) == 0 {
		return model.Log{}, nil
	}
	return searched[len(searched)-1], nil
}
