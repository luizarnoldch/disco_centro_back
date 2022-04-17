package infraestructure

import (
	"encoding/json"

	"github.com/luizarnoldch/disco_centro_back/src/discos/domain"
	"github.com/luizarnoldch/disco_centro_back/src/errs"
)

type DiscoRepository interface {
	RepoGetAllDiscos() ([]domain.Disco, *errs.AppError)
	RepoGetDisco(uint64) (*domain.Disco, *errs.AppError)
	RepoPostDisco(domain.Disco) (*domain.Disco, *errs.AppError)
	RepoUpdateDisco(domain.Disco) (*domain.Disco, *errs.AppError)
	RepoDeleteDisco(uint64) (*json.RawMessage, *errs.AppError)
}
