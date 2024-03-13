package transactionDelivery

import (
	"bike-rent-express/model/dto"
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/model/dto/transactionDto"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var accessToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA0MTI0MTgsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbiIsImlkIjoiIiwicm9sZSI6IkFETUlOIn0.VL7b2UjsczLaEyLNJFYhpbUCSlbfG4tAMGpEPXi9JOk"

type mockTransactionUC struct {
	mock.Mock
}

func (m *mockTransactionUC) AddTransaction(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.Transaction, error) {
	args := m.Called(transactionRequest)
	return args.Get(0).(transactionDto.Transaction), args.Error(1)
}

func (m *mockTransactionUC) GetTransactionById(id string) (transactionDto.ResponseTransaction, error) {
	args := m.Called(id)
	return args.Get(0).(transactionDto.ResponseTransaction), args.Error(1)
}

func (m *mockTransactionUC) GetTransactionAll() ([]transactionDto.ResponseTransaction, error) {
	args := m.Called()
	return args.Get(0).([]transactionDto.ResponseTransaction), args.Error(1)
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

type TestTransactionDelierySuite struct {
	suite.Suite
	mockTransactionUC *mockTransactionUC
	router            *gin.Engine
}

func (suite *TestTransactionDelierySuite) SetupTest() {
	suite.mockTransactionUC = new(mockTransactionUC)
	suite.router = gin.Default()
	api := suite.router.Group("/api")
	v1 := api.Group("/v1")
	NewTransactionDelivery(v1, suite.mockTransactionUC)
}

func (suite *TestTransactionDelierySuite) TestCreateTransaction_Success() {
	transactionRequest := transactionDto.AddTransactionRequest{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "12-09-2024",
		EndDate:        "10-09-2024",
	}
	expectResponse := `{"responseCode":"2010101","responseMessage":"Transaction Created","data":{"id":"1","user_id":"","motor_vehicle_id":"","employee_id":"","start_date":"13-09-2024","end_date":"13-09-2025","price":20000,"created_at":"test","updated_at":"test"}}`

	suite.mockTransactionUC.On("AddTransaction", transactionRequest).Return(expectTransaction, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(transactionRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/transaction", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestCreateTransaction_FailedBind() {
	transactionRequest := transactionDto.AddTransactionRequest{
		ID:         "1",
		UserID:     "1",
		EmployeeId: "1",
		StartDate:  "12-09-2024",
		EndDate:    "10-09-2024",
	}
	expectResponse := `{"responseCode":"4000101","responseMessage":"Bad Request","error_description":[{"field":"MotorVehicleId","message":"field is required"}]}`

	suite.mockTransactionUC.On("AddTransaction", transactionRequest).Return(expectTransaction, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(transactionRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/transaction", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestCreateTransaction_Failed() {
	transactionRequest := transactionDto.AddTransactionRequest{
		ID:             "1",
		UserID:         "1",
		MotorVehicleId: "1",
		EmployeeId:     "1",
		StartDate:      "12-09-2024",
		EndDate:        "10-09-2024",
	}
	expectResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"error"}`

	suite.mockTransactionUC.On("AddTransaction", transactionRequest).Return(expectTransaction, errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(transactionRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/transaction", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestGetTransactionById_Success() {
	suite.mockTransactionUC.On("GetTransactionById", expectTransaction.ID).Return(expectTransactionResponse, nil)
	expectResponse := `{"responseCode":"2000202","responseMessage":"Success get transaction by id","data":{"id":"1","start_date":"13-09-2024","end_date":"13-09-2025","price":20000,"motor_vehicle":{"id":"1","name":"test","type":"test","price":2000,"plat":"test","created_at":"test","updated_at":"test","production_year":"2020","status":"AVAILABLE"},"employee":{"id":"1","name":"test","telp":"08123","username":"test","created_at":"test","updated_at":"test"},"customer":{"id":"1","nama":"test","username":"test","alamat":"test","role":"USER","cant_rent":true,"created_at":"test","updated_at":"test","telepon":"0812312"},"created_at":"test","updated_at":"test"}}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/transaction/"+expectTransaction.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestGetTransactionById_FailedDataNotFound() {
	suite.mockTransactionUC.On("GetTransactionById", expectTransaction.ID).Return(expectTransactionResponse, errors.New("1"))
	expectResponse := `{"responseCode":"2000201","responseMessage":"Data not found"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/transaction/"+expectTransaction.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestGetTransactionById_Failed() {
	suite.mockTransactionUC.On("GetTransactionById", expectTransaction.ID).Return(expectTransactionResponse, errors.New("error"))
	expectResponse := `{"responseCode":"5000201","responseMessage":"internal server error","error":"error"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/transaction/"+expectTransaction.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestGetTransactionAll_Success() {
	allResponseTransaction := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	suite.mockTransactionUC.On("GetTransactionAll").Return(allResponseTransaction, nil)
	expectResponse := `{"responseCode":"2000202","responseMessage":"Success get all transaction","data":[{"id":"1","start_date":"13-09-2024","end_date":"13-09-2025","price":20000,"motor_vehicle":{"id":"1","name":"test","type":"test","price":2000,"plat":"test","created_at":"test","updated_at":"test","production_year":"2020","status":"AVAILABLE"},"employee":{"id":"1","name":"test","telp":"08123","username":"test","created_at":"test","updated_at":"test"},"customer":{"id":"1","nama":"test","username":"test","alamat":"test","role":"USER","cant_rent":true,"created_at":"test","updated_at":"test","telepon":"0812312"},"created_at":"test","updated_at":"test"}]}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/transaction", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *TestTransactionDelierySuite) TestGetTransactionAll_Failed() {
	allResponseTransaction := []transactionDto.ResponseTransaction{
		expectTransactionResponse,
	}

	suite.mockTransactionUC.On("GetTransactionAll").Return(allResponseTransaction, errors.New("error"))
	expectResponse := `{"responseCode":"5000301","responseMessage":"internal server error","error":"error"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/transaction", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())

}

func (suite *TestTransactionDelierySuite) TestGetTransactionAll_FailedEmpty() {
	allResponseTransaction := []transactionDto.ResponseTransaction{}

	suite.mockTransactionUC.On("GetTransactionAll").Return(allResponseTransaction, nil)
	expectResponse := `{"responseCode":"2000201","responseMessage":"Data empty"}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/transaction", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func TestTransactionDelivery(t *testing.T) {
	suite.Run(t, new(TestTransactionDelierySuite))
}
