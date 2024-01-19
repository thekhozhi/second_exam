package models

type Trip struct {
	ID           string `json:"id"`
	TripNumberID string `json:"trip_number_id"`
	FromCityID   string `json:"from_city_id"`
	FromCityData City   `json:"from_city_data"`
	ToCityID     string `json:"to_city_id"`
	ToCityData   City   `json:"to_city_data"`
	DriverID     string `json:"driver_id"`
	DriverData   Driver `json:"driver_data"`
	Price        int    `json:"price"`
	CreatedAt    string `json:"created_at"`
}



type CreateTrip struct {
	TripNumberID string `json:"trip_number_id"`
	FromCityID   string `json:"from_city_id"`
	ToCityID     string `json:"to_city_id"`
	DriverID     string `json:"driver_id"`
	Price        int    `json:"price"`
	CreatedAt    string `json:"created_at"`
}

type GetTripList struct {
	TripNumberID   string `json:"trip_number_id"`
	FromCity 	   string `json:"from_city"`
	ToCity		   string `json:"to_city"`
	DriverFullName string `json:"driver_full_name"`
	DriverPhone    string `json:"driver_phone"`
}