package usersDelivery

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto"
	"bytes"
	"database/sql"
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

var expectUsers = dto.GetUsers{
	Uuid:       "omosiof32131",
	Name:       "test",
	Username:   "test",
	Address:    "test",
	Role:       "USER",
	Can_rent:   true,
	Password:   "$2y$10$VU8yVSpQeECxjpB40IfLY.8FTtWWRnxySvIEKOJpUUHkd32Strtdq",
	Created_at: "0000",
	Updated_at: "0000",
	Telp:       "0813123",
}

var accessToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjM0NjAsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJhZG1pbjEyMjMiLCJyb2xlIjoiQURNSU4ifQ.nC0gb9Vey1JXz2zQZgvg96NJIPm6CFSXHcA5OEWTD2I"

type mockUserUC struct {
	mock.Mock
}

func (m *mockUserUC) RegisterUsers(newUsers dto.RegisterUsers) error {
	args := m.Called(newUsers)
	return args.Error(0)
}

func (m *mockUserUC) GetByID(id string) (dto.GetUsers, error) {
	args := m.Called(id)
	return args.Get(0).(dto.GetUsers), args.Error(1)
}

func (m *mockUserUC) GetAllUsers() ([]dto.GetUsers, error) {
	args := m.Called()
	return args.Get(0).([]dto.GetUsers), args.Error(1)
}

func (m *mockUserUC) UpdateUsers(usersItem dto.Users) error {
	args := m.Called(usersItem)
	return args.Error(0)
}

func (m *mockUserUC) LoginUsers(loginRequest model.LoginRequest) (dto.LoginResponse, error) {
	args := m.Called(loginRequest)
	return args.Get(0).(dto.LoginResponse), args.Error(1)
}

func (m *mockUserUC) TopUp(topUpRequest dto.TopUpRequest) error {
	args := m.Called(topUpRequest)
	return args.Error(0)
}

func (m *mockUserUC) ChangePassword(changePasswordRequest dto.ChangePassword) error {
	args := m.Called(changePasswordRequest)
	return args.Error(0)
}

type UsersDeliveryTestSuite struct {
	suite.Suite
	mockUserUC *mockUserUC
	router     *gin.Engine
}

func (suite *UsersDeliveryTestSuite) SetupTest() {
	suite.mockUserUC = new(mockUserUC)
	suite.router = gin.Default()
	api := suite.router.Group("/api")
	v1 := api.Group("/v1")
	NewUsersDelivery(v1, suite.mockUserUC)
}

