package motorVehicleRepository

import (
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/src/motorVehicle"
	"database/sql"
)

type motorVehicleRepository struct {
	db *sql.DB
}

func NewMotorVehicleRepository(db *sql.DB) motorVehicle.MotorVechileRepository {
	return &motorVehicleRepository{db}
}

// get all motor vehicle
func (mr motorVehicleRepository) RetrieveAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error) {
	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE deleted_at IS NULL;"
	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var motors []motorVehicleDto.MotorVehicle

	for rows.Next() {
		motor := motorVehicleDto.MotorVehicle{}
		err := rows.Scan(&motor.Id, &motor.Name, &motor.Type, &motor.Price, &motor.Plat, &motor.CreatedAt, &motor.UpdatedAt, &motor.ProductionYear, &motor.Status)
		if err != nil {
			return nil, err
		}
		motors = append(motors, motor)
	}
	return motors, nil
}

// get by id
func (mr *motorVehicleRepository) RetrieveMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error) {

	var motor motorVehicleDto.MotorVehicle
	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE id = $1 AND deleted_at IS NULL"
	if err := mr.db.QueryRow(query, id).Scan(&motor.Id, &motor.Name, &motor.Type, &motor.Price, &motor.Plat, &motor.CreatedAt, &motor.UpdatedAt, &motor.ProductionYear, &motor.Status); err != nil {
		return motor, err
	}

	return motor, nil
}

// insert
func (mr *motorVehicleRepository) InsertMotorVehicle(motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {

	query := "INSERT INTO motor_vehicle (name, type, price, plat, production_year, status) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;"
	err := mr.db.QueryRow(query, motor.Name, motor.Type, motor.Price, motor.Plat, motor.ProductionYear, motor.Status).Scan(&motor.Id)
	if err != nil {
		return motor, err
	}

	return mr.RetrieveMotorVehicleById(motor.Id)
}

func (mr *motorVehicleRepository) ChangeMotorVehicle(id string, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {

	// now := time.Now().Truncate(time.Second)

	query := "UPDATE motor_vehicle SET name = $1, type = $2, price = $3, plat = $4, production_year = $5, status = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7;"
	_, err := mr.db.Exec(query, motor.Name, motor.Type, motor.Price, motor.Plat, motor.ProductionYear, motor.Status, id)
	if err != nil {
		return motor, err
	}
	return mr.RetrieveMotorVehicleById(motor.Id)
}

func (mr *motorVehicleRepository) DropMotorVehicle(id string) error {

	query := "UPDATE motor_vehicle SET deleted_at = CURRENT_DATE WHERE id = $1 AND status = 'AVAILABLE';"
	_, err := mr.db.Exec(query, id)
	if err != nil {
		return err
	}

	return err
}
