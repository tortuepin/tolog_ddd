package search

import (
	"strings"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

type SearchFormatter struct {
}

func NewSearchFormatter() *SearchFormatter { return &SearchFormatter{} }

func (f *SearchFormatter) Format(log model.Log) []string {
	ret := []string{}
	timeline := "[" + f.formatTime(log.Time()) + "]"
	tagline := " " + f.formatTags(log.Tags())
	firstline := strings.TrimSpace(timeline + tagline)
	ret = append(ret, firstline)
	ret = append(ret, f.formatContent(log.Content())...)
	return ret
}

func (f *SearchFormatter) formatTime(t model.LogTime) string {
	if t.Time().Second() == 0 {
		return t.Time().Format("\\060102 15:04")
	}
	return t.Time().Format("\\060102 15:04:05")
}

func (f *SearchFormatter) formatTags(tags []model.Tag) string {
	tagsstr := ""
	for _, t := range tags {
		tagsstr = tagsstr + " " + t.Tag()
	}
	return strings.TrimSpace(tagsstr)
}
func (f *SearchFormatter) formatContent(content model.LogContent) []string {
	return content.Content()
}
