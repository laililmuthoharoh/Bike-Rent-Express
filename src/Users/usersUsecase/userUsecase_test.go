package usersUsecase

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto"
	"bike-rent-express/src/Users"
	"database/sql"
	"errors"
	"testing"

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

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) RegisterUsers(newUsers dto.RegisterUsers) error {
	args := m.Called()
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
	args := m.Called()
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

type UserUCTestSuite struct {
	suite.Suite
	userUC             Users.UsersUsecase
	mockUserRepository *mockUserRepository
}

func (suite *UserUCTestSuite) SetupTest() {
	suite.mockUserRepository = new(mockUserRepository)
	suite.userUC = NewUsersUsecase(suite.mockUserRepository)
}

func (suite *UserUCTestSuite) TestGetAllUser_Success() {
	expectUserGetAll := []dto.GetUsers{
		expectUsers,
	}

	suite.mockUserRepository.On("GetAll").Return(expectUserGetAll, nil)

	acutalUserGetAll, err := suite.userUC.GetAllUsers()
	suite.mockUserRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectUserGetAll, acutalUserGetAll)
}

func (suite *UserUCTestSuite) TestGetAllUser_Failed() {
	expectUserGetAll := []dto.GetUsers{
		expectUsers,
	}
	expectError := errors.New("error")
	suite.mockUserRepository.On("GetAll").Return(expectUserGetAll, expectError)

	acutalUserGetAll, err := suite.userUC.GetAllUsers()
	suite.mockUserRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.ErrorIs(suite.T(), expectError, err)
	assert.NotEqual(suite.T(), expectUserGetAll, acutalUserGetAll)
}

func (suite *UserUCTestSuite) TestUpdateUser_Success() {
	expectUpdateUser := dto.Users{
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

	suite.mockUserRepository.On("UpdateUsers", expectUpdateUser).Return(nil)
	err := suite.userUC.UpdateUsers(expectUpdateUser)
	assert.Nil(suite.T(), err)
}

func (suite *UserUCTestSuite) TestGetById_Success() {
	suite.mockUserRepository.On("GetByID", expectUsers.Uuid).Return(expectUsers, nil)

	actualUser, err := suite.userUC.GetByID(expectUsers.Uuid)
	suite.mockUserRepository.AssertExpectations(suite.T())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectUsers, actualUser)
}

func (suite *UserUCTestSuite) TestGetById_FailedInvalidInputOrSqlNoRows() {
	suite.mockUserRepository.On("GetByID", expectUsers.Uuid).Return(dto.GetUsers{}, sql.ErrNoRows)

	actualUser, err := suite.userUC.GetByID(expectUsers.Uuid)
	suite.mockUserRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectUsers, actualUser)
}

func (suite *UserUCTestSuite) TestGetById_FailedInvalid() {
	suite.mockUserRepository.On("GetByID", expectUsers.Uuid).Return(dto.GetUsers{}, errors.New("error"))

	actualUser, err := suite.userUC.GetByID(expectUsers.Uuid)
	suite.mockUserRepository.AssertExpectations(suite.T())
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.NotEqual(suite.T(), expectUsers, actualUser)
}

func (suite *UserUCTestSuite) TestRegisterUser_Success() {
	registerUser := dto.RegisterUsers{
		ID:       expectUsers.Uuid,
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Telp,
	}

	suite.mockUserRepository.On("UsernameIsReady", expectUsers.Username).Return(true, nil)
	suite.mockUserRepository.On("RegisterUsers").Return(nil)

	err := suite.userUC.RegisterUsers(registerUser)
	assert.Nil(suite.T(), err)
}

func (suite *UserUCTestSuite) TestRegisterUser_FailedGetUsernameReady() {
	registerUser := dto.RegisterUsers{
		ID:       expectUsers.Uuid,
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Telp,
	}

	suite.mockUserRepository.On("UsernameIsReady", expectUsers.Username).Return(false, nil)
	suite.mockUserRepository.On("RegisterUsers").Return(nil)

	err := suite.userUC.RegisterUsers(registerUser)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *UserUCTestSuite) TestRegisterUser_FailedGetUsename() {
	registerUser := dto.RegisterUsers{
		ID:       expectUsers.Uuid,
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Telp,
	}

	suite.mockUserRepository.On("UsernameIsReady", expectUsers.Username).Return(false, errors.New("error"))
	suite.mockUserRepository.On("RegisterUsers").Return(nil)

	err := suite.userUC.RegisterUsers(registerUser)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *UserUCTestSuite) TestLoginUser_Success() {
	loginRequest := model.LoginRequest{
		Username: expectUsers.Username,
		Password: "test",
	}
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

	suite.mockUserRepository.On("GetByUsername", loginRequest.Username).Return(user, nil)

	_, err := suite.userUC.LoginUsers(loginRequest)
	assert.Nil(suite.T(), err)
}

func (suite *UserUCTestSuite) TestLoginUser_FailedInvalidInputOrSqlNoRow() {
	loginRequest := model.LoginRequest{
		Username: expectUsers.Username,
		Password: "daniel",
	}
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

	suite.mockUserRepository.On("GetByUsername", loginRequest.Username).Return(user, sql.ErrNoRows)

	_, err := suite.userUC.LoginUsers(loginRequest)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *UserUCTestSuite) TestLoginUser_Failed() {
	loginRequest := model.LoginRequest{
		Username: expectUsers.Username,
		Password: "daniel",
	}
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

	suite.mockUserRepository.On("GetByUsername", loginRequest.Username).Return(user, errors.New("error"))

	_, err := suite.userUC.LoginUsers(loginRequest)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *UserUCTestSuite) TestTopUpSuccess() {
	topUpRequest := dto.TopUpRequest{
		Amount: 100,
		UserID: expectUsers.Uuid,
	}

	suite.mockUserRepository.On("UpdateBalance", topUpRequest).Return(nil)
	err := suite.userUC.TopUp(topUpRequest)

	assert.Nil(suite.T(), err)
}

func (suite *UserUCTestSuite) TestChagePassword_Success() {
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserRepository.On("GetByID", expectUsers.Uuid).Return(expectUsers, nil)
	suite.mockUserRepository.On("UpdatePassword").Return(nil)
	err := suite.userUC.ChangePassword(changePassword)
	assert.Nil(suite.T(), err)
}

func (suite *UserUCTestSuite) TestChagePassword_FailedInvalidInputOrSqlNoRows() {
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserRepository.On("GetByID", expectUsers.Uuid).Return(expectUsers, sql.ErrNoRows)
	suite.mockUserRepository.On("UpdatePassword").Return(nil)
	err := suite.userUC.ChangePassword(changePassword)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *UserUCTestSuite) TestChagePassword_Failed() {
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: "test",
		NewPassword: "test",
	}

	suite.mockUserRepository.On("GetByID", expectUsers.Uuid).Return(expectUsers, errors.New("error"))
	suite.mockUserRepository.On("UpdatePassword").Return(nil)
	err := suite.userUC.ChangePassword(changePassword)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestUserUCTestSuite(t *testing.T) {
	suite.Run(t, new(UserUCTestSuite))
}
