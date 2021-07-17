package repository

import (
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

func NewLogForTest(time model.LogTime, tags []model.Tag, content model.LogContent) model.Log {
	log, _ := model.NewLog(time, tags, content)
	return log
}

func NewLogTimeForTest(time time.Time) model.LogTime {
	ret, _ := model.NewLogTime(time)
	return ret
}

func NewTagForTest(tag string) model.Tag {
	ret, _ := model.NewTag(tag)
	return ret
}

func NewLogContentForTest(content []string) model.LogContent {
	ret, _ := model.NewLogContent(content)
	return ret
}

func (m *MarkdownParser) ParseLinesForTest(lines []string) ([][]string, error) {
	return m.parseLines(lines)
}

func (m *MarkdownParser) ParseLogForTest(lines []string) (ParseReturn, error) {
	return m.parseLog(lines)
}

func NewParseReturnForTest(time time.Time, tags []model.Tag, content model.LogContent) ParseReturn {
	return ParseReturn{logTimePart{time}, tags, content}
}
