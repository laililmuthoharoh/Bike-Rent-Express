package transactionUsecase

import (
	"bike-rent-express/model/dto"
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/model/dto/transactionDto"
	"bike-rent-express/src/transaction"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockTransactionRepository struct {
	mock.Mock
}

func (m *mockTransactionRepository) Add(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.AddTransactionRequest, error) {
	args := m.Called(transactionRequest)
	return args.Get(0).(transactionDto.AddTransactionRequest), args.Error(1)
}

func (m *mockTransactionRepository) GetById(id string) (transactionDto.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(transactionDto.Transaction), args.Error(1)
}

func (m *mockTransactionRepository) GetAll() ([]transactionDto.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]transactionDto.Transaction), args.Error(1)
}

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) RegisterUsers(newUsers dto.RegisterUsers) error {
	args := m.Called(newUsers)
	return args.Error(0)
}

func (m *mockUserRepository) GetByID(id string) (dto.GetUsers, error) {
	args := m.Called(id)
	return args.Get(0).(dto.GetUsers), args.Error(1)
}
func (m *mockUserRepository) GetAll() ([]dto.GetUsers, error) {
	args := m.Called()
	return args.Get(0).([]dto.GetUsers), args.Error(1)
}
func (m *mockUserRepository) UpdateUsers(usersItem dto.Users) error {
	args := m.Called(usersItem)
	return args.Error(0)
}
func (m *mockUserRepository) GetByUsername(username string) (dto.Users, error) {
	args := m.Called(username)
	return args.Get(0).(dto.Users), args.Error(1)
}
func (m *mockUserRepository) UpdateBalance(topUpRequest dto.TopUpRequest) error {
	args := m.Called(topUpRequest)
	return args.Error(0)
}
func (m *mockUserRepository) UpdatePassword(changePasswordRequest dto.ChangePassword) error {
	args := m.Called(changePasswordRequest)
	return args.Error(0)
}
func (m *mockUserRepository) UsernameIsReady(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func (m *mockUserRepository) GetBalance(id string) (dto.Balance, error) {
	args := m.Called(id)
	return args.Get(0).(dto.Balance), args.Error(1)
}

type mockEmployeeRepository struct {
	mock.Mock
}

func (m *mockEmployeeRepository) Add(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error) {
	args := m.Called(employee)
	return args.Get(0).(employeeDto.CreateEmployeeRequest), args.Error(1)
}
func (m *mockEmployeeRepository) Get() ([]employeeDto.Employee, error) {
	args := m.Called()
	return args.Get(0).([]employeeDto.Employee), args.Error(1)
}
func (m *mockEmployeeRepository) UsernameIsReady(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}
func (m *mockEmployeeRepository) GetByUsername(username string) (employeeDto.Employee, error) {
	args := m.Called(username)
	return args.Get(0).(employeeDto.Employee), args.Error(1)
}
func (m *mockEmployeeRepository) GetById(id string) (employeeDto.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(employeeDto.Employee), args.Error(1)
}
func (m *mockEmployeeRepository) Update(employeeUpdateRequest employeeDto.UpdateEmployeeRequest) (employeeDto.Employee, error) {
	args := m.Called(employeeUpdateRequest)
	return args.Get(0).(employeeDto.Employee), args.Error(1)
}
func (m *mockEmployeeRepository) UpdatePassword(employee employeeDto.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}
func (m *mockEmployeeRepository) Delete(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

type mockMotorVehicleRepository struct {
	mock.Mock
}

func (m *mockMotorVehicleRepository) RetrieveAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error) {
	args := m.Called()
	return args.Get(0).([]motorVehicleDto.MotorVehicle), args.Error(1)
}
func (m *mockMotorVehicleRepository) RetrieveMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error) {
	args := m.Called(id)
	return args.Get(0).(motorVehicleDto.MotorVehicle), args.Error(1)
}
func (m *mockMotorVehicleRepository) InsertMotorVehicle(motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	args := m.Called(motor)
	return args.Get(0).(motorVehicleDto.MotorVehicle), args.Error(1)
}
func (m *mockMotorVehicleRepository) ChangeMotorVehicle(id string, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	args := m.Called(id, motor)
	return args.Get(0).(motorVehicleDto.MotorVehicle), args.Error(1)
}
func (m *mockMotorVehicleRepository) DropMotorVehicle(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockMotorVehicleRepository) CheckPlatMotor(plat string) (bool, error) {
	args := m.Called(plat)
	return args.Bool(0), args.Error(1)
}

