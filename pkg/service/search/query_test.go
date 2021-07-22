package search_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/service/search"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

func TestTagQueryBuilder_Build(t *testing.T) {
	type args struct {
		queryString string
	}
	tests := []struct {
		name    string
		args    args
		want    search.Query
		wantErr bool
	}{
		{
			name: "",
			args: args{
				queryString: "@tag1",
			},
			want: search.NewTagQuery([]string{"@tag1"}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := search.NewTagQueryBuilder()
			got, err := sut.Build(tt.args.queryString)
			if (err != nil) != tt.wantErr {
				t.Errorf("TagQueryBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagQueryBuilder.Build(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestTagQuery_Satisfy(t *testing.T) {
	type args struct {
		log model.Log
	}
	tests := []struct {
		name string
		sut  *search.TagQuery
		args args
		want bool
	}{
		{
			name: "1番目のタグにマッチする場合",
			sut: search.NewTagQuery([]string{
				"@tag1",
				"@tag2",
			}),
			args: args{
				log: testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{"@tag1"},
					[]string{"", "content1"},
				),
			},
			want: true,
		},
		{
			name: "2番目のタグにマッチする場合",
			sut: search.NewTagQuery([]string{
				"@tag1",
				"@tag2",
			}),
			args: args{
				log: testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{"@tag2"},
					[]string{"", "content1"},
				),
			},
			want: true,
		},
		{
			name: "どのタグにもマッチしない場合",
			sut: search.NewTagQuery([]string{
				"@tag1",
				"@tag2",
			}),
			args: args{
				log: testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{"@tag3"},
					[]string{"", "content1"},
				),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sut.Satisfy(tt.args.log)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagQuery.Satisfy(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}
