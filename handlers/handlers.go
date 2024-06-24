package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/million_dollar_space_programme/exoplanets/errorhandler"
	"github.com/million_dollar_space_programme/exoplanets/services"
	"github.com/million_dollar_space_programme/exoplanets/utils"
)

func CreateExoPlanet(w http.ResponseWriter, r *http.Request) {
	var exoPlanetInfo utils.ExoPlanetInfo

	err := json.NewDecoder(r.Body).Decode(&exoPlanetInfo)
	if err != nil {
		err := errorhandler.New("invalid request body", http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	id, err := services.CreateExoPlanet(r.Context(), exoPlanetInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	exoPlanetInfo.ID = id

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoPlanetInfo)
}

func GetExoPlanets(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	offsetParam := queryParams.Get("offset")
	sizeParam := queryParams.Get("size")

	offset, err := strconv.ParseUint(offsetParam, 10, 64)
	if err != nil {
		err := errorhandler.New("invalid offset", http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	size, err := strconv.ParseUint(sizeParam, 10, 64)
	if err != nil {
		err := errorhandler.New("invalid size", http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	input := utils.FindAllInput{
		Offset: offset,
		Size:   size,
	}

	findAllOutPut, err := services.GetExoPlanets(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(findAllOutPut)
}

func GetExoPlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	exoPlanetInfo, err := services.GetExoPlanet(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exoPlanetInfo)
}

func UpdateExoPlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var exoPlanetInfo utils.ExoPlanetInfo
	err := json.NewDecoder(r.Body).Decode(&exoPlanetInfo)
	if err != nil {
		err := errorhandler.New("invalid request body", http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = services.UpdateExoPlanet(r.Context(), id, exoPlanetInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteExoPlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := services.DeleteExoPlanet(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func FuelEstimation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	mp := make(map[string]int)

	err := json.NewDecoder(r.Body).Decode(&mp)
	if err != nil || mp["crew_members"] <= 0 {
		err := errorhandler.New("invalid request body", http.StatusBadRequest, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	estimation, err := services.FuelEstimation(r.Context(), id, mp["crew_members"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]float64{"estimated_fuel": estimation})
}
