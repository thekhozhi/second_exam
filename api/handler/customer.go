package handler

import (
	"city2city/api/models"
	"encoding/json"
	"errors"
	"net/http"
)

func (h Handler) Customer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			h.GetCustomerList(w, r)
		} else {
			h.GetCustomerByID(w, r)
		}
	case http.MethodPut:
		h.UpdateCustomer(w, r)
	case http.MethodDelete:
		h.DeleteCustomer(w, r)
	}
}

func (h Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	createCustomer := models.CreateCustomer{}
	
	if err := json.NewDecoder(r.Body).Decode(&createCustomer); err != nil {
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	pKey, err := h.storage.Customer().Create(createCustomer)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	customer, err := h.storage.Customer().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusCreated, customer)
}

func (h Handler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]
	var err error

	customer, err := h.storage.Customer().Get(id)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, customer)
}

func (h Handler) GetCustomerList(w http.ResponseWriter, r *http.Request) {
	customer, err := h.storage.Customer().GetList()
	if err != nil{
		handleResponse(w,http.StatusInternalServerError,err)
	}
	handleResponse(w,http.StatusOK,customer)
}

func (h Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	updateCus := models.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&updateCus); err != nil {
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	pKey, err := h.storage.Customer().Update(updateCus)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	customer, err := h.storage.Car().Get(pKey)
	if err != nil {
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleResponse(w, http.StatusOK, customer)
}

func (h Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	if len(values["id"]) <= 0 {
		handleResponse(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	id := values["id"][0]

	err := h.storage.Customer().Delete(id)
	if err != nil{
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(w, http.StatusOK, "data successfully deleted")
}
