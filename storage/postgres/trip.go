package postgres

import (
	"city2city/api/models"
	"city2city/storage"
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type tripRepo struct {
	db *sql.DB
}

func NewTripRepo(db *sql.DB) storage.ITripRepo {
	return &tripRepo{
		db: db,
	}
}

func (t *tripRepo) Create(trip models.CreateTrip) (string, error) {
	id := uuid.New()
	createdAt := time.Now()
	query := `INSERT INTO trips (id, from_city_id, to_city_id, 
	driver_id, price, created_at) values ($1,$2,$3,$4,$5,$6)`
	
	_, err := t.db.Exec(query, id, trip.FromCityID, trip.ToCityID,
	trip.DriverID, trip.Price, createdAt)
	if err != nil{
		return "", err
	}

	return id.String(), nil
}

func (t *tripRepo) Get(id string) (models.Trip, error) {
	trip := models.Trip{}
	query := `SELECT trip_number_id, from_city_id, to_city_id, driver_id, price, 
	created_at from trips where id = $1
	`
	err := t.db.QueryRow(query,id).Scan(
		&trip.TripNumberID,
		&trip.FromCityID,
		&trip.ToCityID,
		&trip.DriverID,
		&trip.Price,
		&trip.CreatedAt,
	)

	if err != nil{
		return models.Trip{}, err
	}
	
	return trip, nil
}

func (t *tripRepo) GetList(req models.GetListRequest) ([]models.GetTripList, error) {
	trips := []models.GetTripList{}
	query := `SELECT trips.trip_number_id, from_cities.name, to_cities.name, drivers.full_name, drivers.phone
FROM trips
JOIN cities AS from_cities ON trips.from_city_id = from_cities.id
JOIN cities AS to_cities ON trips.to_city_id = to_cities.id
JOIN drivers ON drivers.from_city_id = trips.from_city_id`

rows,err := t.db.Query(query)
if err != nil{
	return []models.GetTripList{},err
}

for rows.Next(){
	trip := models.GetTripList{}
	err := rows.Scan(
		&trip.TripNumberID,
		&trip.FromCity,
		&trip.ToCity,
		&trip.DriverFullName,
		&trip.DriverPhone,
	)
	if err != nil{
		return []models.GetTripList{}, err
	}

	trips = append(trips, trip)	
}
	return trips, nil
}

func (t *tripRepo) Update(trip models.Trip) (string, error) {
	query := `UPDATE trips set price = $1 where id = $2`
	_, err := t.db.Exec(query,trip.Price,trip.ID)
	if err != nil{
		return "", err
	}
	return strconv.Itoa(trip.Price), nil
}

func (t *tripRepo) Delete(id string) error {
	query := `DELETE from trips where id = $1`
	_, err := t.db.Exec(query,id)
	if err != nil{
		return err
	}

	return nil
}
