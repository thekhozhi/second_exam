package handler

import (
	"city2city/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCar(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetCarList(w,r)
		} else {
			h.GetCarByID(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		if _, ok := values["route"]; ok {
			h.UpdateCarRoute(w, r)
		} else if _, ok := values["status"]; ok {
			h.UpdateCarStatus(w, r)
		} else {
			h.UpdateCar(w, r)
		}
	case http.MethodDelete:
		h.DeleteCar(w, r)
	}
}

func (h Handler) CreateCar(w http.ResponseWriter, r *http.Request) {
	createCar := models.CreateCar{}

	if err := json.NewDecoder(r.Body).Decode(&createCar); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Car().Create(createCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	car, err := h.storage.Car().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, car)
}

func (h Handler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	car, err := h.storage.Car().Get(id)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, car)
}

func (h Handler) GetCarList(w http.ResponseWriter, r *http.Request) {
	var (
		page, limit = 1, 10
		err         error
	)
	values := r.URL.Query()

	if len(values["page"]) > 0 {
		page, err = strconv.Atoi(values["page"][0])
		if err != nil {
			page = 1
		}
	}

	if len(values["limit"]) > 0 {
		limit, err = strconv.Atoi(values["limit"][0])
		if err != nil {
			fmt.Println("limit", values["limit"])
			limit = 10
		}
	}

	resp, err := h.storage.Car().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
	})

		 
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (h Handler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	updateCar := models.Car{}

	if err := json.NewDecoder(r.Body).Decode(&updateCar); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Car().Update(updateCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	car, err := h.storage.Car().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, car)
}

func (h Handler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.storage.Car().Delete(id)
	if err != nil{
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, "data successfully deleted")
}

func (h Handler) UpdateCarRoute(w http.ResponseWriter, r *http.Request) {
	updateCar := models.UpdateCarRoute{}

	if err := json.NewDecoder(r.Body).Decode(&updateCar); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.storage.Car().UpdateCarRoute(updateCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK,nil)
}

func (h Handler) UpdateCarStatus(w http.ResponseWriter, r *http.Request) {
	updateCar := models.UpdateCarStatus{}

	if err := json.NewDecoder(r.Body).Decode(&updateCar); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := h.storage.Car().UpdateCarStatus(updateCar)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK,nil)
}
