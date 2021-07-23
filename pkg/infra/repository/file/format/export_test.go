package format

import "github.com/tortuepin/tolog_ddd/pkg/domain/repository/file"

func (m *MarkdownParser) ParseLinesForTest(lines []string) ([][]string, error) {
	return m.parseLines(lines)
}

func (m *MarkdownParser) ParseLogForTest(lines []string) (file.ParseReturn, error) {
	return m.parseLog(lines)
}
