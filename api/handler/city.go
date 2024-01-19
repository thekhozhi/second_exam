package handler

import (
	"city2city/api/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (h Handler) City(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCity(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetCityList(w)
		} else {
			h.GetCityByID(w, r)
		}
	case http.MethodPut:
		h.UpdateCity(w, r)
	case http.MethodDelete:
		h.DeleteCity(w, r)
	}
}

func (h Handler) CreateCity(w http.ResponseWriter, r *http.Request) {
	createCity := models.CreateCity{}

	if err := json.NewDecoder(r.Body).Decode(&createCity); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.City().Create(createCity)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	city, err := h.storage.City().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, city)
}

func (h Handler) GetCityByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	city, err := h.storage.City().Get(id)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, city)
}

func (h Handler) GetCityList(w http.ResponseWriter) {
	city, err := h.storage.City().GetList()
	if err != nil{
		handleResponse(w,http.StatusInternalServerError,err)
	}
	handleResponse(w,http.StatusOK,city)

}

func (h Handler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	updateCity := models.City{}

	if err := json.NewDecoder(r.Body).Decode(&updateCity); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.City().Update(updateCity)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	city, err := h.storage.Car().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, city)
}

func (h Handler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.storage.City().Delete(id)
	if err != nil{
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, "data successfully deleted")
}
