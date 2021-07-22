package search_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service/search"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

type fakeExtractor struct {
	search.Extractor
	fakeExtract func([]model.Log, search.Query) ([]model.Log, error)
}

func (f *fakeExtractor) Extract(logs []model.Log, query search.Query) ([]model.Log, error) {
	return f.fakeExtract(logs, query)
}

func TestSearchService_Search(t *testing.T) {
	type args struct {
		logs  []model.Log
		query search.Query
	}
	type fields struct {
		extractor search.Extractor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Log
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				extractor: &fakeExtractor{
					fakeExtract: func(logs []model.Log, query search.Query) ([]model.Log, error) {
						ex, _ := search.NewDefaultExtractor()
						return ex.Extract(logs, query)
					},
				},
			},
			args: args{
				logs: []model.Log{
					testhelper.NewLogForTest(
						t,
						time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
						[]string{"@tag1"},
						[]string{"", "content1"},
					),
					testhelper.NewLogForTest(
						t,
						time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
						[]string{"@tag2"},
						[]string{"", "content1"},
					),
				},
				query: &fakeQuery{
					fakeSatisfy: func(log model.Log) bool {
						q := search.NewTagQuery([]string{"@tag1"})
						return q.Satisfy(log)
					},
				},
			},
			want: []model.Log{
				testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{"@tag1"},
					[]string{"", "content1"},
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := search.NewSearchService(tt.fields.extractor)
			got, err := sut.Search(tt.args.logs, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchService.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchService.Search(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}
