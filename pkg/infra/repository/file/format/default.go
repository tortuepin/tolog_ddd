package format

import (
	"fmt"
	"strings"
	"time"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
	"github.com/tortuepin/tolog_ddd/pkg/domain/repository/file"
)

func NewMarkdownParser() *MarkdownParser {
	return &MarkdownParser{}
}

type MarkdownParser struct{}

func (m *MarkdownParser) Parse(lines []string) ([]file.ParseReturn, error) {
	// log毎に分割
	parsedLines, err := m.parseLines(lines)
	if err != nil {
		return []file.ParseReturn{}, fmt.Errorf("error in parseLines: %w", err)
	}

	returns := []file.ParseReturn{}
	for _, plines := range parsedLines {
		parseReturn, err := m.parseLog(plines)
		if err != nil {
			return []file.ParseReturn{}, fmt.Errorf("error in parseLines: %w", err)
		}
		returns = append(returns, parseReturn)
	}
	return returns, nil
}

func (m *MarkdownParser) parseLog(lines []string) (file.ParseReturn, error) {
	// 最初のlineからtimeとtagを取得
	firstline := lines[0]
	splited := strings.Split(firstline, " ")

	timestr := splited[0]
	timepart, err := m.parseTime(timestr)
	if err != nil {
		return file.ParseReturn{}, fmt.Errorf("error in parseLog: %w", err)
	}

	tagsstr := splited[1:]
	tags, err := model.NewTags(tagsstr)
	if err != nil {
		return file.ParseReturn{}, fmt.Errorf("error in parseLog: %w", err)
	}

	// parse content
	c := make([]string, len(lines)-1)
	copy(c, lines[1:])
	content, err := model.NewLogContent(c)
	if err != nil {
		return file.ParseReturn{}, fmt.Errorf("error in parseLog: %w", err)
	}

	return file.NewParseReturn(timepart, tags, content), nil
}

func (m *MarkdownParser) parseTime(line string) (time.Time, error) {
	layouts := []string{"[15:04]", "[15:04:05]"}
	for _, l := range layouts {
		t, err := time.Parse(l, line)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("error in parseTime: cannnot parse %s", line)
}

func (m *MarkdownParser) parseLines(lines []string) ([][]string, error) {
	ret := [][]string{}
	l := []string{}
	for _, line := range lines {
		if m.isFirstLine(line) {
			ret = append(ret, l)
			l = []string{}
		}
		l = append(l, line)
	}
	ret = append(ret, l)

	return deletetop(ret), nil
}

var replaceKey = []string{"0", "?", "1", "?", "2", "?", "3", "?", "4", "?", "5", "?", "6", "?", "7", "?", "8", "?", "9", "?"}
var replacer = strings.NewReplacer(replaceKey...)

func (m *MarkdownParser) isFirstLine(line string) bool {
	replaced := replacer.Replace(line)
	if strings.HasPrefix(replaced, "[??:??]") || strings.HasPrefix(replaced, "[??:??:??]") {
		return true
	}

	return false
}

func deletetop(s [][]string) [][]string {
	s = s[1:]
	n := make([][]string, len(s))
	copy(n, s)
	return n
}

func NewMarkdownFormatter() *MarkdownFormatter {
	return &MarkdownFormatter{}
}

type MarkdownFormatter struct{}

func (m *MarkdownFormatter) Format(log model.Log) []string {
	f := []string{}
	firstline := strings.TrimSpace(m.formatTime(log.Time()) + " " + m.formatTags(log.Tags()))
	f = append(f, firstline)
	f = append(f, m.formatContent(log.Content())...)
	return f
}

func (m *MarkdownFormatter) formatTime(t model.LogTime) string {
	if t.Time().Second() == 0 {
		return t.Time().Format("[15:04]")
	}
	return t.Time().Format("[15:04:05]")
}
func (m *MarkdownFormatter) formatTags(tags []model.Tag) string {
	tagsstr := ""
	for _, t := range tags {
		tagsstr = tagsstr + " " + t.Tag()
	}
	return strings.TrimSpace(tagsstr)
}
func (m *MarkdownFormatter) formatContent(content model.LogContent) []string {
	return content.Content()
}
