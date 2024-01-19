package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type tripCustomerRepo struct {
	db *sql.DB
}

func NewTripCustomerRepo(db *sql.DB) storage.ITripCustomerRepo {
	return &tripCustomerRepo{
		db: db,
	}
}

func (tc *tripCustomerRepo) Create(tripCus models.CreateTripCustomer) (string, error) {
	id := uuid.New()
	createdAt := time.Now()
	query := `INSERT INTO trip_customers (id, trip_id, customer_id, created_at)
	values ($1,$2,$3,$4)`

	_, err := tc.db.Exec(query, id, tripCus.TripID, tripCus.CustomerID, createdAt)
	if err != nil{
		return "", err
	}

	return id.String(), nil
}

func (tc *tripCustomerRepo) Get(id string) (models.TripCustomer, error) {
	tripCustomer := models.TripCustomer{}
	query := `SELECT trip_id, customer_id, created_at from trip_customers`

	err := tc.db.QueryRow(query,id).Scan(
		&tripCustomer.TripID,
		&tripCustomer.CustomerID,
		&tripCustomer.CreatedAt,
	)

	if err != nil{
		return models.TripCustomer{}, err
	}
	return  tripCustomer, nil
}

func (tc *tripCustomerRepo) GetList(req models.GetListRequest) ([]models.GetTripCusList, error) {
	var (
		tripCustomers            = []models.GetTripCusList{}
		query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)


	query = `select trip_customers.id as trip_customers_id, customers.full_name, customers.phone, customers.email, 
	customers.created_at from trip_customers join customers on trip_customers.customer_id = customers.id`

	query += ` LIMIT $1 OFFSET $2`

	rows, err := tc.db.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return []models.GetTripCusList{}, err
	}

	for rows.Next() {
		tripCus := models.GetTripCusList{}

		err := rows.Scan(
			 &tripCus.TripCusId,
			 &tripCus.CustomerFullName,
			 &tripCus.CustomerPhone,
			 &tripCus.CustomerEmail,
			 &tripCus.CustomerCreatedAt,
		)
		if err != nil {
			fmt.Println("error while scanning row", err.Error())
			return []models.GetTripCusList{}, err
		}
		 tripCustomers = append(tripCustomers, tripCus)

	}
	return tripCustomers, nil
}

func (tc *tripCustomerRepo) Update(tCus models.TripCustomer) (string, error) {
	query := `UPDATE trip_customers set created_at = $1 where id = $2`
	_, err := tc.db.Exec(query,tCus.ID)
	if err != nil{
		return "", err
	}
	return tCus.ID, nil
}

func (tc *tripCustomerRepo) Delete(id string) error {
	query := `DELETE FROM trip_customers where id = $1`
	_, err := tc.db.Exec(query,id)
	if err != nil{
		return err
	}
	return nil
}
