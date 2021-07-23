package testhelper

import (
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

func NewLogForTest(t *testing.T, ti time.Time, tagsstr []string, contentstr []string) model.Log {
	t.Helper()
	lt, err := model.NewLogTime(ti)
	if err != nil {
		t.Fatal(err)
	}
	tags, err := model.NewTags(tagsstr)
	if err != nil {
		t.Fatal(err)
	}
	con, err := model.NewLogContent(contentstr)
	if err != nil {
		t.Fatal(err)
	}
	l, err := model.NewLog(lt, tags, con)
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func NewLogTimeForTest(t *testing.T, time time.Time) model.LogTime {
	t.Helper()
	ret, _ := model.NewLogTime(time)
	return ret
}
func NewTagForTest(t *testing.T, tag string) model.Tag {
	t.Helper()
	ret, _ := model.NewTag(tag)
	return ret
}

func NewLogContentForTest(t *testing.T, content []string) model.LogContent {
	t.Helper()
	ret, _ := model.NewLogContent(content)
	return ret
}
