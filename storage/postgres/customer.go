package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type customerRepo struct {
	DB *sql.DB
}

func NewCustomerRepo(db *sql.DB) storage.ICustomerRepo {
	return customerRepo{
		DB: db,
	}
}

func (c customerRepo) Create(customer models.CreateCustomer) (string, error) {
	id := uuid.New()
	createdAt := time.Now()
	query := `INSERT INTO customers (id, full_name, phone, email)
	values ($1,$2,$3,$4,$5)`

	_, err := c.DB.Exec(query, id, customer.FullName, customer.Phone, customer.Email, createdAt)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

func (c customerRepo) Get(id string) (models.Customer, error) {
	customer := models.Customer{}
	query := `SELECT full_name, phone, email, createdAt where id = $1`

	err := c.DB.QueryRow(query, id).Scan(
		&customer.FullName,
		&customer.Phone,
		&customer.Email,
		&customer.CreatedAt,
	)

	if err != nil{
		return models.Customer{}, err
	}

	return customer, nil
}

func (c customerRepo) GetList() ([]models.Customer, error) {
	customers := []models.Customer{}
	query := `SELECT full_name,phone,email,created_at from customers`

	rows, err := c.DB.Query(query)
	if err != nil{
		return []models.Customer{}, err
	}

	for rows.Next(){
		customer := models.Customer{}
		err := rows.Scan(
			&customer.FullName,
			&customer.Phone,
			&customer.Email,
			&customer.CreatedAt,
		)
		if err != nil{
			return []models.Customer{}, err
		}

		customers = append(customers, customer)
	}
	return customers, nil
}

func (c customerRepo) Update(customer models.Customer) (string, error) {
	query := `UPDATE customers set full_name = $1, phone = $2, email = $3`
	_, err := c.DB.Exec(query, customer.FullName, customer.Phone, customer.ID)
	if err != nil{
		return "", err
	}
	return customer.ID, nil
}

func (c customerRepo) Delete(id string) error {
	query := `DELETE FROM customers where id = $1`
	_, err := c.DB.Exec(query,id)
	if err != nil{
		return err
	}

	return nil
}
