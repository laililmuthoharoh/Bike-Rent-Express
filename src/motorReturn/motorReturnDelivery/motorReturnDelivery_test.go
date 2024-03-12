package motorReturnDelivery

import (
	"bike-rent-express/model/dto"
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/model/dto/transactionDto"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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

var expectedCustomer = dto.GetUsers{
	Uuid:       "f4884dfc-7ef3-4d84-b77e-4fb930069da5",
	Name:       "billkin",
	Username:   "billkin",
	Address:    "Bekasi",
	Role:       "USER",
	Can_rent:   true,
	Password:   "tempo",
	Created_at: "2024-03-07T23:39:42.63419Z",
	Updated_at: "2024-03-07T23:39:42.63419Z",
	Telp:       "08123456789",
}
var expectedMotorReturnResponse = motorReturnDto.MotorReturnResponse{
	ID:             expectedMotorReturn.ID,
	ReturnDate:     expectedMotorReturn.ReturnDate,
	ExtraCharge:    expectedMotorReturn.ExtraCharge,
	ConditionMotor: expectedMotorReturn.ConditionMotor,
	Descrption:     expectedMotorReturn.Descrption,
	Customer:       expectedCustomer,
	CreatedAt:      expectedMotorReturn.CreatedAt,
	UpdatedAt:      expectedMotorReturn.UpdatedAt,
}

var expectTransaction = transactionDto.Transaction{
	ID:             expectedMotorReturn.TrasactionID,
	UserID:         expectedCustomer.Uuid,
	MotorVehicleId: "1",
	EmployeeId:     "1",
	StartDate:      "13-09-2023",
	EndDate:        "13-09-2024",
	Price:          2000,
	CreatedAt:      "123",
	UpdatedAt:      "123",
}

var expectedAllMotorReturnResponse = []motorReturnDto.MotorReturnResponse{
	expectedMotorReturnResponse,
}

var expectedCreateMotorReturn = motorReturnDto.CreateMotorReturnRequest{
	ID:             expectedMotorReturn.ID,
	TransactionID:  expectedMotorReturn.TrasactionID,
	ExtraCharge:    expectedMotorReturn.ExtraCharge,
	ConditionMotor: expectedMotorReturn.ConditionMotor,
	Description:    expectedMotorReturn.Descrption,
}

var tokenAdmin = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMTk5NzQsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbjEzIiwicm9sZSI6IkFETUlOIn0.Kr-6qbAUKDBikHdJMUEZ90GG0DvfM_xUo7gxG25nAOI"

type mockMotorReturnUsecase struct {
	mock.Mock
}

func (m *mockMotorReturnUsecase) AddMotorReturn(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error) {
	arg := m.Called(createMotorReturnRequest)
	return arg.Get(0).(motorReturnDto.CreateMotorReturnRequest), arg.Error(1)
}

func (m *mockMotorReturnUsecase) GetMotorReturnById(id string) (motorReturnDto.MotorReturnResponse, error) {
	arg := m.Called(id)
	return arg.Get(0).(motorReturnDto.MotorReturnResponse), arg.Error(1)
}

func (m *mockMotorReturnUsecase) GetMotorReturnAll() ([]motorReturnDto.MotorReturnResponse, error) {
	arg := m.Called()
	return arg.Get(0).([]motorReturnDto.MotorReturnResponse), arg.Error(1)
}

type MotorReturnDeliveryTestSuite struct {
	suite.Suite
	router  *gin.Engine
	usecase *mockMotorReturnUsecase
}

func (suite *MotorReturnDeliveryTestSuite) SetupTest() {
	suite.usecase = new(mockMotorReturnUsecase)
	suite.router = gin.Default()
	api := suite.router.Group("/api")
	v1 := api.Group("/v1")
	NewMotorReturnDelivey(v1, suite.usecase)
}

// create success
func (suite *MotorReturnDeliveryTestSuite) TestCreateMotorReturn_Success() {

	expectedResposnse := `{"responseCode":"2010101","responseMessage":"Motor return created","data":{"id":"907698c8-ae04-47b2-a7b9-68c46690c3f8","transaction_id":"621dfcb6-06df-4420-b98e-3ec04def9547","extra_charge":25000,"condition_motor":"Ban depan bocor","description":"bocor di jalan"}}`

	suite.usecase.On("AddMotorReturn", expectedCreateMotorReturn).Return(expectedCreateMotorReturn, nil)

	jsonData, _ := json.Marshal(expectedCreateMotorReturn)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/"+expectTransaction.EmployeeId+"/motor-return", bytes.NewBuffer(jsonData))
	accessTokenEmployee := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjE2ODksImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJkaW5vMTI0NTEiLCJyb2xlIjoiRU1QTE9ZRUUifQ.lGqvAnJ2q7E34ZnA2IXpYY5EfMNI1pjxDf4ZKhnh_oA"
	req.Header.Add("Authorization", accessTokenEmployee)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorReturnDeliveryTestSuite) TestCreateMotorReturn_FailedBind() {

	expectedResposnse := `{"responseCode":"4000101","responseMessage":"Bad Request","error_description":[{"field":"TransactionID","message":"field is required"},{"field":"ExtraCharge","message":"field is required"},{"field":"ConditionMotor","message":"field is required"},{"field":"Description","message":"field is required"}]}`

	suite.usecase.On("AddMotorReturn", expectedCreateMotorReturn).Return(expectedCreateMotorReturn, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/"+expectTransaction.EmployeeId+"/motor-return", nil)
	accessTokenEmployee := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjE2ODksImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJkaW5vMTI0NTEiLCJyb2xlIjoiRU1QTE9ZRUUifQ.lGqvAnJ2q7E34ZnA2IXpYY5EfMNI1pjxDf4ZKhnh_oA"
	req.Header.Add("Authorization", accessTokenEmployee)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorReturnDeliveryTestSuite) TestCreateMotorReturn_Failed() {

	expectedResposnse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"error"}`

	suite.usecase.On("AddMotorReturn", expectedCreateMotorReturn).Return(expectedCreateMotorReturn, errors.New("error"))

	jsonData, _ := json.Marshal(expectedCreateMotorReturn)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/"+expectTransaction.EmployeeId+"/motor-return", bytes.NewBuffer(jsonData))
	accessTokenEmployee := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjE2ODksImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJkaW5vMTI0NTEiLCJyb2xlIjoiRU1QTE9ZRUUifQ.lGqvAnJ2q7E34ZnA2IXpYY5EfMNI1pjxDf4ZKhnh_oA"
	req.Header.Add("Authorization", accessTokenEmployee)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	fmt.Println(w.Body.String())
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

// get by id success
func (suite *MotorReturnDeliveryTestSuite) TestGetMotorReturnById_Succes() {

	expectedResposnse := `{"responseCode":"2000201","responseMessage":"Success get motor return by id","data":{"id":"907698c8-ae04-47b2-a7b9-68c46690c3f8","return_date":"2024-03-07T00:00:00Z","extra_charge":25000,"condition_motor":"Ban depan bocor","description":"bocor di jalan","customer":{"id":"f4884dfc-7ef3-4d84-b77e-4fb930069da5","nama":"billkin","username":"billkin","alamat":"Bekasi","role":"USER","cant_rent":true,"created_at":"2024-03-07T23:39:42.63419Z","updated_at":"2024-03-07T23:39:42.63419Z","telepon":"08123456789"},"created_at":"2024-03-07T23:39:42.63419Z","updatad_at":"2024-03-07T23:39:42.63419Z"}}`

	suite.usecase.On("GetMotorReturnById", expectedMotorReturnResponse.ID).Return(expectedMotorReturnResponse, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/"+expectTransaction.EmployeeId+"/motor-return/"+expectedMotorReturnResponse.ID, nil)

	req.Header.Add("Authorization", tokenAdmin)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

// get by id fail
func (suite *MotorReturnDeliveryTestSuite) TestGetMotorReturnById_Fail() {

	expectedResposnse := `{"responseCode":"5000201","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	suite.usecase.On("GetMotorReturnById", expectedMotorReturnResponse.ID).Return(expectedMotorReturnResponse, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/"+expectTransaction.EmployeeId+"/motor-return/"+expectedMotorReturnResponse.ID, nil)

	req.Header.Add("Authorization", tokenAdmin)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

// get all success
func (suite *MotorReturnDeliveryTestSuite) TestGetAllMotorReturn_Succes() {

	expectedResposnse := `{"responseCode":"2000301","responseMessage":"Success get all motor return","data":[{"id":"907698c8-ae04-47b2-a7b9-68c46690c3f8","return_date":"2024-03-07T00:00:00Z","extra_charge":25000,"condition_motor":"Ban depan bocor","description":"bocor di jalan","customer":{"id":"f4884dfc-7ef3-4d84-b77e-4fb930069da5","nama":"billkin","username":"billkin","alamat":"Bekasi","role":"USER","cant_rent":true,"created_at":"2024-03-07T23:39:42.63419Z","updated_at":"2024-03-07T23:39:42.63419Z","telepon":"08123456789"},"created_at":"2024-03-07T23:39:42.63419Z","updatad_at":"2024-03-07T23:39:42.63419Z"}]}`

	suite.usecase.On("GetMotorReturnAll").Return(expectedAllMotorReturnResponse, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/motor-return", nil)

	req.Header.Add("Authorization", tokenAdmin)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

// get all fail
func (suite *MotorReturnDeliveryTestSuite) TestGetAllMotorReturn_Fail() {

	expectedResposnse := `{"responseCode":"5000301","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	suite.usecase.On("GetMotorReturnAll").Return(expectedAllMotorReturnResponse, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/motor-return", nil)

	req.Header.Add("Authorization", tokenAdmin)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func TestMotorReturnDelivery(t *testing.T) {
	suite.Run(t, new(MotorReturnDeliveryTestSuite))
}
