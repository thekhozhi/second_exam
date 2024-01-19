package handler

import (
	"city2city/api/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (h Handler) TripCustomer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTripCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		if _, ok := values["id"]; !ok {
			h.GetTripCustomerList(w, r)
		} else {
			h.GetTripCustomerByID(w, r)
		}
	case http.MethodPut:
		h.UpdateTripCustomer(w, r)
	case http.MethodDelete:
		h.DeleteTripCustomer(w, r)
	}
}

func (h Handler) CreateTripCustomer(w http.ResponseWriter, r *http.Request) {
	createTripCus := models.CreateTripCustomer{}
	
	if err := json.NewDecoder(r.Body).Decode(&createTripCus); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.TripCustomer().Create(createTripCus)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	tripCustomer, err := h.storage.TripCustomer().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, tripCustomer)
}

func (h Handler) GetTripCustomerByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	tripCus, err := h.storage.TripCustomer().Get(id)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, tripCus)
}

func (h Handler) GetTripCustomerList(w http.ResponseWriter, r *http.Request) {
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

	resp, err := h.storage.TripCustomer().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
	})

		 
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, resp)
}

func (h Handler) UpdateTripCustomer(w http.ResponseWriter, r *http.Request) {
	updateTripCus := models.TripCustomer{}

	if err := json.NewDecoder(r.Body).Decode(&updateTripCus); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.TripCustomer().Update(updateTripCus)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	tripCus, err := h.storage.Driver().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, tripCus)
}

func (h Handler) DeleteTripCustomer(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.storage.TripCustomer().Delete(id)
	if err != nil{
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, "data successfully deleted")
}
