package motorReturnRepository

import (
	"bike-rent-express/model/dto/motorReturnDto"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var expectedMotorReturn = motorReturnDto.MotorReturn{
	ID:             "907698c8-ae04-47b2-a7b9-68c46690c3f8",
	TrasactionID:   "621dfcb6-06df-4420-b98e-3ec04def9547",
	ReturnDate:     "2024-03-07T00:00:00Z",
	ExtraCharge:    25000,
	ConditionMotor: "Ban depan bocor",
	Descrption:     "bocor di jalan",
	CreatedAt:      "2024-03-07T23:39:42.63419Z",
	UpdatedAt:      "2024-03-07T23:39:42.63419Z",
}

var expectedCreateMotorReturn = motorReturnDto.CreateMotorReturnRequest{
	ID:             expectedMotorReturn.ID,
	TransactionID:  expectedMotorReturn.TrasactionID,
	ExtraCharge:    expectedMotorReturn.ExtraCharge,
	ConditionMotor: expectedMotorReturn.ConditionMotor,
	Description:    expectedMotorReturn.Descrption,
}

// test add success
func TestAdd_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow("907698c8-ae04-47b2-a7b9-68c46690c3f8")
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT amount FROM balance WHERE user_id = \\$1;"
	rows = sqlmock.NewRows([]string{"amount"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance SET amount = \\$1 WHERE user_id = \\$2;"
	mock.ExpectExec(query).WithArgs(5000, "907698c8-ae04-47b2-a7b9-68c46690c3f8").WillReturnResult(sqlmock.NewResult(0, 1))

	query = "UPDATE motor_vehicle SET status = 'AVAILABLE';"
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))

	query = "INSERT INTO motor_return(.+) RETURNING id;"
	rows = sqlmock.NewRows([]string{"id"}).AddRow(expectedCreateMotorReturn.ID)
	mock.ExpectQuery(query).WillReturnRows(rows)

	mock.ExpectCommit()

	result, err := repository.Add(expectedCreateMotorReturn)
	assert.Nil(t, err)
	assert.Equal(t, expectedCreateMotorReturn, result)

}

// test fail to get user_id from transaction
func TestAdd_FailToGetUserIdTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	mock.ExpectRollback()

	_, err = repository.Add(expectedCreateMotorReturn)
	assert.Error(t, err)
}

// test fail to get amount from balance
func TestAdd_FailToGetAmountBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow("907698c8-ae04-47b2-a7b9-68c46690c3f8")
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT amount FROM balance WHERE user_id = \\$1;"
	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	mock.ExpectRollback()

	_, err = repository.Add(expectedCreateMotorReturn)
	assert.Error(t, err)
}