var expectMotorVehicle = motorVehicleDto.MotorVehicle{
	Id:             "1",
	Name:           "test",
	Type:           "test",
	Price:          2000,
	Plat:           "test",
	CreatedAt:      "test",
	UpdatedAt:      "test",
	ProductionYear: "2020",
	Status:         "AVAILABLE",
}
var expectEmployee = employeeDto.Employee{
	ID:        "1",
	Name:      "test",
	Telp:      "08123",
	Username:  "test",
	Password:  "test",
	CreatedAt: "test",
	UpdatedAt: "test",
}
var expectCustomer = dto.GetUsers{
	Uuid:       "1",
	Name:       "test",
	Username:   "test",
	Address:    "test",
	Role:       "USER",
	Can_rent:   true,
	Password:   "test",
	Created_at: "test",
	Updated_at: "test",
	Telp:       "0812312",
}
var expectTransaction = transactionDto.Transaction{
	ID:        "1",
	StartDate: "13-09-2024",
	EndDate:   "13-09-2025",
	Price:     20000,
	CreatedAt: "test",
	UpdatedAt: "test",
}
var expectTransactionResponse = transactionDto.ResponseTransaction{
	ID:           "1",
	StartDate:    "13-09-2024",
	EndDate:      "13-09-2025",
	Price:        20000,
	MotorVehicle: expectMotorVehicle,
	Employee:     expectEmployee,
	Customer:     expectCustomer,
	CreatedAt:    "test",
	UpdatedAt:    "test",
}

type TransactionUseCaseSuite struct {
	suite.Suite
	mockTransactionRepository  *mockTransactionRepository
	mockUserRepository         *mockUserRepository
	mockEmployeeRepository     *mockEmployeeRepository
	mockMotorVehicleRepository *mockMotorVehicleRepository
	transactionUsecase         transaction.TransactionUsecase
}

func (suite *TransactionUseCaseSuite) SetupTest() {
	suite.mockTransactionRepository = new(mockTransactionRepository)
	suite.mockUserRepository = new(mockUserRepository)
	suite.mockEmployeeRepository = new(mockEmployeeRepository)
	suite.mockMotorVehicleRepository = new(mockMotorVehicleRepository)
	suite.transactionUsecase = NewTransactionRepository(suite.mockTransactionRepository, suite.mockUserRepository, suite.mockEmployeeRepository, suite.mockMotorVehicleRepository)
}

func (suite *TransactionUseCaseSuite) TestAddTransaction_Success() {
	addTransaction := transactionDto.AddTransactionRequest{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "13-09-2023",
		EndDate:        "15-09-2024",
	}
	expectTransaction := transactionDto.Transaction{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "13-09-2023",
		EndDate:        "13-09-2024",
		Price:          2000,
		CreatedAt:      "123",
		UpdatedAt:      "123",
	}

	suite.mockTransactionRepository.On("Add", addTransaction).Return(addTransaction, nil)
	suite.mockTransactionRepository.On("GetById", addTransaction.ID).Return(expectTransaction, nil)

	actualTransaction, err := suite.transactionUsecase.AddTransaction(addTransaction)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectTransaction, actualTransaction)
}

func (suite *TransactionUseCaseSuite) TestAddTransaction_FailedAdd() {
	addTransaction := transactionDto.AddTransactionRequest{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "13-09-2023",
		EndDate:        "15-09-2024",
	}
	expectTransaction := transactionDto.Transaction{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "13-09-2023",
		EndDate:        "13-09-2024",
		Price:          2000,
		CreatedAt:      "123",
		UpdatedAt:      "123",
	}
	expectError := errors.New("error")

	suite.mockTransactionRepository.On("Add", addTransaction).Return(addTransaction, expectError)
	suite.mockTransactionRepository.On("GetById", addTransaction.ID).Return(expectTransaction, nil)

	actualTransaction, err := suite.transactionUsecase.AddTransaction(addTransaction)
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransaction, actualTransaction)
}

func (suite *TransactionUseCaseSuite) TestAddTransaction_FailedGetById() {
	addTransaction := transactionDto.AddTransactionRequest{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "13-09-2023",
		EndDate:        "15-09-2024",
	}
	expectTransaction := transactionDto.Transaction{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "13-09-2023",
		EndDate:        "13-09-2024",
		Price:          2000,
		CreatedAt:      "123",
		UpdatedAt:      "123",
	}
	expectError := errors.New("error")

	suite.mockTransactionRepository.On("Add", addTransaction).Return(addTransaction, nil)
	suite.mockTransactionRepository.On("GetById", addTransaction.ID).Return(transactionDto.Transaction{}, expectError)

	actualTransaction, err := suite.transactionUsecase.AddTransaction(addTransaction)
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransaction, actualTransaction)
}

func (suite *TransactionUseCaseSuite) TestGetTransactionById_Success() {
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, nil)
	suite.mockEmployeeRepository.On("GetById", expectTransaction.EmployeeId).Return(expectEmployee, nil)
	suite.mockUserRepository.On("GetByID", expectTransaction.UserID).Return(expectCustomer, nil)

	actualExpectTransactionResponse, err := suite.transactionUsecase.GetTransactionById(expectTransaction.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectTransactionResponse, actualExpectTransactionResponse)
}

func (suite *TransactionUseCaseSuite) TestGetTransactionById_FailedGetByIdInvalidInputOrSqlNoRows() {
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, sql.ErrNoRows)

	actualExpectTransactionResponse, err := suite.transactionUsecase.GetTransactionById(expectTransaction.ID)
	assert.NotNil(suite.T(), err)
	assert.NotEqual(suite.T(), expectTransactionResponse, actualExpectTransactionResponse)
}

