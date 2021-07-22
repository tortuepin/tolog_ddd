package file

import (
	"fmt"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

type Parser interface {
	Parse([]string) ([]ParseReturn, error)
}

type Formatter interface {
	Format(model.Log) []string
}

type ParseReturn struct {
	logTimePart logTimePart
	tags        []model.Tag
	content     model.LogContent
}

func NewParseReturn(timepart time.Time, tags []model.Tag, content model.LogContent) ParseReturn {
	return ParseReturn{logTimePart{timepart}, tags, content}
}

func (p ParseReturn) ToLog(year int, month time.Month, day int) (model.Log, error) {
	logtime, err := p.LogTime(year, month, day)
	if err != nil {
		return model.Log{}, fmt.Errorf("error in ToLog: %w", err)
	}
	return model.NewLog(logtime, p.Tags(), p.LogContent())
}

func (p ParseReturn) LogTime(year int, month time.Month, day int) (model.LogTime, error) {
	return p.logTimePart.toLogTime(year, month, day)
}

func (p ParseReturn) Tags() []model.Tag {
	return p.tags
}

func (p ParseReturn) LogContent() model.LogContent {
	return p.content
}

type logTimePart struct {
	t time.Time
}

func (l logTimePart) toLogTime(year int, month time.Month, day int) (model.LogTime, error) {
	ti := time.Date(year, month, day, l.t.Hour(), l.t.Minute(), l.t.Second(), l.t.Nanosecond(), l.t.Location())
	return model.NewLogTime(ti)
}