// test fail when balance less than extra charge
func TestAdd_FailBalanceLessThanExtraCharge(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow("907698c8-ae04-47b2-a7b9-68c46690c3f8")
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT amount FROM balance WHERE user_id = \\$1;"
	rows = sqlmock.NewRows([]string{"amount"}).AddRow(10000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	mock.ExpectRollback()

	_, err = repository.Add(expectedCreateMotorReturn)
	assert.Error(t, err)
}

// test fail to update balance
func TestAdd_FailToUpdateBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow("907698c8-ae04-47b2-a7b9-68c46690c3f8")
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT amount FROM balance WHERE user_id = \\$1;"
	rows = sqlmock.NewRows([]string{"amount"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance SET amount = \\$1 WHERE user_id = \\$2;"
	mock.ExpectExec(query).WillReturnError(errors.New("error sql"))

	mock.ExpectRollback()

	_, err = repository.Add(expectedCreateMotorReturn)
	assert.Error(t, err)
}

// test fail to update Motor Vehicle
func TestAdd_FailToUpdateMotorVehicle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow("907698c8-ae04-47b2-a7b9-68c46690c3f8")
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT amount FROM balance WHERE user_id = \\$1;"
	rows = sqlmock.NewRows([]string{"amount"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance SET amount = \\$1 WHERE user_id = \\$2;"
	mock.ExpectExec(query).WithArgs(5000, "907698c8-ae04-47b2-a7b9-68c46690c3f8").WillReturnResult(sqlmock.NewResult(0, 1))

	query = "UPDATE motor_vehicle SET status = 'AVAILABLE';"
	mock.ExpectExec(query).WillReturnError(errors.New("error sql"))

	mock.ExpectRollback()

	_, err = repository.Add(expectedCreateMotorReturn)
	assert.Error(t, err)
}

// test fail add motor return
func TestAdd_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	mock.ExpectBegin()

	query := "SELECT user_id FROM transaction WHERE id = \\$1;"
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow("907698c8-ae04-47b2-a7b9-68c46690c3f8")
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT amount FROM balance WHERE user_id = \\$1;"
	rows = sqlmock.NewRows([]string{"amount"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance SET amount = \\$1 WHERE user_id = \\$2;"
	mock.ExpectExec(query).WithArgs(5000, "907698c8-ae04-47b2-a7b9-68c46690c3f8").WillReturnResult(sqlmock.NewResult(0, 1))

	query = "UPDATE motor_vehicle SET status = 'AVAILABLE';"
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))

	query = "INSERT INTO motor_return(.+) RETURNING id;"
	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	mock.ExpectRollback()

	_, err = repository.Add(expectedCreateMotorReturn)
	assert.Error(t, err)
}

// test get all success
func TestGetAll_Success(t *testing.T) {

	expected := []motorReturnDto.MotorReturn{
		expectedMotorReturn,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	//mock database
	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return;"
	rows := mock.NewRows([]string{"id", "transaction_id", "return_date", "extra_charge", "condition_motor", "description", "created_at", "updatad_at"}).
		AddRow(expected[0].ID, expected[0].TrasactionID, expected[0].ReturnDate, expected[0].ExtraCharge, expected[0].ConditionMotor, expected[0].Descrption, expected[0].CreatedAt, expected[0].UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

// tes get all (if error when db.Query)
func TestGetAll_Fail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	//mock database
	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return;"

	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	result, err := repository.GetAll()

	assert.Error(t, err)
	assert.Empty(t, result)
}

// tes get all fail to rows.Scan
// source: https://github.com/DATA-DOG/go-sqlmock/issues/47
func TestGetAll_FailRowsScan(t *testing.T) {

	expected := []motorReturnDto.MotorReturn{
		expectedMotorReturn,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	//mock database
	// mengubah input id menjadi nil sehingga nantinya id tidak akan terbaca
	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return;"
	rows := mock.NewRows([]string{"id", "transaction_id", "return_date", "extra_charge", "condition_motor", "description", "created_at", "updatad_at"}).
		AddRow(nil, expected[0].TrasactionID, expected[0].ReturnDate, expected[0].ExtraCharge, expected[0].ConditionMotor, expected[0].Descrption, expected[0].CreatedAt, expected[0].UpdatedAt).RowError(2, errors.New("scanErr"))

	mock.ExpectQuery(query).WillReturnRows(rows)

	_, err = repository.GetAll()

	assert.Error(t, err)
}

// test get by id
func TestGetById_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return WHERE id = \\$1;"

	rows := mock.NewRows([]string{"id", "transaction_id", "return_date", "extra_charge", "condition_motor", "description", "created_at", "updatad_at"}).
		AddRow(expectedMotorReturn.ID, expectedMotorReturn.TrasactionID, expectedMotorReturn.ReturnDate, expectedMotorReturn.ExtraCharge, expectedMotorReturn.ConditionMotor, expectedMotorReturn.Descrption, expectedMotorReturn.CreatedAt, expectedMotorReturn.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repository.GetById(expectedMotorReturn.ID)

	assert.Nil(t, err)
	assert.Equal(t, expectedMotorReturn, result)

}

// test get by id fail
func TestGetById_Fail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorRepository(db)

	query := "SELECT id, transaction_id, return_date, extra_charge, condition_motor, description, created_at, updated_at FROM motor_return WHERE id = \\$1;"

	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	result, err := repository.GetById(expectedMotorReturn.ID)

	assert.Error(t, err)
	assert.Empty(t, result)

}
