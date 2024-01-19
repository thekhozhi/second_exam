package storage

import (
	"city2city/api/models"
)

type IStorage interface {
	CloseDB()
	City() ICityRepo
	Customer() ICustomerRepo
	Driver() IDriverRepo
	Car() ICarRepo
	Trip() ITripRepo
	TripCustomer() ITripCustomerRepo
}

type ICityRepo interface {
	Create(city models.CreateCity) (string, error)
	Get(id string) (models.City, error)
	GetList() ([]models.City, error)
	Update(city models.City) (string, error)
	Delete(id string) error
}

type ICustomerRepo interface {
	Create(customer models.CreateCustomer) (string, error)
	Get(id string) (models.Customer, error)
	GetList() ([]models.Customer, error)
	Update(customer models.Customer) (string, error)
	Delete(id string) error
}

type IDriverRepo interface {
	Create(driver models.CreateDriver) (string, error)
	Get(id string) (models.GetDriver, error)
	GetList(req models.GetListRequest) ([]models.GetDriverList, error)
	Update(driver models.Driver) (string, error)
	Delete(id string) error
}

type ICarRepo interface {
	Create(car models.CreateCar) (string, error)
	Get(id string) (models.Car, error)
	GetList(req models.GetListRequest) ([]models.GetCarList, error)
	Update(car models.Car) (string, error)
	Delete(id string) error
	UpdateCarStatus(updateCarStatus models.UpdateCarStatus) error
	UpdateCarRoute(updateCarRoute models.UpdateCarRoute) error
}

type ITripRepo interface {
	Create(trip models.CreateTrip) (string, error)
	Get(id string) (models.Trip, error)
	GetList(req models.GetListRequest) ([]models.GetTripList, error)
	Update(trip models.Trip) (string, error)
	Delete(id string) error
}

type ITripCustomerRepo interface {
	Create(tripCustomer models.CreateTripCustomer) (string, error)
	Get(id string) (models.TripCustomer, error)
	GetList(req models.GetListRequest) ([]models.GetTripCusList, error)
	Update(tripCus models.TripCustomer) (string, error)
	Delete(id string) error
}
