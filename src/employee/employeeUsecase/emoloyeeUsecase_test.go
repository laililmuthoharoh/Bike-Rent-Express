package employeeUsecase

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/src/employee"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var expectEmployee = employeeDto.Employee{
	ID:        "ec22df4c-1c1c-4012-9395-bc0994807e35",
	Name:      "dino",
	Telp:      "0812321412312",
	Username:  "dino123",
	Password:  "$2y$10$kGRStv6avxUMxjxRePgyh.i7fNbtJgV7WNb93DfS.Os2Lj2oU20Q2",
	CreatedAt: "2024-03-07T00:00:00Z",
	UpdatedAt: "2024-03-07T00:00:00Z",
}

type mockEmployeeRepository struct {
	mock.Mock
}

func (m *mockEmployeeRepository) Add(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error) {
	args := m.Called()
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
	args := m.Called()
	return args.Error(0)
}

func (m *mockEmployeeRepository) Delete(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

type EmployeeUCTestSuite struct {
	suite.Suite
	employeeUC             employee.EmployeeUsecase
	mockEmployeeRepository *mockEmployeeRepository
}

func (suite *EmployeeUCTestSuite) SetupTest() {
	suite.mockEmployeeRepository = new(mockEmployeeRepository)
	suite.employeeUC = NewEmployeeUsecase(suite.mockEmployeeRepository)
}

func (suite *EmployeeUCTestSuite) TestRegister_Success() {
	expectedCreatedEmployee := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Name,
		Password: expectEmployee.Password,
	}
	suite.mockEmployeeRepository.On("Add").Return(expectedCreatedEmployee, nil)
	actualCreatedEmployee, err := suite.employeeUC.Register(expectedCreatedEmployee)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedCreatedEmployee, actualCreatedEmployee)
}

func (suite *EmployeeUCTestSuite) TestRegister_Failed() {
	expectedCreatedEmployee := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Name,
		Password: expectEmployee.Password,
	}
	expectError := errors.New("error")
	suite.mockEmployeeRepository.On("Add").Return(employeeDto.CreateEmployeeRequest{}, expectError)
	actualCreatedEmployee, err := suite.employeeUC.Register(expectedCreatedEmployee)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectedCreatedEmployee, actualCreatedEmployee)
}

func (suite *EmployeeUCTestSuite) TestGet_Success() {
	expectAllEmployee := []employeeDto.Employee{
		expectEmployee,
	}

	suite.mockEmployeeRepository.On("Get").Return(expectAllEmployee, nil)
	actualGetEmployee, err := suite.employeeUC.Get()

	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectAllEmployee, actualGetEmployee)
}

func (suite *EmployeeUCTestSuite) TestGet_Failed() {
	expectAllEmployee := []employeeDto.Employee{
		expectEmployee,
	}

	expectError := errors.New("error")
	suite.mockEmployeeRepository.On("Get").Return([]employeeDto.Employee{}, expectError)

	actualGetEmployee, err := suite.employeeUC.Get()
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.NotEqual(suite.T(), expectAllEmployee, actualGetEmployee)
	assert.ErrorIs(suite.T(), expectError, err)
}

func (suite *EmployeeUCTestSuite) TestGetById_Success() {

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, nil)

	actualEmployee, err := suite.employeeUC.GetById(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUCTestSuite) TestGetById_FailedInvalidInput() {
	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(employeeDto.Employee{}, errors.New("invalid input syntax for type uuid"))

	actualEmployee, err := suite.employeeUC.GetById(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUCTestSuite) TestGetById_FailedNoRows() {
	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(employeeDto.Employee{}, sql.ErrNoRows)

	actualEmployee, err := suite.employeeUC.GetById(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), sql.ErrNoRows, err)
	assert.NotEqual(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUCTestSuite) TestUpdate_Success() {
	employeeUpdateRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
	}

	suite.mockEmployeeRepository.On("Update", employeeUpdateRequest).Return(expectEmployee, nil)

	actualEmployee, err := suite.employeeUC.Update(employeeUpdateRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUCTestSuite) TestUpdate_FailedInvalidInput() {
	employeeUpdateRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
	}

	suite.mockEmployeeRepository.On("Update", employeeUpdateRequest).Return(employeeDto.Employee{}, errors.New("invalid input syntax for type uuid"))

	actualEmployee, err := suite.employeeUC.Update(employeeUpdateRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUCTestSuite) TestUpdate_FailedNoRows() {
	employeeUpdateRequest := employeeDto.UpdateEmployeeRequest{
		ID:   expectEmployee.ID,
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
	}

	suite.mockEmployeeRepository.On("Update", employeeUpdateRequest).Return(employeeDto.Employee{}, sql.ErrNoRows)

	actualEmployee, err := suite.employeeUC.Update(employeeUpdateRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), sql.ErrNoRows, err)
	assert.NotEqual(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUCTestSuite) TestDelete_Success() {
	expectResult := "Sucessfully delete employee"

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, nil)
	suite.mockEmployeeRepository.On("Delete", expectEmployee.ID).Return(expectResult, nil)

	actualResult, err := suite.employeeUC.Delete(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectResult, actualResult)
}

