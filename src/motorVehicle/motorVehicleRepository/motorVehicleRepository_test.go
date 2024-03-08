package motorVehicleRepository

import (
	"bike-rent-express/model/dto/motorVehicleDto"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var expectedAllMotorVehicle = []motorVehicleDto.MotorVehicle{
	{
		Id:             "3a57713c-24d0-41f8-bfa8-f8f721dba9e4",
		Name:           "Vario",
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		CreatedAt:      "2024-03-07T00:00:00Z",
		UpdatedAt:      "2024-03-07T00:00:00Z",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	},
	{
		Id:             "3a57713c-24d0-41f8-bfa8-f8f721dba9e5",
		Name:           "NMAX",
		Type:           "MATIC",
		Price:          50000,
		Plat:           "AB1234I",
		CreatedAt:      "2024-03-07T00:00:00Z",
		UpdatedAt:      "2024-03-07T00:00:00Z",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	},
}

var expectedMotorVehicleById = motorVehicleDto.MotorVehicle{
	Id:             "3a57713c-24d0-41f8-bfa8-f8f721dba9e4",
	Name:           "Vario",
	Type:           "MATIC",
	Price:          50000,
	Plat:           "BA1234I",
	CreatedAt:      "2024-03-07T00:00:00Z",
	UpdatedAt:      "2024-03-07T00:00:00Z",
	ProductionYear: "2023",
	Status:         "AVAILABLE",
}

var expectedCreatedMotorVehicle = motorVehicleDto.CreateMotorVehicle{
	Name:           "Vario",
	Type:           "MATIC",
	Price:          50000,
	Plat:           "BA1234I",
	ProductionYear: "2023",
	Status:         "AVAILABLE",
}

// test get all success
func TestRetrieveAllMotorVehicle_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	//mock database
	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE deleted_at IS NULL;"
	rows := mock.NewRows([]string{"id", "name", "type", "price", "plat", "created_at", "updated_at", "production_year", "status"}).
		AddRow(expectedAllMotorVehicle[0].Id, expectedAllMotorVehicle[0].Name, expectedAllMotorVehicle[0].Type, expectedAllMotorVehicle[0].Price, expectedAllMotorVehicle[0].Plat, expectedAllMotorVehicle[0].CreatedAt, expectedAllMotorVehicle[0].UpdatedAt, expectedAllMotorVehicle[0].ProductionYear, expectedAllMotorVehicle[0].Status).
		AddRow(expectedAllMotorVehicle[1].Id, expectedAllMotorVehicle[1].Name, expectedAllMotorVehicle[1].Type, expectedAllMotorVehicle[1].Price, expectedAllMotorVehicle[1].Plat, expectedAllMotorVehicle[1].CreatedAt, expectedAllMotorVehicle[0].UpdatedAt, expectedAllMotorVehicle[1].ProductionYear, expectedAllMotorVehicle[1].Status)

	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repository.RetrieveAllMotorVehicle()

	assert.Nil(t, err)
	assert.Equal(t, expectedAllMotorVehicle, result)
}

// tes get all (if error when db.Query)
func TestRetrieveAllMotorVehicle_Fail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	//mock database
	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE deleted_at IS NULL;"

	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	result, err := repository.RetrieveAllMotorVehicle()

	assert.Error(t, err)
	assert.Empty(t, result)
}

// tes all empty (if error when empty) --> targetnya buat cover error .Scan
func TestRetrieveAllMotorVehicle_Fail2(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	//mock database
	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE deleted_at IS NULL;"

	mock.ExpectQuery(query)

	result, err := repository.RetrieveAllMotorVehicle()

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectedAllMotorVehicle, result)

}

// test get by id
func TestRetrieveMotorVehicleById_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE id = \\$1 AND deleted_at IS NULL"
	rows := mock.NewRows([]string{"id", "name", "type", "price", "plat", "created_at", "updated_at", "production_year", "status"}).
		AddRow(expectedMotorVehicleById.Id, expectedMotorVehicleById.Name, expectedMotorVehicleById.Type, expectedMotorVehicleById.Price, expectedMotorVehicleById.Plat, expectedMotorVehicleById.CreatedAt, expectedMotorVehicleById.UpdatedAt, expectedMotorVehicleById.ProductionYear, expectedMotorVehicleById.Status)

	mock.ExpectQuery(query).WillReturnRows(rows)

	result, err := repository.RetrieveMotorVehicleById(expectedMotorVehicleById.Id)

	assert.Nil(t, err)
	assert.Equal(t, expectedMotorVehicleById, result)

}

// test get by id fail
func TestRetrieveMotorVehicleById_Fail(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	query := "SELECT id, name, type, price, plat, created_at, updated_at, production_year, status FROM motor_vehicle WHERE id = $1 AND deleted_at IS NULL"

	mock.ExpectQuery(query).WillReturnError(errors.New("error sql"))

	result, err := repository.RetrieveMotorVehicleById(expectedMotorVehicleById.Id)

	assert.Error(t, err)
	assert.Empty(t, result)

}

// test insert motor vehicle success
func TestInsertMotorVehicle_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	query := "INSERT INTO motor_vehicle (name, type, price, plat, production_year, status) VALUES (.+) RETURNING id;"

	//source: https://github.com/DATA-DOG/go-sqlmock/issues/27
	mock.ExpectQuery(query).WithArgs(expectedMotorVehicleById.Name, expectedMotorVehicleById.Type, expectedMotorVehicleById.Price, expectedMotorVehicleById.Plat, expectedMotorVehicleById.ProductionYear, expectedMotorVehicleById.Status).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedMotorVehicleById.Id))

	result, err := repository.InsertMotorVehicle(expectedMotorVehicleById)

	assert.Nil(t, err)
	assert.Equal(t, expectedMotorVehicleById, result)
}

// test update motor vehicle success
func TestChangeMotorVehicle_Success(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error creating mock database: ", err)
	}
	defer db.Close()

	//initialization repository
	repository := NewMotorVehicleRepository(db)

	query := "UPDATE motor_vehicle SET name = $1, type = $2, price = $3, plat = $4, production_year = $5, status = $6, updated_at = $7 WHERE id = $8;"

	//source: https://github.com/DATA-DOG/go-sqlmock/issues/27
	mock.ExpectExec(query).WithArgs(expectedMotorVehicleById.Name, expectedMotorVehicleById.Type, expectedMotorVehicleById.Price, expectedMotorVehicleById.Plat, expectedMotorVehicleById.ProductionYear, expectedMotorVehicleById.Status).
		WillReturnResult(sqlmock.NewRows([]string{"id"}).AddRow(expectedMotorVehicleById.Id))

	result, err := repository.InsertMotorVehicle(expectedMotorVehicleById)

	assert.Nil(t, err)
	assert.Equal(t, expectedMotorVehicleById, result)
}
