package format_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/domain/repository/file"
	"github.com/tortuepin/tolog_ddd/pkg/infra/repository/file/format"
	"github.com/tortuepin/tolog_ddd/pkg/testhelper"
)

func Test_Parse(t *testing.T) {

	type args struct {
		lines []string
	}

	tests := []struct {
		name    string
		args    args
		want    []file.ParseReturn
		wantErr bool
	}{
		{
			name: "",
			args: args{
				lines: []string{
					"[12:34] @tag",
					"",
					"content1",
					"content2",
					"",
					"[12:34:56]",
					"",
					"content1",
					"content2",
				},
			},
			want: []file.ParseReturn{
				file.NewParseReturn(
					time.Date(0, time.January, 1, 12, 34, 0, 0, time.UTC),
					[]model.Tag{testhelper.NewTagForTest(t, "@tag")},
					testhelper.NewLogContentForTest(t, []string{
						"",
						"content1",
						"content2",
						""}),
				),
				file.NewParseReturn(
					time.Date(0, time.January, 1, 12, 34, 56, 0, time.UTC),
					[]model.Tag{},
					testhelper.NewLogContentForTest(t, []string{
						"",
						"content1",
						"content2"}),
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := format.NewMarkdownParser()
			got, err := formatter.Parse(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarkdownFormatter.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownFormatter.Parse() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_ParseLog(t *testing.T) {
	type args struct {
		lines []string
	}

	tests := []struct {
		name    string
		args    args
		want    file.ParseReturn
		wantErr bool
	}{
		{
			name: "",
			args: args{
				lines: []string{
					"[12:34] @tag",
					"",
					"content1",
					"content2",
					"",
				},
			},
			want: file.NewParseReturn(
				time.Date(0, time.January, 1, 12, 34, 0, 0, time.UTC),
				[]model.Tag{testhelper.NewTagForTest(t, "@tag")},
				testhelper.NewLogContentForTest(t, []string{
					"",
					"content1",
					"content2",
					"",
				})),
		},
		{
			name: "",
			args: args{
				lines: []string{
					"[12:34:56] @tag",
					"",
					"content1",
					"content2",
					"",
				},
			},
			want: file.NewParseReturn(
				time.Date(0, time.January, 1, 12, 34, 56, 0, time.UTC),
				[]model.Tag{testhelper.NewTagForTest(t, "@tag")},
				testhelper.NewLogContentForTest(t, []string{
					"",
					"content1",
					"content2",
					"",
				})),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := format.NewMarkdownParser()
			got, err := formatter.ParseLogForTest(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarkdownParser.parseLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownParser.parseLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ParseLines(t *testing.T) {

	type args struct {
		lines []string
	}

	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				lines: []string{
					"gomi",
					"garbage",
					"[12:34] @tag",
					"",
					"content1",
					"content2",
					"",
					"[12:34:56]",
					"",
					"content1",
					"content2",
				},
			},
			want: [][]string{
				[]string{
					"[12:34] @tag",
					"",
					"content1",
					"content2",
					"",
				},
				[]string{
					"[12:34:56]",
					"",
					"content1",
					"content2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := format.NewMarkdownParser()
			got, err := formatter.ParseLinesForTest(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarkdownParser.parseLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownParser.parseLines() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_Format(t *testing.T) {

	type args struct {
		log model.Log
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC),
					[]string{"@tag1"},
					[]string{"", "content1", "content2"},
				),
			},
			want: []string{
				"[12:34] @tag1",
				"",
				"content1",
				"content2",
			},
		},
		{
			name: "????????????????????????????????????",
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 12, 34, 56, 0, time.UTC),
					[]string{"@tag1"},
					[]string{"", "content1", "content2"},
				),
			},
			want: []string{
				"[12:34:56] @tag1",
				"",
				"content1",
				"content2",
			},
		},
		{
			name: "???????????????????????????",
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 12, 34, 56, 0, time.UTC),
					[]string{"@tag1", "@tag2"},
					[]string{"", "content1", "content2"},
				),
			},
			want: []string{
				"[12:34:56] @tag1 @tag2",
				"",
				"content1",
				"content2",
			},
		},
		{
			name: "?????????????????????",
			args: args{
				log: testhelper.NewLogForTest(t,
					time.Date(2021, time.July, 10, 12, 34, 56, 0, time.UTC),
					[]string{},
					[]string{"", "content1", "content2"},
				),
			},
			want: []string{
				"[12:34:56]",
				"",
				"content1",
				"content2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			formatter := format.NewMarkdownFormatter()
			got := formatter.Format(tt.args.log)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarkdownFormatter.Format() got = %v, want %v", got, tt.want)
			}
		})
	}

}
