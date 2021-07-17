package model_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

func Test_NewLog(t *testing.T) {
	type args struct {
		time    model.LogTime
		tags    []model.Tag
		content model.LogContent
	}
	tests := []struct {
		name    string
		args    args
		want    model.Log
		wantErr bool
	}{
		{
			name: "",
			args: args{
				time:    model.NewLogTimeForTest(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
				tags:    []model.Tag{model.NewTagForTest("@tag1"), model.NewTagForTest("@tag2")},
				content: model.NewLogContentForTest([]string{"content", "content"}),
			},
			want: model.NewLogForTest(
				model.NewLogTimeForTest(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)),
				[]model.Tag{model.NewTagForTest("@tag1"), model.NewTagForTest("@tag2")},
				model.NewLogContentForTest([]string{"content", "content"}),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewLog(tt.args.time, tt.args.tags, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("model.NewLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("model.NewLog() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func Test_NewLogTime(t *testing.T) {
	type args struct {
		time time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    model.LogTime
		wantErr bool
	}{
		{
			name: "nanosecondで切り捨てされて正しくNewされる",
			args: args{
				time: time.Date(2009, time.November, 10, 23, 10, 10, 99, time.UTC),
			},
			want: model.NewLogTimeForTest(time.Date(2009, time.November, 10, 23, 10, 10, 0, time.UTC)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewLogTime(tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("model.NewLogTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("model.NewLogTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Tag
		wantErr bool
	}{
		{
			name: "",
			args: args{
				tag: "@tag1",
			},
			want: model.NewTagForTest("@tag1"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewTag(tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("model.NewTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("model.NewTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewTags(t *testing.T) {
	type args struct {
		tags []string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Tag
		wantErr bool
	}{
		{
			name: "",
			args: args{
				tags: []string{"@tag1", "@tag2"},
			},
			want: []model.Tag{
				model.NewTagForTest("@tag1"),
				model.NewTagForTest("@tag2")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewTags(tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("model.NewTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("model.NewTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewLogContent(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name    string
		args    args
		want    model.LogContent
		wantErr bool
	}{
		{
			name: "",
			args: args{
				content: []string{"line1", "line2"},
			},
			want: model.NewLogContentForTest([]string{"line1", "line2"}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.NewLogContent(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("model.NewLogContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("model.NewLogContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
