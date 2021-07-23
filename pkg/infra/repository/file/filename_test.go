package file_test

import (
	"github.com/tortuepin/tolog_ddd/pkg/infra/repository/file"
	"reflect"
	"testing"
	"time"
)

func Test_NewFilenameFromDate(t *testing.T) {
	tests := []struct {
		name    string
		arg     time.Time
		want    file.Filename
		wantErr bool
	}{
		{
			name: "",
			arg:  time.Date(2021, time.July, 10, 0, 0, 0, 0, time.UTC),
			want: file.NewFilenameForTest("210710", ".md"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := file.NewFilenameFromDate(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilenameFromDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewFilenameFromString(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    file.Filename
		wantErr bool
	}{
		{
			name: "",
			arg:  "210710.md",
			want: file.NewFilenameForTest("210710", ".md"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := file.NewFilenameFromString(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFilenameFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilenameFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Filename(t *testing.T) {
	tests := []struct {
		name    string
		sut     file.Filename
		want    string
		wantErr bool
	}{
		{
			name: "",
			sut:  file.NewFilenameFromDate(time.Date(2021, time.July, 10, 0, 0, 0, 0, time.UTC)),
			want: "210710.md",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sut.Filename()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filename.Filename() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Date(t *testing.T) {
	tests := []struct {
		name    string
		sut     file.Filename
		want    time.Time
		wantErr bool
	}{
		{
			name: "",
			sut:  file.NewFilenameFromDate(time.Date(2021, time.July, 10, 0, 0, 0, 0, time.UTC)),
			want: time.Date(2021, time.July, 10, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sut.Date()
			if (err != nil) != tt.wantErr {
				t.Errorf("Filename.Date() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filename.Date() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Go(t *testing.T) {
	tests := []struct {
		name    string
		sut     file.Filename
		arg     int
		want    file.Filename
		wantErr bool
	}{
		{
			name: "",
			sut:  file.NewFilenameFromDate(time.Date(2021, time.July, 10, 0, 0, 0, 0, time.UTC)),
			arg:  1,
			want: file.NewFilenameForTest("210711", ".md"),
		},
		{
			name: "",
			sut:  file.NewFilenameFromDate(time.Date(2021, time.July, 10, 0, 0, 0, 0, time.UTC)),
			arg:  -1,
			want: file.NewFilenameForTest("210709", ".md"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sut.Go(tt.arg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filename.Go() got = %v, want %v", got, tt.want)
			}
		})
	}
}
