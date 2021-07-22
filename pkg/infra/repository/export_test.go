package repository

import (
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/domain/repository/file"
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

func (m *MarkdownParser) ParseLogForTest(lines []string) (file.ParseReturn, error) {
	return m.parseLog(lines)
}

func NewFilenameForTest(date string, ext string) Filename {
	return Filename{date: date, ext: ext}
}
