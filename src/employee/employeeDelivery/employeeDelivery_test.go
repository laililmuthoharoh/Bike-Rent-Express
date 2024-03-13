package employeeDelivery

import (
	employeeDto "bike-rent-express/model/dto/employee"
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

type mockEmployeeUsecase struct {
	mock.Mock
}

var expectEmployee = employeeDto.Employee{
	ID:        "ec22df4c-1c1c-4012-9395-bc0994807e35",
	Name:      "dino",
	Telp:      "0812321412312",
	Username:  "dino123",
	Password:  "dino12345",
	CreatedAt: "2024-03-07T00:00:00Z",
	UpdatedAt: "2024-03-07T00:00:00Z",
}

var accessToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMTk5NzQsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbjEzIiwicm9sZSI6IkFETUlOIn0.Kr-6qbAUKDBikHdJMUEZ90GG0DvfM_xUo7gxG25nAOI"

func (m *mockEmployeeUsecase) Register(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error) {
	args := m.Called(employee)
	return args.Get(0).(employeeDto.CreateEmployeeRequest), args.Error(1)
}

func (m *mockEmployeeUsecase) Get() ([]employeeDto.Employee, error) {
	args := m.Called()
	return args.Get(0).([]employeeDto.Employee), args.Error(1)
}

func (m *mockEmployeeUsecase) GetById(id string) (employeeDto.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(employeeDto.Employee), args.Error(1)
}

func (m *mockEmployeeUsecase) Update(employeUpdateRequest employeeDto.UpdateEmployeeRequest) (employeeDto.Employee, error) {
	args := m.Called(employeUpdateRequest)
	return args.Get(0).(employeeDto.Employee), args.Error(1)
}

func (m *mockEmployeeUsecase) Delete(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *mockEmployeeUsecase) Login(loginRequest employeeDto.LoginRequest) (employeeDto.LoginResponse, error) {
	args := m.Called(loginRequest)
	return args.Get(0).(employeeDto.LoginResponse), args.Error(1)
}

func (m *mockEmployeeUsecase) ChangePassword(id string, changePasswordRequest employeeDto.ChangePasswordRequest) error {
	args := m.Called(id, changePasswordRequest)
	return args.Error(0)
}

type EmployeeDeliverySuite struct {
	suite.Suite
	mockEmployeeUC *mockEmployeeUsecase
	router         *gin.Engine
}

func (suite *EmployeeDeliverySuite) SetupTest() {
	suite.mockEmployeeUC = new(mockEmployeeUsecase)
	suite.router = gin.Default()
	api := suite.router.Group("/api")
	v1 := api.Group("/v1")
	NewEmployeeDelivery(v1, suite.mockEmployeeUC)
}

func (suite *EmployeeDeliverySuite) TestAddEmployee_Success() {
	createEmployeRequest := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}
	expectResponse := `{"responseCode":"2010101","responseMessage":"Employee Created","data":{"id":"ec22df4c-1c1c-4012-9395-bc0994807e35","name":"dino","telp":"0812321412312","username":"dino123","password":"dino12345"}}`
	suite.mockEmployeeUC.On("Register", createEmployeRequest).Return(createEmployeRequest, nil)

	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(createEmployeRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/register", bytes.NewBuffer(jsonData))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestAddEmployee_FailedBind() {
	createEmployeRequest := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
	}
	expectResponse := `{"responseCode":"4000101","responseMessage":"Bad Request","error_description":[{"field":"Password","message":"field is required"}]}`
	suite.mockEmployeeUC.On("Register", createEmployeRequest).Return(createEmployeRequest, nil)

	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(createEmployeRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/register", bytes.NewBuffer(jsonData))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestAddEmployee_FailedUsernameReady() {
	createEmployeRequest := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}
	expectResponse := `{"responseCode":"2000101","responseMessage":"username already in use"}`
	suite.mockEmployeeUC.On("Register", createEmployeRequest).Return(createEmployeRequest, errors.New("1"))

	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(createEmployeRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/register", bytes.NewBuffer(jsonData))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestAddEmployee_FailedUsername() {
	createEmployeRequest := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}
	expectResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"error"}`
	suite.mockEmployeeUC.On("Register", createEmployeRequest).Return(createEmployeRequest, errors.New("error"))

	w := httptest.NewRecorder()
	jsonData, _ := json.Marshal(createEmployeRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/register", bytes.NewBuffer(jsonData))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestGetEmployeeById_Success() {
	suite.mockEmployeeUC.On("GetById", expectEmployee.ID).Return(expectEmployee, nil)

	w := httptest.NewRecorder()
	expectResponse := `{"responseCode":"2000202","data":{"id":"ec22df4c-1c1c-4012-9395-bc0994807e35","name":"dino","telp":"0812321412312","username":"dino123","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z"}}`

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/"+expectEmployee.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestGetEmployeeById_FailedGetByIdDataNotFound() {
	suite.mockEmployeeUC.On("GetById", expectEmployee.ID).Return(expectEmployee, errors.New("1"))

	w := httptest.NewRecorder()
	expectResponse := `{"responseCode":"2000201","responseMessage":"Data not found"}`

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/"+expectEmployee.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestGetEmployeeById_FailedGetById() {
	suite.mockEmployeeUC.On("GetById", expectEmployee.ID).Return(expectEmployee, errors.New("error"))

	w := httptest.NewRecorder()
	expectResponse := `{"responseCode":"5000201","responseMessage":"internal server error","error":"error"}`

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee/"+expectEmployee.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestGetAllEmployee_Success() {
	expectResponse := `{"responseCode":"2000302","responseMessage":"Success Get All Employee","data":[{"id":"ec22df4c-1c1c-4012-9395-bc0994807e35","name":"dino","telp":"0812321412312","username":"dino123","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z"}]}`
	suite.mockEmployeeUC.On("Get").Return([]employeeDto.Employee{expectEmployee}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}
func (suite *EmployeeDeliverySuite) TestGetAllEmployee_FailedDataEmpty() {
	expectResponse := `{"responseCode":"2000301","responseMessage":"Data empty"}`
	suite.mockEmployeeUC.On("Get").Return([]employeeDto.Employee{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestGetAllEmployee_Failed() {
	expectResponse := `{"responseCode":"5000301","responseMessage":"internal server error","error":"error"}`
	suite.mockEmployeeUC.On("Get").Return([]employeeDto.Employee{expectEmployee}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/employee", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestUpdateEmployee_Success() {
	expectResponse := `{"responseCode":"2000401","responseMessage":"Success update employee","data":{"id":"ec22df4c-1c1c-4012-9395-bc0994807e35","name":"dino","telp":"0812321412312","username":"dino123","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z"}}`
	updateEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
	}

	suite.mockEmployeeUC.On("Update", updateEmployeeRequest).Return(expectEmployee, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(updateEmployeeRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestUpdateEmployee_FailedBind() {
	expectResponse := `{"responseCode":"4000401","responseMessage":"Bad Request","error_description":[{"field":"Telp","message":"field is required"}]}`
	updateEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
	}

	suite.mockEmployeeUC.On("Update", updateEmployeeRequest).Return(expectEmployee, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(updateEmployeeRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestUpdateEmployee_FailedDataNotFound() {
	expectResponse := `{"responseCode":"2000401","responseMessage":"Data not found"}`
	updateEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
	}

	suite.mockEmployeeUC.On("Update", updateEmployeeRequest).Return(expectEmployee, errors.New("1"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(updateEmployeeRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestUpdateEmployee_Failed() {
	expectResponse := `{"responseCode":"5000401","responseMessage":"internal server error","error":"error"}`
	updateEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
	}

	suite.mockEmployeeUC.On("Update", updateEmployeeRequest).Return(expectEmployee, errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(updateEmployeeRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestDeleteEmployee_Success() {
	expectResponse := `{"responseCode":"2000502","responseMessage":"Sucessfully delete employee"}`
	suite.mockEmployeeUC.On("Delete", expectEmployee.ID).Return("Sucessfully delete employee", nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/employee/"+expectEmployee.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestDeleteEmployee_FailedDataNotFound() {
	expectResponse := `{"responseCode":"2000501","responseMessage":"Data not found"}`
	suite.mockEmployeeUC.On("Delete", expectEmployee.ID).Return("Sucessfully delete employee", errors.New("1"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/employee/"+expectEmployee.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestDeleteEmployee_Failed() {
	expectResponse := `{"responseCode":"5000501","responseMessage":"internal server error","error":"errror"}`
	suite.mockEmployeeUC.On("Delete", expectEmployee.ID).Return("Sucessfully delete employee", errors.New("errror"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/employee/"+expectEmployee.ID, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestLoginEmployee_Success() {
	expectResponse := `{"responseCode":"2000603","responseMessage":"Login successfully","data":{"access_token":"","employee":{"id":"ec22df4c-1c1c-4012-9395-bc0994807e35","name":"dino","telp":"0812321412312","username":"dino123","created_at":"2024-03-07T00:00:00Z","updated_at":"2024-03-07T00:00:00Z"}}}`
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	loginResponse := employeeDto.LoginResponse{
		AccessToken: "",
		Employee:    expectEmployee,
	}

	suite.mockEmployeeUC.On("Login", loginRequest).Return(loginResponse, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestLoginEmployee_FailedBind() {
	expectResponse := `{"responseCode":"4000601","responseMessage":"Bad Request","error_description":[{"field":"Password","message":"field is required"}]}`
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
	}

	loginResponse := employeeDto.LoginResponse{
		AccessToken: "",
		Employee:    expectEmployee,
	}

	suite.mockEmployeeUC.On("Login", loginRequest).Return(loginResponse, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestLoginEmployee_FailedWrongPassOrUsername1() {
	expectResponse := `{"responseCode":"2000601","responseMessage":"Incorrect username or password"}`
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	loginResponse := employeeDto.LoginResponse{
		AccessToken: "",
		Employee:    expectEmployee,
	}

	suite.mockEmployeeUC.On("Login", loginRequest).Return(loginResponse, errors.New("1"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestLoginEmployee_FailedWrongPassOrUsername2() {
	expectResponse := `{"responseCode":"2000602","responseMessage":"Incorrect username or password"}`
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	loginResponse := employeeDto.LoginResponse{
		AccessToken: "",
		Employee:    expectEmployee,
	}

	suite.mockEmployeeUC.On("Login", loginRequest).Return(loginResponse, errors.New("2"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestLoginEmployee_Failed() {
	expectResponse := `{"responseCode":"5000601","responseMessage":"internal server error","error":"error"}`
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	loginResponse := employeeDto.LoginResponse{
		AccessToken: "",
		Employee:    expectEmployee,
	}

	suite.mockEmployeeUC.On("Login", loginRequest).Return(loginResponse, errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/employee/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestChangePassword_Success() {
	expectResponse := `{"responseCode":"2000703","responseMessage":"Success updated password"}`
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: expectEmployee.Password,
		NewPassword: "test",
	}

	suite.mockEmployeeUC.On("ChangePassword", expectEmployee.ID, changePasswordRequest).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePasswordRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestChangePassword_FailedBind() {
	expectResponse := `{"responseCode":"4000701","responseMessage":"Bad Request","error_description":[{"field":"NewPassword","message":"field is required"}]}`
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: expectEmployee.Password,
	}

	suite.mockEmployeeUC.On("ChangePassword", expectEmployee.ID, changePasswordRequest).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePasswordRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestChangePassword_FailedDataNotFound1() {
	expectResponse := `{"responseCode":"2000701","responseMessage":"Data not found"}`
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: expectEmployee.Password,
		NewPassword: "test",
	}

	suite.mockEmployeeUC.On("ChangePassword", expectEmployee.ID, changePasswordRequest).Return(errors.New("1"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePasswordRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestChangePassword_FailedDataNotFound2() {
	expectResponse := `{"responseCode":"2000702","responseMessage":"password does not match"}`
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: expectEmployee.Password,
		NewPassword: "test",
	}

	suite.mockEmployeeUC.On("ChangePassword", expectEmployee.ID, changePasswordRequest).Return(errors.New("2"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePasswordRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *EmployeeDeliverySuite) TestChangePassword_Failed() {
	expectResponse := `{"responseCode":"5000701","responseMessage":"internal server error","error":"error"}`
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: expectEmployee.Password,
		NewPassword: "test",
	}

	suite.mockEmployeeUC.On("ChangePassword", expectEmployee.ID, changePasswordRequest).Return(errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePasswordRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/employee/"+expectEmployee.ID+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func TestEmployeeDelivery(t *testing.T) {
	suite.Run(t, new(EmployeeDeliverySuite))
}
