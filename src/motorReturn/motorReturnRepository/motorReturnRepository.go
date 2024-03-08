package motorReturnRepository

import (
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/src/motorReturn"
	"database/sql"
	"errors"
	"time"
)

type motorReturnRepository struct {
	db *sql.DB
}

func NewMotorRepository(db *sql.DB) motorReturn.MotorReturnRepository {
	return &motorReturnRepository{db}
}

func (m *motorReturnRepository) Add(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return createMotorReturnRequest, err
	}

	var userId string
	query := "SELECT user_id FROM transaction WHERE id = $1;"
	if err := tx.QueryRow(query, createMotorReturnRequest.TransactionID).Scan(&userId); err != nil {
		tx.Rollback()
		return createMotorReturnRequest, err
	}

	var balanceUser int
	query = "SELECT amount FROM balance WHERE user_id = $1;"

	if err := tx.QueryRow(query, userId).Scan(&balanceUser); err != nil {
		tx.Rollback()
		return createMotorReturnRequest, err
	}

	if balanceUser < createMotorReturnRequest.ExtraCharge {
		tx.Rollback()
		return createMotorReturnRequest, errors.New("1")
	}

	balanceUser -= createMotorReturnRequest.ExtraCharge

	query = "UPDATE balance SET amount = $1 WHERE user_id = $2;"
	_, err = tx.Exec(query, balanceUser, userId)
	if err != nil {
		tx.Rollback()
		return createMotorReturnRequest, err
	}

	query = "UPDATE motor_vehicle SET status = 'AVAILABLE';"
	_, err = tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return createMotorReturnRequest, err
	}

	returnDate := time.Now()

	query = "INSERT INTO motor_return(transaction_id, return_date, extra_charge, condition_motor, description) VALUES($1, $2, $3, $4, $5) RETURNING id;"

	if err := tx.QueryRow(query, createMotorReturnRequest.TransactionID, returnDate, createMotorReturnRequest.ExtraCharge, createMotorReturnRequest.ConditionMotor, createMotorReturnRequest.Description).Scan(&createMotorReturnRequest.ID); err != nil {
		tx.Rollback()
		return createMotorReturnRequest, err
	}
	tx.Commit()

	return createMotorReturnRequest, nil
}

func (m *motorReturnRepository) GetById(id string) (motorReturnDto.MotorReturn, error) {
	var motorReturn motorReturnDto.MotorReturn
	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return WHERE id = $1;"

	if err := m.db.QueryRow(query, id).Scan(&motorReturn.ID, &motorReturn.TrasactionID, &motorReturn.ReturnDate, &motorReturn.ExtraCharge, &motorReturn.ConditionMotor, &motorReturn.Descrption, &motorReturn.CreatedAt, &motorReturn.UpdatedAt); err != nil {
		return motorReturn, err
	}

	return motorReturn, nil
}

func (m *motorReturnRepository) GetAll() ([]motorReturnDto.MotorReturn, error) {
	var motorsReturn []motorReturnDto.MotorReturn

	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return;"
	rows, err := m.db.Query(query)
	if err != nil {
		return motorsReturn, err
	}

	for rows.Next() {
		var motorReturn motorReturnDto.MotorReturn
		if err := rows.Scan(&motorReturn.ID, &motorReturn.TrasactionID, &motorReturn.ReturnDate, &motorReturn.ExtraCharge, &motorReturn.ConditionMotor, &motorReturn.Descrption, &motorReturn.CreatedAt, &motorReturn.UpdatedAt); err != nil {
			return motorsReturn, err
		}
		motorsReturn = append(motorsReturn, motorReturn)
	}

	return motorsReturn, nil
}
