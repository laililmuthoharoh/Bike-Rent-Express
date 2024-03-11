package transactionRepository

import (
	"bike-rent-express/model/dto/transactionDto"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var expectTransaction = transactionDto.Transaction{
	ID:             "1",
	UserID:         "1",
	MotorVehicleId: "1",
	EmployeeId:     "1",
	StartDate:      "13-09-2024",
	EndDate:        "14-09-2024",
	Price:          2000,
	CreatedAt:      "1",
	UpdatedAt:      "1",
}

func TestAddTransaction_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(10000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT (.+) FROM balance WHERE .+"
	rows = sqlmock.NewRows([]string{".+"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance"
	mock.ExpectExec(query).WithArgs(20000, expectAddTransactionRequest.UserID).WillReturnResult(sqlmock.NewResult(1, 1))

	query = "UPDATE motor_vehicle"
	mock.ExpectExec(query).WithArgs(expectAddTransactionRequest.MotorVehicleId).WillReturnResult(sqlmock.NewResult(1, 1))

	query = "INSERT INTO transaction(.+) RETURNING .+;"
	rows = sqlmock.NewRows([]string{".+"}).AddRow(expectAddTransactionRequest.ID)
	mock.ExpectQuery(query).WillReturnRows(rows)

	mock.ExpectCommit()

	actualAddTransaction, err := transactionRepository.Add(expectAddTransactionRequest)
	assert.Nil(t, err)
	assert.Equal(t, expectAddTransactionRequest, actualAddTransaction)
}

func TestAddTransaction_FailedGetPriceMotorVehicle(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	mock.ExpectQuery(query)

	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FailedGetAmountBalance(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(10000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT (.+) FROM balance WHERE .+"
	mock.ExpectQuery(query)
	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FailedSmallerPriceBalance(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(90000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT (.+) FROM balance WHERE .+"
	rows = sqlmock.NewRows([]string{".+"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FailedConvertStartDate(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "asdasdas",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()
	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FailedConvertEndDate(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "14-08-2024",
		EndDate:        "adsasdasd",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()
	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FiledDifference(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "10-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()
	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FailedUpdateBalance(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(10000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT (.+) FROM balance WHERE .+"
	rows = sqlmock.NewRows([]string{".+"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance"
	mock.ExpectExec(query).WithArgs(20000, expectAddTransactionRequest.UserID).WillReturnError(errors.New("error"))
	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_FailedUpdateMotorVehicle(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(10000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT (.+) FROM balance WHERE .+"
	rows = sqlmock.NewRows([]string{".+"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance"
	mock.ExpectExec(query).WithArgs(20000, expectAddTransactionRequest.UserID).WillReturnResult(sqlmock.NewResult(1, 1))

	query = "UPDATE motor_vehicle"
	mock.ExpectExec(query).WithArgs(expectAddTransactionRequest.MotorVehicleId).WillReturnError(errors.New("error"))
	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestAddTransaction_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	expectAddTransactionRequest := transactionDto.AddTransactionRequest{
		ID:             "123",
		UserID:         "123",
		MotorVehicleId: "123",
		EmployeeId:     "123",
		StartDate:      "13-08-2024",
		EndDate:        "14-08-2024",
	}

	transactionRepository := NewTransactionRepository(dbMock)

	mock.ExpectBegin()

	query := "SELECT (.+) FROM motor_vehicle WHERE .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(10000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "SELECT (.+) FROM balance WHERE .+"
	rows = sqlmock.NewRows([]string{".+"}).AddRow(30000)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "UPDATE balance"
	mock.ExpectExec(query).WithArgs(20000, expectAddTransactionRequest.UserID).WillReturnResult(sqlmock.NewResult(1, 1))

	query = "UPDATE motor_vehicle"
	mock.ExpectExec(query).WithArgs(expectAddTransactionRequest.MotorVehicleId).WillReturnResult(sqlmock.NewResult(1, 1))

	query = "INSERT INTO transaction(.+) RETURNING .+;"
	mock.ExpectQuery(query)

	mock.ExpectRollback()

	_, err = transactionRepository.Add(expectAddTransactionRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestGetById_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()
	transactionRepository := NewTransactionRepository(dbMock)

	query := "SELECT (.+) FROM transaction WHERE .+"
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectTransaction.ID, expectTransaction.UserID, expectTransaction.MotorVehicleId, expectTransaction.StartDate, expectTransaction.EndDate, expectTransaction.Price, expectTransaction.CreatedAt, expectTransaction.UpdatedAt, expectTransaction.EmployeeId)
	mock.ExpectQuery(query).WillReturnRows(rows)

	actualTransaction, err := transactionRepository.GetById(expectTransaction.ID)
	assert.Nil(t, err)
	assert.Equal(t, expectTransaction, actualTransaction)
}

func TestGetById_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()
	transactionRepository := NewTransactionRepository(dbMock)

	query := "SELECT (.+) FROM transaction WHERE .+"
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"})
	mock.ExpectQuery(query).WillReturnRows(rows)

	actualTransaction, err := transactionRepository.GetById(expectTransaction.ID)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectTransaction, actualTransaction)
}

func TestGetAll_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()
	expectedGetAllTranasction := []transactionDto.Transaction{
		expectTransaction,
	}
	value := [][]driver.Value{
		{
			expectTransaction.ID,
			expectTransaction.UserID,
			expectTransaction.MotorVehicleId,
			expectTransaction.StartDate,
			expectTransaction.EndDate,
			expectTransaction.Price,
			expectTransaction.CreatedAt,
			expectTransaction.UpdatedAt,
			expectTransaction.EmployeeId,
		},
	}

	transactionRepository := NewTransactionRepository(dbMock)

	query := "SELECT (.+) FROM transaction"
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRows(value...)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualGetAllTransaction, err := transactionRepository.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expectedGetAllTranasction, actualGetAllTransaction)

}

func TestGetAll_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()
	expectedGetAllTranasction := []transactionDto.Transaction{
		expectTransaction,
	}

	transactionRepository := NewTransactionRepository(dbMock)

	query := "SELECT (.+) FROM transaction"

	mock.ExpectQuery(query)

	actualGetAllTransaction, err := transactionRepository.GetAll()
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectedGetAllTranasction, actualGetAllTransaction)

}
