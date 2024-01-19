package models

type Driver struct {
	ID           string `json:"id"`
	FullName     string `json:"full_name"`
	Phone        string `json:"phone"`
	FromCityID   string `json:"from_city_id"`
	FromCityData City   `json:"from_city_data"`
	ToCityID     string `json:"to_city_id"`
	ToCityData   City   `json:"to_city_data"`
	CreatedAt    string `json:"created_at"`
}

type CreateDriver struct {
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	FromCityID string `json:"from_city_id"`
	ToCityID   string `json:"to_city_id"`
}

type GetDriver struct {
	DriverFullName   string`json:"driver_full_name"`
	DriverPhone      string`json:"driver_phone"`
	FromCitiesName   string`json:"from_cities_name"`
	ToCitiesName     string`json:"to_cities_name"`
	CustomerFullName string`json:"customer_full_name"`
	CustomerPhone    string`json:"customer_phone"`
	CustomerEmail    string`json:"customer_email"`
	TripsDate        string`json:"trips_date"`
}

type GetDriverList struct {
	FullName string`json:"full_name"`
	Phone    string`json:"phone"`
	FromCity string`json:"from_city"`
	ToCity   string`json:"to_city"`
}