package generator

import (
	"fmt"

	"github.com/tortuepin/tolog_ddd/pkg/domain/model"
)

type LogGenerator interface {
	Generate(model.Log) (model.Log, error)
}

type MultiLogGenerator struct {
	generators []func(model.Log) (model.Log, error)
}

func NewMultiLogGenerator(generators []func(model.Log) (model.Log, error)) (*MultiLogGenerator, error) {
	return &MultiLogGenerator{generators: generators}, nil
}

func (g *MultiLogGenerator) Generate(log model.Log) (model.Log, error) {
	ret := log
	var err error
	for _, gen := range g.generators {
		ret, err = gen(ret)
		if err != nil {
			return log, fmt.Errorf("error in MultiLogGenerator.Generate(): %w", err)
		}
	}

	return ret, nil
}
