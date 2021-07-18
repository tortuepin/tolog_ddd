package service

import (
	"fmt"
	"reflect"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

func EqualWithoutTime(got model.Log, want model.Log) error {
	dummytime, _ := model.NewLogTime(time.Time(time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC)))
	g, _ := model.NewLog(dummytime, got.Tags(), got.Content())
	w, _ := model.NewLog(dummytime, want.Tags(), want.Content())

	if !reflect.DeepEqual(g, w) {
		return fmt.Errorf("got = %v, want %v", g, w)
	}
	return nil
}

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
