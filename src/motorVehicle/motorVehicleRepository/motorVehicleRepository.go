package motorVehicleRepository

import (
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/src/motorVehicle"
	"database/sql"

	"github.com/google/uuid"
)

type motorVehicleRepository struct {
	db *sql.DB
}

func NewMotorVehicleRepository(db *sql.DB) motorVehicle.MotorVechileRepository {
	return &motorVehicleRepository{db}
}

// get all motor vehicle
// get all cust
func (mr motorVehicleRepository) RetrieveAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error) {

	query := "SELECT id, name, type, price, plat, production_year, status FROM motor_vehicle;"
	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	motors, err := scanMotorVehicle(rows)
	if err != nil {
		return nil, err
	}

	return motors, nil
}

func scanMotorVehicle(rows *sql.Rows) ([]motorVehicleDto.MotorVehicle, error) {
	var motors []motorVehicleDto.MotorVehicle

	for rows.Next() {
		motor := motorVehicleDto.MotorVehicle{}
		err := rows.Scan(&motor.Id, &motor.Name, &motor.Type, &motor.Price, &motor.Plat, &motor.CreatedAt, &motor.UpdatedAt, &motor.ProductionYear, &motor.Status, &motor.DeletedAt)
		if err != nil {
			return nil, err
		}
		motors = append(motors, motor)
	}

	err := rows.Err()
	if err != nil {
		return nil, err
	}

	return motors, nil
}

// get by id
func (mr *motorVehicleRepository) RetrieveMotorVehicleById(id uuid.UUID) (motorVehicleDto.MotorVehicle, error) {

	var motor motorVehicleDto.MotorVehicle
	query := "SELECT id, name, type, price, plat, production_year, status FROM motor_vehicle WHERE id = $1"
	if err := mr.db.QueryRow(query, id).Scan(&motor.Id, &motor.Name, &motor.Type, &motor.Price, &motor.Plat, &motor.ProductionYear, &motor.Status); err != nil {
		return motor, err
	}

	return motor, nil
}

// insert
func (mr *motorVehicleRepository) InsertMotorVehicle(motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {

	query := "INSERT INTO motor_vehicle (name, type, price, plat, production_year, status) VALUES ($1, $2, $3, $4, $5, $6);"
	_, err := mr.db.Exec(query, motor.Name, motor.Type, motor.Price, motor.Plat, motor.ProductionYear, motor.Status)
	if err != nil {
		return motor, err
	}
	return mr.RetrieveMotorVehicleById(motor.Id)
}

func (mr *motorVehicleRepository) ChangeMotorVehicle(id uuid.UUID, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {

	query := "UPDATE motor_vehicle SET name = $1, type =$2, price=$3, plat=$4, production_year=$5, status=$6;"
	_, err := mr.db.Exec(query, motor.Name, motor.Type, motor.Price, motor.Plat, motor.ProductionYear, motor.Status)
	if err != nil {
		return motor, err
	}
	return mr.RetrieveMotorVehicleById(motor.Id)
}

func (mr *motorVehicleRepository) DropMotorVehicle(id uuid.UUID) error {

	query := "DELETE FROM customer WHERE id = $1;"
	_, err := mr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
