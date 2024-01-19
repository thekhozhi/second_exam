package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type carRepo struct {
	DB *sql.DB
}

func NewCarRepo(db *sql.DB) storage.ICarRepo {
	return carRepo{
		DB: db,
	}
}

func (c carRepo) Create(car models.CreateCar) (string, error) {
	query := `INSERT INTO cars (id, model, brand, number, status, created_at) 
	values ($1,$2,$3,$4,$5)`
	id := uuid.New()

    _, err := c.DB.Exec(query, id, car.Model, car.Brand, car.Number, car.DriverID)
	if err != nil{
		return "",err
	} 

	return id.String(), nil
}

func (c carRepo) Get(id string) (models.Car, error) {
	car := models.Car{}
	query := `SELECT cars.model,cars.brand,cars.number,cars.created_at,drivers.full_name,drivers.phone,
	drivers.created_at FROM cars INNER JOIN drivers ON $1 = drivers.id`

	err := c.DB.QueryRow(query,id).Scan(
		&car.Model,
		&car.Brand,
		&car.Number,
		&car.CreatedAt,
		&car.DriverData.FullName,
		&car.DriverData.Phone,
		&car.DriverData.CreatedAt,
	)

	if err != nil{
		return models.Car{}, err
	}

	return car, nil
}

func (c carRepo) GetList(req models.GetListRequest) ([]models.GetCarList, error) {
	var (
		cars            = []models.GetCarList{}
		query string
		page              = req.Page
		offset            = (page - 1) * req.Limit
	)

	 query = `SELECT cars.model,cars.brand,cars.number,cars.created_at,drivers.full_name,drivers.phone,
	drivers.created_at FROM cars INNER JOIN drivers ON cars.driver_id = drivers.id`

	query += ` LIMIT $1 OFFSET $2`

	rows, err := c.DB.Query(query, req.Limit, offset)
	if err != nil {
		fmt.Println("error while query rows", err.Error())
		return []models.GetCarList{}, err
	}

	for rows.Next() {
		car := models.GetCarList{}

		err := rows.Scan(
			 &car.Model,
			 &car.Brand,
			 &car.Number,
			 &car.CarsCreatedAt,
			 &car.DriversFullName,
			 &car.DriversPhone,
			 &car.DriversCreatedAt,
		)
		if err != nil {
			fmt.Println("error while scanning row", err.Error())
			return []models.GetCarList{}, err
		}
		cars = append(cars, car)

	}
	return cars, nil
}

func (c carRepo) Update(car models.Car) (string, error) {
	query := `UPDATE cars set model = $1, brand = $2, number = $3 where id = $4`
	_, err := c.DB.Exec(query,car.Model, car.Brand, car.Number, car.ID)
	if err != nil{
		return "", err
	}
	return car.ID, nil
}

func (c carRepo) Delete(id string) error {
	query := `DELETE FROM cars where id = $1`
	_, err := c.DB.Exec(query,id)
	if err != nil{
		return err
	}

	return nil
}

func (c carRepo) UpdateCarRoute(car models.UpdateCarRoute) error {
	query := `UPDATE drivers set from_city_id = $1, to_city_id = $2 where id = $3`
	_, err := c.DB.Exec(query,car.FromCityID, car.ToCityID, car.CarID)
	if err != nil{
		return err
	}
	return nil
}
func (c carRepo) UpdateCarStatus(car models.UpdateCarStatus) error {
	query := `UPDATE cars set status = $1 where id = $2`
	_, err := c.DB.Exec(query,car.Status, car.ID)
	if err != nil{
		return err
	}
	return nil
}

