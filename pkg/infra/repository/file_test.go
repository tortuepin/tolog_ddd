package repository_test

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/domain/repository/file"
	"github.com/tortuepin/tolog_ddd/pkg/infra/repository"
)

type fakeParse struct {
	file.Parser
	fakeParse func([]string) ([]file.ParseReturn, error)
}

func (f *fakeParse) Parse(s []string) ([]file.ParseReturn, error) {
	return f.fakeParse(s)
}

type fakeFormat struct {
	file.Formatter
	fakeFormat func(model.Log) []string
}

func (f *fakeFormat) Format(log model.Log) []string {
	return f.fakeFormat(log)
}

func Test_Read(t *testing.T) {
	type fields struct {
		dir   string
		parse file.Parser
	}
	tests := []struct {
		name    string
		fields  fields
		want    []model.Log
		wantErr bool
	}{
		{
			name: "MarkdownParserを使ったテスト",
			fields: fields{
				dir: "testdata/read",
				parse: &fakeParse{
					fakeParse: func(lines []string) ([]file.ParseReturn, error) {
						p := repository.NewMarkdownParser()
						return p.Parse(lines)
					},
				},
			},
			want: []model.Log{
				repository.NewLogForTest(
					repository.NewLogTimeForTest(time.Date(2021, time.July, 10, 12, 34, 0, 0, time.UTC)),
					[]model.Tag{repository.NewTagForTest("@tag1")},
					repository.NewLogContentForTest([]string{"", "content1", "content2"}),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := repository.NewFile(tt.fields.dir, tt.fields.parse, nil)
			if err != nil {
				t.Errorf("repository.Newfile() error = %w", err)
			}
			got, err := f.Read()
			if (err != nil) != tt.wantErr {
				t.Errorf("File.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.Read() got = %v, want %v", got, tt.want)
			}
		})
	}

}

const CreateTestDataDir = "testdata/create"

func Test_Create(t *testing.T) {
	type fields struct {
		dir    string
		format file.Formatter
	}
	type args struct {
		log model.Log
	}
	type want struct {
		filename string
		contents []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "秒まで",
			args: args{
				log: repository.NewLogForTest(
					repository.NewLogTimeForTest(time.Date(2021, time.July, 11, 12, 34, 56, 0, time.UTC)),
					[]model.Tag{repository.NewTagForTest("@tag1"), repository.NewTagForTest("@tag2")},
					repository.NewLogContentForTest([]string{"", "content1", "content2"}),
				),
			},
			fields: fields{
				dir: CreateTestDataDir,
				format: &fakeFormat{
					fakeFormat: func(log model.Log) []string {
						f := repository.NewMarkdownFormatter()
						return f.Format(log)
					},
				},
			},
			want: want{
				filename: CreateTestDataDir + "/210711.md",
				contents: []string{"", "[12:34:56] @tag1 @tag2", "", "content1", "content2"},
			},
		},
		{
			name: "秒なし",
			args: args{
				log: repository.NewLogForTest(
					repository.NewLogTimeForTest(time.Date(2021, time.July, 11, 12, 34, 0, 0, time.UTC)),
					[]model.Tag{repository.NewTagForTest("@tag1"), repository.NewTagForTest("@tag2")},
					repository.NewLogContentForTest([]string{"", "content1", "content2"}),
				),
			},
			fields: fields{
				dir: CreateTestDataDir,
				format: &fakeFormat{
					fakeFormat: func(log model.Log) []string {
						f := repository.NewMarkdownFormatter()
						return f.Format(log)
					},
				},
			},
			want: want{
				filename: CreateTestDataDir + "/210711.md",
				contents: []string{"", "[12:34] @tag1 @tag2", "", "content1", "content2"},
			},
		},
		{
			name: "tagなし",
			args: args{
				log: repository.NewLogForTest(
					repository.NewLogTimeForTest(time.Date(2021, time.July, 11, 12, 34, 56, 0, time.UTC)),
					[]model.Tag{},
					repository.NewLogContentForTest([]string{"", "content1", "content2"}),
				),
			},
			fields: fields{
				dir: CreateTestDataDir,
				format: &fakeFormat{
					fakeFormat: func(log model.Log) []string {
						f := repository.NewMarkdownFormatter()
						return f.Format(log)
					},
				},
			},
			want: want{
				filename: CreateTestDataDir + "/210711.md",
				contents: []string{"", "[12:34:56]", "", "content1", "content2"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := creanUpDirForCreate(); err != nil {
				t.Errorf("File.Create() failed to crean up: %v", err)
				return
			}
			f, err := repository.NewFile(tt.fields.dir, nil, tt.fields.format)
			if err != nil {
				t.Errorf("repository.Newfile() error = %w", err)
			}
			err = f.Create(tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("File.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 正しく書き込まれたかを確認
			if err := verifyWrittenContents(tt.want.filename, tt.want.contents); err != nil {
				t.Errorf("File.Create() error = %v", err)
				return
			}
		})
	}
}

func verifyWrittenContents(filename string, want []string) error {
	// ファイルをreadする
	got, err := readLines(filename)
	if err != nil {
		return fmt.Errorf("error in verifyWrittenContents: %w", err)
	}
	// 内容がcontentsと一致することを確認する
	if !reflect.DeepEqual(got, want) {
		return fmt.Errorf("got  = %v\n, want = %v", got, want)
	}
	return nil
}

func readLines(filename string) ([]string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return []string{}, fmt.Errorf("cannnot open %s", filename)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	ret := []string{}
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret, nil
}

func creanUpDirForCreate() error {
	if err := os.RemoveAll(CreateTestDataDir); err != nil {
		return fmt.Errorf("error in creanUpDirForCreate(): %w", err)
	}
	if err := os.MkdirAll(CreateTestDataDir, 0755); err != nil {
		return fmt.Errorf("error in creanUpDirForCreate(): %w", err)
	}
	return nil
}