func (suite *TransactionUseCaseSuite) TestGetTransactionById_FailedGetById() {
	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, expectError)

	actualExpectTransactionResponse, err := suite.transactionUsecase.GetTransactionById(expectTransaction.ID)
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionResponse, actualExpectTransactionResponse)
}

func (suite *TransactionUseCaseSuite) TestGetTransactionById_FailedGetMotorVehicle() {
	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, expectError)

	actualExpectTransactionResponse, err := suite.transactionUsecase.GetTransactionById(expectTransaction.ID)
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionResponse, actualExpectTransactionResponse)
}

func (suite *TransactionUseCaseSuite) TestGetTransactionById_FailedGetEmployee() {
	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, nil)
	suite.mockEmployeeRepository.On("GetById", expectTransaction.EmployeeId).Return(expectEmployee, expectError)

	actualExpectTransactionResponse, err := suite.transactionUsecase.GetTransactionById(expectTransaction.ID)
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionResponse, actualExpectTransactionResponse)
}

func (suite *TransactionUseCaseSuite) TestGetTransactionById_FailedGetCustomer() {
	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, nil)
	suite.mockEmployeeRepository.On("GetById", expectTransaction.EmployeeId).Return(expectEmployee, nil)
	suite.mockUserRepository.On("GetByID", expectTransaction.UserID).Return(expectCustomer, expectError)

	actualExpectTransactionResponse, err := suite.transactionUsecase.GetTransactionById(expectTransaction.ID)
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionResponse, actualExpectTransactionResponse)
}

func (suite *TransactionUseCaseSuite) TestGetAllTransaction_Success() {
	allTransaction := []transactionDto.Transaction{
		expectTransaction,
	}

	expectTransactionGetAllResponse := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	suite.mockTransactionRepository.On("GetAll").Return(allTransaction, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, nil)
	suite.mockEmployeeRepository.On("GetById", expectTransaction.EmployeeId).Return(expectEmployee, nil)
	suite.mockUserRepository.On("GetByID", expectTransaction.UserID).Return(expectCustomer, nil)

	actualTransactionGetAllResponse, err := suite.transactionUsecase.GetTransactionAll()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectTransactionGetAllResponse, actualTransactionGetAllResponse)
}

func (suite *TransactionUseCaseSuite) TestGetAllTransaction_FailedGetAll() {
	allTransaction := []transactionDto.Transaction{
		expectTransaction,
	}

	expectTransactionGetAllResponse := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}
	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetAll").Return(allTransaction, expectError)

	actualTransactionGetAllResponse, err := suite.transactionUsecase.GetTransactionAll()
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionGetAllResponse, actualTransactionGetAllResponse)
}

func (suite *TransactionUseCaseSuite) TestGetAllTransaction_FailedTransactionGetById() {
	allTransaction := []transactionDto.Transaction{
		expectTransaction,
	}

	expectTransactionGetAllResponse := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetAll").Return(allTransaction, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, expectError)

	actualTransactionGetAllResponse, err := suite.transactionUsecase.GetTransactionAll()
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionGetAllResponse, actualTransactionGetAllResponse)
}

func (suite *TransactionUseCaseSuite) TestGetAllTransaction_FailedGetMotorById() {
	allTransaction := []transactionDto.Transaction{
		expectTransaction,
	}

	expectTransactionGetAllResponse := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetAll").Return(allTransaction, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, expectError)

	actualTransactionGetAllResponse, err := suite.transactionUsecase.GetTransactionAll()
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionGetAllResponse, actualTransactionGetAllResponse)
}

func (suite *TransactionUseCaseSuite) TestGetAllTransaction_FailedGetEmployeeId() {
	allTransaction := []transactionDto.Transaction{
		expectTransaction,
	}

	expectTransactionGetAllResponse := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetAll").Return(allTransaction, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, nil)
	suite.mockEmployeeRepository.On("GetById", expectTransaction.EmployeeId).Return(expectEmployee, expectError)

	actualTransactionGetAllResponse, err := suite.transactionUsecase.GetTransactionAll()
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionGetAllResponse, actualTransactionGetAllResponse)
}

func (suite *TransactionUseCaseSuite) TestGetAllTransaction_FailedCustomerGetById() {
	allTransaction := []transactionDto.Transaction{
		expectTransaction,
	}

	expectTransactionGetAllResponse := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	expectError := errors.New("error")
	suite.mockTransactionRepository.On("GetAll").Return(allTransaction, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockMotorVehicleRepository.On("RetrieveMotorVehicleById", expectTransaction.MotorVehicleId).Return(expectMotorVehicle, nil)
	suite.mockEmployeeRepository.On("GetById", expectTransaction.EmployeeId).Return(expectEmployee, nil)
	suite.mockUserRepository.On("GetByID", expectTransaction.UserID).Return(expectCustomer, expectError)

	actualTransactionGetAllResponse, err := suite.transactionUsecase.GetTransactionAll()
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectTransactionGetAllResponse, actualTransactionGetAllResponse)
}

func TestTransactionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseSuite))
}
