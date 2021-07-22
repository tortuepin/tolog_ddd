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
