package task_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service"
	"github.com/tortuepin/tolog_ddd/pkg/service/generator/task"
	"github.com/tortuepin/tolog_ddd/pkg/service/search"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

type doNothingLogService struct {
	service.LogServiceInterface
}

func (l *doNothingLogService) ReadLogs() ([]model.Log, error) {
	return []model.Log{}, nil
}

type fakeSearchService struct {
	search.SearchServiceInterface
	fakeSearch func([]model.Log, search.Query) ([]model.Log, error)
}

func (f *fakeSearchService) Search(logs []model.Log, query search.Query) ([]model.Log, error) {
	return f.fakeSearch(logs, query)
}

func CreateFakeSearchServiceReturnLogs(logs []model.Log) *fakeSearchService {
	fakeSearch := func([]model.Log, search.Query) ([]model.Log, error) {
		return logs, nil
	}
	return &fakeSearchService{
		fakeSearch: fakeSearch,
	}
}

func TestTaskGenerator_Generate(t *testing.T) {
	type fields struct {
		logservice    service.LogServiceInterface
		searchservice search.SearchServiceInterface
	}
	type args struct {
		log model.Log
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Log
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				logservice: &doNothingLogService{},
				searchservice: CreateFakeSearchServiceReturnLogs(
					[]model.Log{
						testhelper.NewLogForTest(t,
							time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
							[]string{"@Task/list"},
							[]string{"tasklist1"},
						),
						testhelper.NewLogForTest(t,
							time.Date(2021, time.July, 10, 13, 34, 0, 0, time.UTC),
							[]string{"@Task/list"},
							[]string{"tasklist2"},
						),
					},
				),
			},
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 14, 34, 0, 0, time.UTC),
					[]string{"@Task/list"},
					[]string{""},
				),
			},
			want: testhelper.NewLogForTest(t,
				time.Date(2021, time.July, 10, 14, 34, 0, 0, time.UTC),
				[]string{"@Task/list"},
				[]string{"tasklist2"},
			),
		},
		{
			name: "When Tags does not have @Task/list",
			fields: fields{
				logservice: &doNothingLogService{},
				searchservice: CreateFakeSearchServiceReturnLogs(
					[]model.Log{
						testhelper.NewLogForTest(t,
							time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
							[]string{"@Task/list"},
							[]string{"tasklist1"},
						),
						testhelper.NewLogForTest(t,
							time.Date(2021, time.July, 10, 13, 34, 0, 0, time.UTC),
							[]string{"@Task/list"},
							[]string{"tasklist2"},
						),
					},
				),
			},
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 14, 34, 0, 0, time.UTC),
					[]string{"@invalidTag"},
					[]string{""},
				),
			},
			want: testhelper.NewLogForTest(t,
				time.Date(2021, time.July, 10, 14, 34, 0, 0, time.UTC),
				[]string{"@invalidTag"},
				[]string{""},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := task.NewTaskGenerator(tt.fields.logservice, tt.fields.searchservice)
			if err != nil {
				t.Errorf("task.NewTaskGenerator() error = %w", err)
			}
			got, err := g.Generate(tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaskGenerator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskGenerator.Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
