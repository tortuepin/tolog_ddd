package registry

import "fmt"
import "github.com/tortuepin/tolog_ddd/pkg/infra/repository"
import "github.com/tortuepin/tolog_ddd/pkg/service"

func RegisterFileLogService(dir string) (service.LogServiceInterface, error) {
	parser := repository.NewMarkdownParser()
	formatter := repository.NewMarkdownFormatter()
	f, err := repository.NewFile(dir, parser, formatter)
	if err != nil {
		return nil, fmt.Errorf("error in RegisterFileLogService(): %w", err)
	}

	logservice, err := service.NewLogService(f, f, nil)
	if err != nil {
		return nil, fmt.Errorf("error in RegisterFileLogService(): %w", err)
	}

	return logservice, nil
}
