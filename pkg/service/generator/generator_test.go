package generator_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service/generator"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

func addTag(l model.Log, tag string) model.Log {
	t, _ := model.NewTag(tag)
	m, _ := model.NewLog(l.Time(), append(l.Tags(), t), l.Content())
	return m
}

func TestMultiLogGenerator_Generate(t *testing.T) {
	type fields struct {
		generators []func(model.Log) (model.Log, error)
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
			name: "tagを追加するgenerator",
			fields: fields{
				generators: []func(model.Log) (model.Log, error){
					func(l model.Log) (model.Log, error) {
						return addTag(l, "@tag1"), nil
					},
					func(l model.Log) (model.Log, error) {
						return addTag(l, "@tag2"), nil
					},
				},
			},
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{},
					[]string{"", "content1", "content2"},
				),
			},
			want: testhelper.NewLogForTest(t,
				time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
				[]string{"@tag1", "@tag2"},
				[]string{"", "content1", "content2"},
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := generator.NewMultiLogGenerator(tt.fields.generators)
			if err != nil {
				t.Errorf("service.NewMultiLogGenerator() error = %w", err)
			}
			got, err := g.Generate(tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiLogGenerator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiLogGenerator.Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
