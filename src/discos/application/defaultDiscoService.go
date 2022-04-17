package application

import (
	"encoding/json"

	"github.com/luizarnoldch/disco_centro_back/src/discos/domain"
	"github.com/luizarnoldch/disco_centro_back/src/discos/infraestructure"
	"github.com/luizarnoldch/disco_centro_back/src/errs"
)

type DefaultDiscoService struct {
	repo infraestructure.DiscoRepository
}

func (s DefaultDiscoService) ServiceGetAllDiscos() ([]domain.DiscoResponse, *errs.AppError) {

	discos, err := s.repo.RepoGetAllDiscos()
	if err != nil {
		return nil, err
	}
	response := make([]domain.DiscoResponse, 0)
	for _, c := range discos {
		response = append(response, c.ToResponse())
	}

	return response, err
}

func (s DefaultDiscoService) ServiceGetDisco(id uint64) (*domain.DiscoResponse, *errs.AppError) {
	c, err := s.repo.RepoGetDisco(id)
	if err != nil {
		return nil, err
	}
	response := c.ToResponse()
	return &response, err
}

func (s DefaultDiscoService) ServicePostDisco(req domain.DiscoRequest) (*domain.DiscoResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	disco := domain.NewDisco(0, req.Nombre, req.Album, req.Artista, req.Fecha_Lanzamiento, req.Estado, req.Stock, req.Calidad)
	if newDisco, err := s.repo.RepoPostDisco(disco); err != nil {
		return nil, err
	} else {
		return newDisco.ToNewDiscoResponse(), nil
	}
}

func (s DefaultDiscoService) ServiceUpdateDisco(req domain.DiscoRequest) (*domain.DiscoResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	disco := domain.NewDisco(req.Id, req.Nombre, req.Album, req.Artista, req.Fecha_Lanzamiento, req.Estado, req.Stock, req.Calidad)
	if newDisco, err := s.repo.RepoUpdateDisco(disco); err != nil {
		return nil, err
	} else {
		return newDisco.ToNewDiscoResponse(), nil
	}
}

func (s DefaultDiscoService) ServiceDeleteDisco(id uint64) (**json.RawMessage, *errs.AppError) {
	c, err := s.repo.RepoDeleteDisco(id)
	if err != nil {
		return nil, err
	}
	return &c, err
}

func NewDiscoService(repo infraestructure.DiscoRepository) DefaultDiscoService {
	return DefaultDiscoService{repo}
}
