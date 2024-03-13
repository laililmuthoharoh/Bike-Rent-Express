package motorVehicleDelivery

import (
	"bike-rent-express/model/dto/motorVehicleDto"
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

var expectedUpdateMotorVehicle = motorVehicleDto.UpdateMotorVehicle{
	Name:           "Vario",
	Type:           "MATIC",
	Price:          50000,
	Plat:           "BA1234I",
	ProductionYear: "2023",
	Status:         "AVAILABLE",
}

var expectedCreateMotorVehicle = motorVehicleDto.CreateMotorVehicle{
	Name:           "Vario",
	Type:           "MATIC",
	Price:          50000,
	Plat:           "BA1234I",
	ProductionYear: "2023",
	Status:         "AVAILABLE",
}

var token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA0MTI0MTgsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbiIsImlkIjoiIiwicm9sZSI6IkFETUlOIn0.VL7b2UjsczLaEyLNJFYhpbUCSlbfG4tAMGpEPXi9JOk"

type mockMotorVehicleUsecase struct {
	mock.Mock
}

func (m *mockMotorVehicleUsecase) GetAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error) {
	arg := m.Called()
	return arg.Get(0).([]motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleUsecase) GetMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error) {
	arg := m.Called(id)
	return arg.Get(0).(motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleUsecase) CreateMotorVehicle(motor motorVehicleDto.CreateMotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	arg := m.Called(expectedCreateMotorVehicle)
	return arg.Get(0).(motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleUsecase) UpdateMotorVehicle(id string, input motorVehicleDto.UpdateMotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	arg := m.Called(id, expectedUpdateMotorVehicle)
	return arg.Get(0).(motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleUsecase) DeleteMotorVehicle(id string) error {
	arg := m.Called(id)
	return arg.Error(0)
}

type MotorVehicleDeliveryTestSuite struct {
	suite.Suite
	router  *gin.Engine
	usecase *mockMotorVehicleUsecase
}

func (suite *MotorVehicleDeliveryTestSuite) SetupTest() {
	suite.router = gin.Default()
	suite.usecase = new(mockMotorVehicleUsecase)

	api := suite.router.Group("/api")
	v1 := api.Group("/v1")
	NewMotorVehicleDelivery(v1, suite.usecase)
}

func (suite *MotorVehicleDeliveryTestSuite) TestGetAllMotorVehicle_Success() {
	expected := []motorVehicleDto.MotorVehicle{expectedMotorVehicleById}
	expectedResposnse := `{"responseCode":"2000102","responseMessage":"success","data":[{"id":"3a57713c-24d0-41f8-bfa8-f8f721dba9e4","name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z","production_year":"2023","status":"AVAILABLE"}]}`

	suite.usecase.On("GetAllMotorVehicle").Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/motor-vehicles/", nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestGetAllMotorVehicle_FailedDataEmpty() {
	expectedResposnse := `{"responseCode":"2000101","responseMessage":"Empty data"}`

	suite.usecase.On("GetAllMotorVehicle").Return([]motorVehicleDto.MotorVehicle{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/motor-vehicles/", nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestGetAllMotorVehicle_Fail() {
	expected := []motorVehicleDto.MotorVehicle{expectedMotorVehicleById}
	expectedResposnse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	suite.usecase.On("GetAllMotorVehicle").Return(expected, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/motor-vehicles/", nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestGetMotorVehicleById_Success() {

	expectedResposnse := `{"responseCode":"2000202","responseMessage":"success get data by id","data":{"id":"3a57713c-24d0-41f8-bfa8-f8f721dba9e4","name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z","production_year":"2023","status":"AVAILABLE"}}`

	suite.usecase.On("GetMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestGetMotorVehicleById_Fail() {

	expectedResposnse := `{"responseCode":"5000201","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	suite.usecase.On("GetMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestGetMotorVehicleById_FailDataNotFound() {

	expectedResposnse := `{"responseCode":"2000201","responseMessage":"Data not found"}`
	expectedError := errors.New("1")

	suite.usecase.On("GetMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestCreateMotorVehicle_Success() {

	expectedResposnse := `{"responseCode":"2010301","responseMessage":"motor vehicle created","data":{"id":"3a57713c-24d0-41f8-bfa8-f8f721dba9e4","name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z","production_year":"2023","status":"AVAILABLE"}}`

	suite.usecase.On("CreateMotorVehicle", expectedCreateMotorVehicle).Return(expectedMotorVehicleById, nil)

	requestbody := []byte(`{"name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","production_year":"2023","status":"AVAILABLE"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/motor-vehicles/", bytes.NewBuffer(requestbody))

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestCreateMotorVehicle_FailBadRequest() {

	expectedCreateMotorVehicle = motorVehicleDto.CreateMotorVehicle{
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	}

	expectedResposnse := `{"responseCode":"4000301","responseMessage":"Bad Request","error_description":[{"field":"Name","message":"field is required"}]}`

	suite.usecase.On("CreateMotorVehicle", expectedCreateMotorVehicle).Return(expectedMotorVehicleById, nil)

	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(expectedCreateMotorVehicle)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/motor-vehicles/", bytes.NewBuffer(jsonData))

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestCreateMotorVehicle_fail() {

	expectedResposnse := `{"responseCode":"5000301","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	requestbody := []byte(`{"name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","production_year":"2023","status":"AVAILABLE"}`)

	suite.usecase.On("CreateMotorVehicle", expectedCreateMotorVehicle).Return(expectedMotorVehicleById, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/motor-vehicles/", bytes.NewBuffer(requestbody))

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestUpdateMotorVehicle_Success() {

	expectedResposnse := `{"responseCode":"2000401","responseMessage":"motor vehicle updated","data":{"id":"3a57713c-24d0-41f8-bfa8-f8f721dba9e4","name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z","production_year":"2023","status":"AVAILABLE"}}`

	requestbody := []byte(`{"name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","production_year":"2023","status":"AVAILABLE"}`)

	suite.usecase.On("UpdateMotorVehicle", mock.Anything, mock.Anything).Return(expectedMotorVehicleById, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, bytes.NewBuffer(requestbody))

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestUpdateMotorVehicle_FailBadRequest() {

	expectedUpdateMotorVehicle = motorVehicleDto.UpdateMotorVehicle{
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	}

	expectedResposnse := `{"responseCode":"4000401","responseMessage":"Bad Request","error_description":[{"field":"Name","message":"field is required"}]}`

	suite.usecase.On("UpdateMotorVehicle", mock.Anything, mock.Anything).Return(expectedMotorVehicleById, nil)

	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(expectedUpdateMotorVehicle)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, bytes.NewBuffer(jsonData))

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestUpdateMotorVehicle_fail() {

	expectedResposnse := `{"responseCode":"5000401","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	requestbody := []byte(`{"name":"Vario","type":"MATIC","price":50000,"plat":"BA1234I","production_year":"2023","status":"AVAILABLE"}`)

	suite.usecase.On("UpdateMotorVehicle", expectedMotorVehicleById.Id, expectedUpdateMotorVehicle).Return(expectedMotorVehicleById, expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, bytes.NewBuffer(requestbody))

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestDeleteMotorVehicle_Success() {

	expectedResposnse := `{"responseCode":"2000501","responseMessage":"Sucessfully deleted motor vehicle"}`

	suite.usecase.On("DeleteMotorVehicle", expectedMotorVehicleById.Id).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func (suite *MotorVehicleDeliveryTestSuite) TestDeleteMotorVehicle_Fail() {

	expectedResposnse := `{"responseCode":"5000501","responseMessage":"internal server error","error":"error"}`
	expectedError := errors.New("error")

	suite.usecase.On("DeleteMotorVehicle", expectedMotorVehicleById.Id).Return(expectedError)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/motor-vehicles/"+expectedMotorVehicleById.Id, nil)

	req.Header.Add("Authorization", token)
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectedResposnse, w.Body.String())
}

func TestMotorVehicleDelivery(t *testing.T) {
	suite.Run(t, new(MotorVehicleDeliveryTestSuite))
}
