package models

type TripCustomer struct {
	ID           string   `json:"id"`
	TripID       string   `json:"trip_id"`
	CustomerID   string   `json:"customer_id"`
	CustomerData Customer `json:"customer_data"`
	CreatedAt    string   `json:"created_at"`
}

type CreateTripCustomer struct {
	TripID     string `json:"trip_id"`
	CustomerID string `json:"customer_id"`
}


type GetTripCusList struct {
	TripCusId         string`json:"trip_cus_id"`
	CustomerFullName  string`json:"customer_full_name"`
	CustomerPhone     string`json:"customer_phone"`
	CustomerEmail 	  string`json:"customer_email"`
	CustomerCreatedAt string`json:"customer_created_at"`
}
