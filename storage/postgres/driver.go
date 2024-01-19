package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) storage.IDriverRepo {
	return driverRepo{
		DB: db,
	}
}

func (d driverRepo) Create(driver models.CreateDriver) (string, error) {

	id := uuid.New()
	query := `INSERT INTO drivers (id, full_name, phone, from_city_id, to_city_id) values ($1,$2,$3,$4,$5)`

	_, err := d.DB.Exec(query, id, driver.FullName, driver.Phone, driver.FromCityID, driver.ToCityID)
	if err != nil{
		return "", err
	}

	return id.String(), nil
}

func (d driverRepo) Get(id string) (models.GetDriver, error) {
	driver := models.GetDriver{}
	 query := `SELECT
	 drivers.full_name AS driver_full_name,
	 drivers.phone AS driver_phone,
	 from_cities.name AS from_city,
	 to_cities.name AS to_city,
	 customers.full_name AS customer_full_name,
	 customers.phone AS customer_phone,
	 customers.email AS customer_email,
	 trips.created_at AS trip_date
 FROM
	 trips
	 JOIN drivers ON trips.driver_id = $1
	 JOIN cities AS from_cities ON trips.from_city_id = from_cities.id
	 JOIN cities AS to_cities ON trips.to_city_id = to_cities.id
	 JOIN trip_customers tc ON trips.id = tc.trip_id
	 JOIN customers ON tc.customer_id = customers.id;`

	 err := d.DB.QueryRow(query,id).Scan(
		&driver.DriverFullName,
		&driver.DriverPhone,
		&driver.FromCitiesName,
		&driver.ToCitiesName,
		&driver.CustomerFullName,
		&driver.CustomerPhone,
		driver.CustomerEmail,
		driver.TripsDate,
	 )
	 if err != nil{
		return models.GetDriver{},err
	 }
	 return driver,nil
}

func (d driverRepo) GetList(req models.GetListRequest) ([]models.GetDriverList, error) {

	var (
		drivers             = []models.GetDriverList{}
		query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	query = `SELECT drivers.full_name, drivers.phone, from_cities.name as from_city, to_cities.name as to_city
	FROM drivers
	JOIN cities from_cities ON drivers.from_city_id = from_cities.id
	JOIN cities to_cities ON drivers.to_city_id = to_cities.id`

	query += ` LIMIT $1 OFFSET $2`

	rows, err := d.DB.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return []models.GetDriverList{}, err
	}

	for rows.Next() {
		driver := models.GetDriverList{}

		err := rows.Scan(
			 &driver.FullName,
			 &driver.Phone,
			 &driver.FromCity,
			 &driver.ToCity,
		)
		if err != nil {
			fmt.Println("error while scanning row", err.Error())
			return []models.GetDriverList{}, err
		}
		drivers = append(drivers, driver)

	}
	return drivers, nil
}

func (d driverRepo) Update(driver models.Driver) (string, error) {
	query := `UPDATE drivers set full_name = $1, phone = $2, from_city_id = $3, to_city_id = $4 where id = $5`
	_, err := d.DB.Exec(query, driver.FullName, driver.Phone, driver.FromCityID, driver.ToCityID, driver.ID)
	if err != nil{
		return "", err
	}
	return driver.ID, nil
}

func (d driverRepo) Delete(id string) error {
	query := `DELETE FROM drivers where id = $1`
	_, err := d.DB.Exec(query,id)
	if err != nil{
		return err
	}
	return nil
}
