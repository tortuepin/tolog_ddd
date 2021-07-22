package extractor

import "github.com/tortuepin/tolog_ddd/pkg/domain/model"

type Query interface {
	Satisfy(model.Log) bool
}

type TagQuery struct {
	targettags []string
}

func NewTagQuery(tags []string) *TagQuery {
	return &TagQuery{tags}
}

func (q *TagQuery) Satisfy(log model.Log) bool {
	for _, l := range log.Tags() {
		if q.contain(l.Tag()) {
			return true
		}
	}
	return false
}

func (q *TagQuery) contain(tag string) bool {
	for _, t := range q.targettags {
		if tag == t {
			return true
		}
	}
	return false
}
