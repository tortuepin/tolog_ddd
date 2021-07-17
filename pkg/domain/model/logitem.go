package model

import "time"

type LogItem struct {
	time    logTime
	tags    tags
	content logContent
	id      logID
}

type logContent []string

type tags []tag
type tag struct {
	tag string
}

type logTime time.Time

type logID int

type IDGenerator func(logTime) logId

func NewLogItem(time time.Time, tags []string, content []string, dup DuplicateCount) LogItem {
	t, err := newLogTime(time)
	if err != nil {
		return LogItem{}, err
	}

	tag, err := newTags(tags)
	if err != nil {
		return LogItem{}, err
	}

	c, err := newContent(content)
	if err != nil {
		return LogItem{}, err
	}

	id := DuplicateCount(t) + 1

	return LogItem{
		time:    t,
		tags:    tag,
		content: c,
		id:      id}
}

func newLogTime(t time.Time) (logTime, error) {
	return logTime{t}, nil
}

func newTags(tags []string) (tags, error) {
	ret := []tag{}
	for i := 0; i < len(tags); i++ {
		t, err := newTag()
		if err != nil {
			return tags{}, err
		}
		ret = append(ret, t)
	}
	return t, nil
}

func newContent(c []string) (logContent, error) {
	return logContent{c}, nil
}
