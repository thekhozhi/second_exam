package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type cityRepo struct {
	DB *sql.DB
}

func NewCityRepo(db *sql.DB) storage.ICityRepo {
	return cityRepo{
		DB: db,
	}
}

func (c cityRepo) Create(city models.CreateCity) (string, error) {
	id := uuid.New()
	createdAt := time.Now()
	query := `INSERT INTO cities (id, name, created_at) values ($1,$2,$3)`

	_, err := c.DB.Exec(query, id, city.Name, createdAt)
	if err != nil{
		return "", err
	}

	return id.String(), nil
}

func (c cityRepo) Get(id string) (models.City, error) {
	city := models.City{}
	query := `SELECT FROM name, created_at cities where id = $1`

	err := c.DB.QueryRow(query,id).Scan(
		&city.Name,
		&city.CreatedAt,
	)

	if err != nil{
		return models.City{}, err
	}
	return city, nil
}

func (c cityRepo) GetList() ([]models.City, error) {
	cities := []models.City{}
	query := `SELECT name, created_at from cities`
	rows, err := c.DB.Query(query)
	if err != nil{
		return []models.City{}, err
	}
	for rows.Next(){
		city := models.City{}
		err := rows.Scan(
			&city.Name,
			&city.CreatedAt,
		)
		if err != nil{
			return []models.City{},err
		}
		cities = append(cities, city)
	}
	return cities, nil
}

func (c cityRepo) Update(city models.City) (string, error) {
	query := `UPDATE cities set name = $1 where id = $2`
	_, err := c.DB.Exec(query,city.Name, city.ID)
	if err != nil{
		return "", err
	}
	return city.ID, nil
}

func (c cityRepo) Delete(id string) error {
	query := `DELETE from cities where id = $1`
	_, err := c.DB.Exec(query,id)
	if err != nil{
		return err
	}
	return nil
}
