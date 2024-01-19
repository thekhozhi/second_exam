package models

import "time"

type Car struct {
	ID         string `json:"id"`
	Model      string `json:"model"`
	Brand      string `json:"brand"`
	Number     int    `json:"number"`
	DriverID   string `json:"driver_id"`
	DriverData Driver `json:"driver_data"`
	CreatedAt  string `json:"created_at"`
}

type CreateCar struct {
	Model    string `json:"model"`
	Brand    string `json:"brand"`
	Number   int    `json:"number"`
	DriverID string `json:"driver_id"`
}

type UpdateCarStatus struct {
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

type UpdateCarRoute struct {
	CarID         string    `json:"car_id"`
	DepartureTime time.Time `json:"departure_time"`
	FromCityID    string    `json:"from_city_id"`
	ToCityID      string    `json:"to_city_id"`
}

type GetCarList struct {
	Model 			 string`json:"model"`
	Brand 			 string`json:"brand"`
	Number 			 string`json:"number"`
	CarsCreatedAt    string`json:"cars_created_at"`
	DriversFullName  string`json:"drivers_full_name"`
	DriversPhone     string`json:"drivers_phone"`
	DriversCreatedAt string`json:"derivers_created_at"`
}