package model

import "time"

func NewLogTimeForTest(time time.Time) LogTime {
	return LogTime{time}
}

func NewLogForTest(time LogTime, tags []Tag, content LogContent) Log {
	return Log{
		time,
		tags,
		content,
	}
}

func NewTagForTest(tag string) Tag {
	return Tag{tag}
}

func NewLogContentForTest(content []string) LogContent {
	return LogContent{content}
}