func (suite *EmployeeUCTestSuite) TestDelete_FailedInvalidInputOrSqlNoRows() {
	expectResult := "Sucessfully delete employee"

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, sql.ErrNoRows)

	actualResult, err := suite.employeeUC.Delete(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectResult, actualResult)
}

func (suite *EmployeeUCTestSuite) TestDelete_FailedGetById() {
	expectResult := "Sucessfully delete employee"

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, errors.New("error"))

	actualResult, err := suite.employeeUC.Delete(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectResult, actualResult)
}

func (suite *EmployeeUCTestSuite) TestDelete_Failed() {
	expectResult := "Sucessfully delete employee"

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, nil)
	suite.mockEmployeeRepository.On("Delete", expectEmployee.ID).Return("", errors.New("error"))

	actualResult, err := suite.employeeUC.Delete(expectEmployee.ID)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectResult, actualResult)
}

func (suite *EmployeeUCTestSuite) TestLogin_Success() {
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: "daniel",
	}

	suite.mockEmployeeRepository.On("GetByUsername", loginRequest.Username).Return(expectEmployee, nil)

	_, err := suite.employeeUC.Login(loginRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
}

func (suite *EmployeeUCTestSuite) TestLogin_FailedInvalidInputOrNoRows() {
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: "daniel",
	}

	expectLoginResponse := employeeDto.LoginResponse{}

	suite.mockEmployeeRepository.On("GetByUsername", loginRequest.Username).Return(expectEmployee, errors.New("invalid input syntax for type uuid"))

	actualLoginResposne, err := suite.employeeUC.Login(loginRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), expectLoginResponse, actualLoginResposne)
}

func (suite *EmployeeUCTestSuite) TestLogin_FailedGetUsername() {
	loginRequest := employeeDto.LoginRequest{
		Username: expectEmployee.Username,
		Password: "daniel",
	}

	expectLoginResponse := employeeDto.LoginResponse{}

	suite.mockEmployeeRepository.On("GetByUsername", loginRequest.Username).Return(expectEmployee, errors.New("error"))

	actualLoginResposne, err := suite.employeeUC.Login(loginRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), expectLoginResponse, actualLoginResposne)
}

func (suite *EmployeeUCTestSuite) TestChangePassword_Success() {
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: "daniel",
		NewPassword: "daniel",
	}

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, nil)
	suite.mockEmployeeRepository.On("UpdatePassword").Return(nil)

	err := suite.employeeUC.ChangePassword(expectEmployee.ID, changePasswordRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
}

func (suite *EmployeeUCTestSuite) TestChangePassword_FailedIvalidInputOrNoRows() {
	changePasswordRequest := employeeDto.ChangePasswordRequest{
		PasswordOld: "daniel",
		NewPassword: "daniel",
	}

	suite.mockEmployeeRepository.On("GetById", expectEmployee.ID).Return(expectEmployee, sql.ErrNoRows)

	err := suite.employeeUC.ChangePassword(expectEmployee.ID, changePasswordRequest)
	suite.mockEmployeeRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "1", err.Error())
}

func TestEmployeeUCTestSuite(t *testing.T) {
	suite.Run(t, new(EmployeeUCTestSuite))
}
