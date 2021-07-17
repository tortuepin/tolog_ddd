package model

import (
	"time"
)

type Log struct {
	time    LogTime
	tags    []Tag
	content LogContent
}

func NewLog(time LogTime, tags []Tag, content LogContent) (Log, error) {

	return Log{
		time:    time,
		tags:    tags,
		content: content}, nil
}

func (l Log) Time() LogTime {
	return l.time
}
func (l Log) Tags() []Tag {
	return l.tags
}
func (l Log) Content() LogContent {
	return l.content
}

type LogTime struct {
	time time.Time
}

func (l LogTime) Time() time.Time {
	return l.time
}

func truncate(t time.Time) time.Time {
	return t.Truncate(time.Nanosecond).Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)
}

func NewLogTime(t time.Time) (LogTime, error) {
	return LogTime{truncate(t)}, nil
}

func NewLogTimeNow() (LogTime, error) {
	now := time.Now()
	return NewLogTime(now)
}

type Tag struct {
	tag string
}

func (t Tag) Tag() string {
	return t.tag
}

func NewTag(t string) (Tag, error) {
	return Tag{t}, nil
}

func NewTags(tags []string) ([]Tag, error) {
	ret := []Tag{}
	for i := 0; i < len(tags); i++ {
		t, err := NewTag(tags[i])
		if err != nil {
			return []Tag{}, err
		}
		ret = append(ret, t)
	}
	return ret, nil
}

type LogContent struct {
	content []string
}

func (l LogContent) Content() []string {
	return l.content
}

func NewLogContent(c []string) (LogContent, error) {
	return LogContent{c}, nil
}
