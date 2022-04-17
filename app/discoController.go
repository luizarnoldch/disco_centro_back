package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luizarnoldch/disco_centro_back/src/discos/application"
	"github.com/luizarnoldch/disco_centro_back/src/discos/domain"
)

type DiscoController struct {
	service application.DiscoService
}

func (dc DiscoController) GetAllDiscos(w http.ResponseWriter, r *http.Request) {
	discos, err := dc.service.ServiceGetAllDiscos()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, discos)
	}
}

func (dc DiscoController) GetDisco(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	discoId := vars["id_disco"]
	id, _ := strconv.ParseUint(discoId, 10, 64)
	discos, err := dc.service.ServiceGetDisco(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, discos)
	}
}

func (dc DiscoController) PostDisco(w http.ResponseWriter, r *http.Request) {
	var request domain.DiscoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		disco, appError := dc.service.ServicePostDisco(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, disco)
		}
	}
}

func (dc DiscoController) UpdateDisco(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	discoId, _ := strconv.ParseUint(vars["id_disco"], 10, 64)
	var request domain.DiscoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.Id = discoId
		disco, appError := dc.service.ServiceUpdateDisco(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, disco)
		}
	}
}

func (dc DiscoController) DeleteDisco(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	discoId := vars["id_disco"]
	id, _ := strconv.ParseUint(discoId, 10, 64)
	discos, err := dc.service.ServiceDeleteDisco(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, discos)
	}
}
