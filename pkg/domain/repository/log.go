package repository

import "github.com/tortuepin/tolog_ddd/pkg/domain/model"

type Reader interface {
	Read() ([]model.Log, error)
}

type Creater interface {
	Create(model.Log) error
}

type Updater interface {
	Update(model.Log, model.Log) error
}
