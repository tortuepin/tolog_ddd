package registry

import "fmt"
import "github.com/tortuepin/tolog_ddd/pkg/infra/repository/file"
import "github.com/tortuepin/tolog_ddd/pkg/infra/repository/file/format"
import "github.com/tortuepin/tolog_ddd/pkg/service"

func RegisterFileLogService(dir string) (service.LogServiceInterface, error) {
	parser := format.NewMarkdownParser()
	formatter := format.NewMarkdownFormatter()
	f, err := file.NewFile(dir, parser, formatter)
	if err != nil {
		return nil, fmt.Errorf("error in RegisterFileLogService(): %w", err)
	}

	logservice, err := service.NewLogService(f, f, nil)
	if err != nil {
		return nil, fmt.Errorf("error in RegisterFileLogService(): %w", err)
	}

	return logservice, nil
}
