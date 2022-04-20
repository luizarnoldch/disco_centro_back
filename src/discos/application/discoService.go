package application

import (
	"encoding/json"

	"github.com/luizarnoldch/disco_centro_back/src/discos/domain"
	"github.com/luizarnoldch/disco_centro_lib/errs"
)

type DiscoService interface {
	ServiceGetAllDiscos() ([]domain.DiscoResponse, *errs.AppError)
	ServiceGetDisco(uint64) (*domain.DiscoResponse, *errs.AppError)
	ServicePostDisco(domain.DiscoRequest) (*domain.DiscoResponse, *errs.AppError)
	ServiceUpdateDisco(domain.DiscoRequest) (*domain.DiscoResponse, *errs.AppError)
	ServiceDeleteDisco(uint64) (**json.RawMessage, *errs.AppError)
}
