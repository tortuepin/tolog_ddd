package extractor_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service/extractor"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

type fakeQuery struct {
	extractor.Query
	fakeSatisfy func(model.Log) bool
}

func (f *fakeQuery) Satisfy(log model.Log) bool {
	return f.fakeSatisfy(log)
}

func TestDefaultExtractor_Extract(t *testing.T) {
	type arg struct {
		logs  []model.Log
		query extractor.Query
	}
	tests := []struct {
		name    string
		arg     arg
		want    []model.Log
		wantErr bool
	}{
		{
			name: "",
			arg: arg{
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
						if log.Tags()[0].Tag() == "@tag1" {
							return true
						}
						return false
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
			sut, _ := extractor.NewDefaultExtractor()
			got, err := sut.Extract(tt.arg.logs, tt.arg.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractor.Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractor.Extract(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}
