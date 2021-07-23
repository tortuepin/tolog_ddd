package format_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/infra/repository/file/format"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

func TestSearchFormatter_Format(t *testing.T) {

	type args struct {
		log model.Log
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				log: testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{"@tag"},
					[]string{"", "content"},
				),
			},
			want: []string{
				"[\\210710 12:34] @tag",
				"",
				"content",
			},
		},
		{
			name: "",
			args: args{
				log: testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{},
					[]string{"", "content"},
				),
			},
			want: []string{
				"[\\210710 12:34]",
				"",
				"content",
			},
		},
		{
			name: "",
			args: args{
				log: testhelper.NewLogForTest(
					t,
					time.Date(2021, time.July, 10, 12, 34, 56, 0, time.UTC),
					[]string{},
					[]string{"", "content"},
				),
			},
			want: []string{
				"[\\210710 12:34:56]",
				"",
				"content",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := format.NewSearchFormatter()
			got := sut.Format(tt.args.log)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractorFormatter.Format(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}
