package repository

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

type File struct {
	dir    string
	parse  Parser
	format Formatter
}

func NewFile(dir string, parse Parser, format Formatter) (*File, error) {
	_, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("dir is not exist: %s", dir)
	}
	return &File{dir: dir, parse: parse, format: format}, nil
}

func (f *File) Dir() string {
	return f.dir
}

func (f *File) Read() ([]model.Log, error) {
	pattern := filepath.Join(f.dir, "/*"+EXT)
	files, err := filepath.Glob(pattern)
	if err != nil {
		return []model.Log{}, fmt.Errorf("failed in File.Read(): %w", err)
	}

	logs := []model.Log{}
	for _, file := range files {
		filename, err := NewFilenameFromString(filepath.Base(file))
		if err != nil {
			return []model.Log{}, fmt.Errorf("failed in File.Read(): %w", err)
		}
		date, err := filename.Date()
		if err != nil {
			return []model.Log{}, fmt.Errorf("failed in File.Read(): %w", err)
		}
		lines, err := readLines(file)
		if err != nil {
			return []model.Log{}, fmt.Errorf("failed in File.Read(): %w", err)
		}
		parseReturns, err := f.parse.Parse(lines)
		if err != nil {
			return []model.Log{}, fmt.Errorf("failed in File.Read(): %w", err)
		}
		for _, p := range parseReturns {
			l, err := p.ToLog(date.Year(), date.Month(), date.Day()) // TODO
			if err != nil {
				return []model.Log{}, fmt.Errorf("failed in File.Read(): %w", err)
			}
			logs = append(logs, l)
		}
	}

	return logs, nil
}

func (f *File) Create(log model.Log) error {
	// logをformatする
	lines := f.format.Format(log)
	// 追記する
	filename := NewFilenameFromDate(log.Time().Time())
	path := filepath.Join(f.dir, filename.Filename())
	if err := appendLines(path, lines); err != nil {
		return fmt.Errorf("error in Create(): %w", err)
	}
	return nil
}

func appendLines(filename string, lines []string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("cannnot open %s", filename)
	}
	defer f.Close()

	for _, l := range lines {
		if _, err = f.WriteString("\n" + l); err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}
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