func (suite *UsersDeliveryTestSuite) TestGetAllUsers_Success() {
	expectResponse := `{"responseCode":"2000102","responseMessage":"Success","data":[{"id":"omosiof32131","nama":"test","username":"test","alamat":"test","role":"USER","cant_rent":true,"created_at":"0000","updated_at":"0000","telepon":"0813123"}]}`
	allUser := []dto.GetUsers{
		expectUsers,
	}
	suite.mockUserUC.On("GetAllUsers").Return(allUser, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestGetAllUsers_FailedDataEmpty() {
	expectResponse := `{"responseCode":"2000101","responseMessage":"Data empty"}`
	suite.mockUserUC.On("GetAllUsers").Return([]dto.GetUsers{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestGetAllUsers_Failed() {
	expectResponse := `{"responseCode":"5000101","responseMessage":"internal server error","error":"error"}`
	allUser := []dto.GetUsers{
		expectUsers,
	}
	suite.mockUserUC.On("GetAllUsers").Return(allUser, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestUpdateUsers_Success() {
	expectResponse := `{"responseCode":"2000202","responseMessage":"Account Updated"}`
	user := dto.Users{
		ID:         expectUsers.Uuid,
		Name:       expectUsers.Name,
		Username:   expectUsers.Username,
		Password:   expectUsers.Password,
		Address:    expectUsers.Address,
		Role:       expectUsers.Role,
		CanRent:    expectUsers.Can_rent,
		Updated_at: expectUsers.Updated_at,
		Telp:       expectUsers.Telp,
	}

	suite.mockUserUC.On("UpdateUsers", user).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+user.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestUpdateUsers_FailedBind() {
	expectResponse := `{"responseCode":"4000201","responseMessage":"Bad Request","error_description":[{"field":"Telp","message":"field is required"}]}`
	user := dto.Users{
		ID:         expectUsers.Uuid,
		Name:       expectUsers.Name,
		Username:   expectUsers.Username,
		Password:   expectUsers.Password,
		Address:    expectUsers.Address,
		Role:       expectUsers.Role,
		CanRent:    expectUsers.Can_rent,
		Updated_at: expectUsers.Updated_at,
	}

	suite.mockUserUC.On("UpdateUsers", user).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+user.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestUpdateUsers_FailedInputOrNoRows() {
	expectResponse := `{"responseCode":"2000201","responseMessage":"User not found"}`
	user := dto.Users{
		ID:         expectUsers.Uuid,
		Name:       expectUsers.Name,
		Username:   expectUsers.Username,
		Password:   expectUsers.Password,
		Address:    expectUsers.Address,
		Role:       expectUsers.Role,
		CanRent:    expectUsers.Can_rent,
		Updated_at: expectUsers.Updated_at,
		Telp:       expectUsers.Telp,
	}

	suite.mockUserUC.On("UpdateUsers", user).Return(sql.ErrNoRows)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+user.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestUpdateUsers_Failed() {
	expectResponse := `{"responseCode":"5000201","responseMessage":"internal server error","error":"error"}`
	user := dto.Users{
		ID:         expectUsers.Uuid,
		Name:       expectUsers.Name,
		Username:   expectUsers.Username,
		Password:   expectUsers.Password,
		Address:    expectUsers.Address,
		Role:       expectUsers.Role,
		CanRent:    expectUsers.Can_rent,
		Updated_at: expectUsers.Updated_at,
		Telp:       expectUsers.Telp,
	}

	suite.mockUserUC.On("UpdateUsers", user).Return(errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+user.ID, bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestGetById_Success() {
	expectResponse := `{"responseCode":"2000301","responseMessage":"Success","data":{"id":"omosiof32131","nama":"test","username":"test","alamat":"test","role":"USER","cant_rent":true,"created_at":"0000","updated_at":"0000","telepon":"0813123"}}`
	suite.mockUserUC.On("GetByID", expectUsers.Uuid).Return(expectUsers, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/"+expectUsers.Uuid, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestGetById_FailedDataNotFound() {
	expectResponse := `{"responseCode":"2000301","responseMessage":"User not found"}`
	suite.mockUserUC.On("GetByID", expectUsers.Uuid).Return(expectUsers, errors.New("1"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/"+expectUsers.Uuid, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestGetById_Failed() {
	expectResponse := `{"responseCode":"5000301","responseMessage":"internal server error","error":"error"}`
	suite.mockUserUC.On("GetByID", expectUsers.Uuid).Return(expectUsers, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/"+expectUsers.Uuid, nil)
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestRegisterUser_Success() {
	expectResponse := `{"responseCode":"2010402","responseMessage":"Account Created","data":{"name":"test","username":"test","password":"$2y$10$VU8yVSpQeECxjpB40IfLY.8FTtWWRnxySvIEKOJpUUHkd32Strtdq","address":"test","role":"USER","telp":"0813123"}}`
	newUser := dto.RegisterUsers{
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Telp,
	}

	suite.mockUserUC.On("RegisterUsers", newUser).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(newUser)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 201, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestRegisterUser_FailedBind() {
	expectResponse := `{"responseCode":"4000401","responseMessage":"Bad Request","error_description":[{"field":"Telp","message":"field is required"}]}`
	newUser := dto.RegisterUsers{
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
	}

	suite.mockUserUC.On("RegisterUsers", newUser).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(newUser)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestRegisterUser_FailedUsername() {
	expectResponse := `{"responseCode":"2000401","responseMessage":"username already in use"}`
	newUser := dto.RegisterUsers{
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Telp,
	}

	suite.mockUserUC.On("RegisterUsers", newUser).Return(errors.New("1"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(newUser)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestRegisterUser_Failed() {
	expectResponse := `{"responseCode":"5000402","responseMessage":"internal server error","error":"error"}`
	newUser := dto.RegisterUsers{
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Telp,
	}

	suite.mockUserUC.On("RegisterUsers", newUser).Return(errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(newUser)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestLoginUser_Success() {
	expectResponse := `{"responseCode":"2000502","responseMessage":"login success","data":{"acces_token":"1","user":{"id":"omosiof32131","name":"test","username":"test","password":"$2y$10$VU8yVSpQeECxjpB40IfLY.8FTtWWRnxySvIEKOJpUUHkd32Strtdq","address":"test","role":"USER","can_rent":true,"Updated_at":"0000","telp":"0813123"}}}`
	loginRequest := model.LoginRequest{
		Username: expectUsers.Name,
		Password: "test",
	}

	loginResponse := dto.LoginResponse{
		AccesToken: "1",
		User: dto.Users{
			ID:         expectUsers.Uuid,
			Name:       expectUsers.Name,
			Username:   expectUsers.Username,
			Password:   expectUsers.Password,
			Address:    expectUsers.Address,
			Role:       expectUsers.Role,
			CanRent:    expectUsers.Can_rent,
			Updated_at: expectUsers.Updated_at,
			Telp:       expectUsers.Telp,
		},
	}

	suite.mockUserUC.On("LoginUsers", loginRequest).Return(loginResponse, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestLoginUser_Bind() {
	expectResponse := `{"responseCode":"4000501","responseMessage":"bad request","error_description":[{"field":"Password","message":"field is required"}]}`
	loginRequest := model.LoginRequest{
		Username: expectUsers.Name,
	}

	loginResponse := dto.LoginResponse{
		AccesToken: "1",
		User: dto.Users{
			ID:         expectUsers.Uuid,
			Name:       expectUsers.Name,
			Username:   expectUsers.Username,
			Password:   expectUsers.Password,
			Address:    expectUsers.Address,
			Role:       expectUsers.Role,
			CanRent:    expectUsers.Can_rent,
			Updated_at: expectUsers.Updated_at,
			Telp:       expectUsers.Telp,
		},
	}

	suite.mockUserUC.On("LoginUsers", loginRequest).Return(loginResponse, nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestLoginUser_FailedWrongUsername() {
	expectResponse := `{"responseCode":"2000501","responseMessage":"Incorrect username or password"}`
	loginRequest := model.LoginRequest{
		Username: expectUsers.Name,
		Password: "test",
	}

	loginResponse := dto.LoginResponse{
		AccesToken: "1",
		User: dto.Users{
			ID:         expectUsers.Uuid,
			Name:       expectUsers.Name,
			Username:   expectUsers.Username,
			Password:   expectUsers.Password,
			Address:    expectUsers.Address,
			Role:       expectUsers.Role,
			CanRent:    expectUsers.Can_rent,
			Updated_at: expectUsers.Updated_at,
			Telp:       expectUsers.Telp,
		},
	}

	suite.mockUserUC.On("LoginUsers", loginRequest).Return(loginResponse, errors.New("1"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestLoginUser_Failed() {
	expectResponse := `{"responseCode":"5000501","responseMessage":"internal server error","error":"error"}`
	loginRequest := model.LoginRequest{
		Username: expectUsers.Name,
		Password: "test",
	}

	loginResponse := dto.LoginResponse{
		AccesToken: "1",
		User: dto.Users{
			ID:         expectUsers.Uuid,
			Name:       expectUsers.Name,
			Username:   expectUsers.Username,
			Password:   expectUsers.Password,
			Address:    expectUsers.Address,
			Role:       expectUsers.Role,
			CanRent:    expectUsers.Can_rent,
			Updated_at: expectUsers.Updated_at,
			Telp:       expectUsers.Telp,
		},
	}

	suite.mockUserUC.On("LoginUsers", loginRequest).Return(loginResponse, errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(json))

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestTopUp_Success() {
	expectResponse := `{"responseCode":"2000601","responseMessage":"Success Top Up"}`
	topUpRequest := dto.TopUpRequest{
		Amount: 1,
		UserID: expectUsers.Uuid,
	}

	suite.mockUserUC.On("TopUp", topUpRequest).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(topUpRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/top-up", bytes.NewBuffer(json))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjg4NTQsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJ1c2VyMjEiLCJpZCI6IiIsInJvbGUiOiJVU0VSIn0.xT64rPCg2Ud9m_BSkXSOscYsQ3fT6-x6yNyttpfUNhI")

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestTopUp_FailedBind() {
	expectResponse := `{"responseCode":"4000601","responseMessage":"Bad Request","error_description":[{"field":"Amount","message":"field is required"}]}`
	topUpRequest := dto.TopUpRequest{
		UserID: expectUsers.Uuid,
	}

	suite.mockUserUC.On("TopUp", topUpRequest).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(topUpRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/top-up", bytes.NewBuffer(json))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjg4NTQsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJ1c2VyMjEiLCJpZCI6IiIsInJvbGUiOiJVU0VSIn0.xT64rPCg2Ud9m_BSkXSOscYsQ3fT6-x6yNyttpfUNhI")

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestTopUp_Failed() {
	expectResponse := `{"responseCode":"5000601","responseMessage":"internal server error","error":"error"}`
	topUpRequest := dto.TopUpRequest{
		Amount: 1,
		UserID: expectUsers.Uuid,
	}

	suite.mockUserUC.On("TopUp", topUpRequest).Return(errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(topUpRequest)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/top-up", bytes.NewBuffer(json))
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjg4NTQsImlzcyI6ImluY3ViYXRpb24tZ29sYW5nIiwidXNlcm5hbWUiOiJ1c2VyMjEiLCJpZCI6IiIsInJvbGUiOiJVU0VSIn0.xT64rPCg2Ud9m_BSkXSOscYsQ3fT6-x6yNyttpfUNhI")

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestChangePassword_Success() {
	expectResponse := `{"responseCode":"2000703","responseMessage":"Success change password"}`
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserUC.On("ChangePassword", changePassword).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePassword)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestChangePassword_FailedBind() {
	expectResponse := `{"responseCode":"4000701","responseMessage":"Bad Request","error_description":[{"field":"OldPassword","message":"field is required"}]}`
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		NewPassword: "test",
	}

	suite.mockUserUC.On("ChangePassword", changePassword).Return(nil)

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePassword)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 400, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestChangePassword_FailedDataNotFound() {
	expectResponse := `{"responseCode":"2000701","responseMessage":"Data not found"}`
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserUC.On("ChangePassword", changePassword).Return(errors.New("1"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePassword)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestChangePassword_FailedPasswordNotMatch() {
	expectResponse := `{"responseCode":"2000702","responseMessage":"password does not match"}`
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserUC.On("ChangePassword", changePassword).Return(errors.New("2"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePassword)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func (suite *UsersDeliveryTestSuite) TestChangePassword_Failed() {
	expectResponse := `{"responseCode":"5000702","responseMessage":"internal server error","error":"error"}`
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserUC.On("ChangePassword", changePassword).Return(errors.New("error"))

	w := httptest.NewRecorder()
	json, _ := json.Marshal(changePassword)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/"+expectUsers.Uuid+"/change-password", bytes.NewBuffer(json))
	req.Header.Add("Authorization", accessToken)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), 500, w.Code)
	assert.Equal(suite.T(), expectResponse, w.Body.String())
}

func TestUsersDelivery(t *testing.T) {
	suite.Run(t, new(UsersDeliveryTestSuite))
}
