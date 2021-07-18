package service_test

import (
	"reflect"
	"testing"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/domain/repository"
	"github.com/tortuepin/tolog_ddd/pkg/service"
)

type fakeReader struct {
	repository.Reader
	fakeRead func() ([]model.Log, error)
}

func (f *fakeReader) Read() ([]model.Log, error) {
	return f.fakeRead()
}

type fakeCreater struct {
	repository.Creater
	fakeCreate func(model.Log) error
}

func (f *fakeCreater) Create(log model.Log) error {
	return f.fakeCreate(log)
}

type fakeUpdater struct {
	repository.Updater
	fakeUpdater func(model.Log, model.Log) error
}

func (f *fakeUpdater) Update(from model.Log, to model.Log) error {
	return f.fakeUpdater(from, to)
}

func Test_NewLog(t *testing.T) {
	type args struct {
		tags    []model.Tag
		content model.LogContent
	}
	type fields struct {
		creater repository.Creater
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		want    model.Log
		wantErr bool
	}{
		{
			name: "",
			args: args{
				tags:    []model.Tag{service.NewTagForTest("@tag1")},
				content: service.NewLogContentForTest([]string{"content"}),
			},
			fields: fields{
				creater: &fakeCreater{
					fakeCreate: func(model.Log) error {
						return nil
					},
				},
			},
			want: service.NewLogForTest(
				model.LogTime{},
				[]model.Tag{
					service.NewTagForTest("@tag1"),
				},
				service.NewLogContentForTest([]string{"content"}),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := service.NewLogService(nil, tt.fields.creater, nil)
			got, err := s.NewLog(tt.args.tags, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.NewLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := service.EqualWithoutTime(got, tt.want); err != nil {
				t.Errorf("service.NewLog(): %v", err)
			}
		})
	}
}

func Test_EditLog(t *testing.T) {
	type args struct {
		from model.Log
		to   model.Log
	}
	type fields struct {
		updater repository.Updater
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		want   error
	}{
		{
			name: "",
			args: args{
				model.Log{},
				model.Log{},
			},
			fields: fields{
				updater: &fakeUpdater{
					fakeUpdater: func(model.Log, model.Log) error {
						return nil
					},
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := service.NewLogService(nil, nil, tt.fields.updater)
			got := s.EditLog(tt.args.from, tt.args.to)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.EditLog() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_ReadLogs(t *testing.T) {
	type fields struct {
		reader repository.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Log
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				reader: &fakeReader{
					fakeRead: func() ([]model.Log, error) {
						return []model.Log{}, nil
					},
				},
			},
			want: []model.Log{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := service.NewLogService(tt.fields.reader, nil, nil)
			got, err := s.ReadLogs()
			if (err != nil) != tt.wantErr {
				t.Errorf("model.NewLogContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.EditLog() got = %v, want %v", got, tt.want)
			}
		})
	}

}
