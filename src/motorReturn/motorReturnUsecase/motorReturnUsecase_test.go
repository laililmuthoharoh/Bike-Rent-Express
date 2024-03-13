package motorReturnUsecase

import (
	"bike-rent-express/model/dto"
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/model/dto/transactionDto"
	"bike-rent-express/src/motorReturn"
	"errors"
	"testing"

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

var expectedAllMotorReturn = []motorReturnDto.MotorReturn{
	expectedMotorReturn,
}

var expectedAllMotorReturnResponse = []motorReturnDto.MotorReturnResponse{
	expectedMotorReturnResponse,
}

type mockMotorReturnRepository struct {
	mock.Mock
}

func (m *mockMotorReturnRepository) Add(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error) {
	arg := m.Called(createMotorReturnRequest)
	return arg.Get(0).(motorReturnDto.CreateMotorReturnRequest), arg.Error(1)
}

func (m *mockMotorReturnRepository) GetById(id string) (motorReturnDto.MotorReturn, error) {
	arg := m.Called(id)
	return arg.Get(0).(motorReturnDto.MotorReturn), arg.Error(1)
}

func (m *mockMotorReturnRepository) GetAll() ([]motorReturnDto.MotorReturn, error) {
	arg := m.Called()
	return arg.Get(0).([]motorReturnDto.MotorReturn), arg.Error(1)
}

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

type MotorReturnUsecaseTestSuite struct {
	suite.Suite
	mockMotorReturnRepository *mockMotorReturnRepository
	mockTransactionRepository *mockTransactionRepository
	mockUserRepository        *mockUserRepository
	motorReturnUsecase        motorReturn.MotorReturnUsecase
}

func (suite *MotorReturnUsecaseTestSuite) SetupTest() {
	suite.mockMotorReturnRepository = new(mockMotorReturnRepository)
	suite.mockTransactionRepository = new(mockTransactionRepository)
	suite.mockUserRepository = new(mockUserRepository)
	suite.motorReturnUsecase = NewMotorReturnUseCase(suite.mockMotorReturnRepository, suite.mockTransactionRepository, suite.mockUserRepository)
}

// test add success
func (suite *MotorReturnUsecaseTestSuite) TestAddMotorReturn_Success() {
	expectedCreateMotorReturn := motorReturnDto.CreateMotorReturnRequest{
		ID:             expectedMotorReturn.ID,
		TransactionID:  expectedMotorReturn.TrasactionID,
		ExtraCharge:    expectedMotorReturn.ExtraCharge,
		ConditionMotor: expectedMotorReturn.ConditionMotor,
		Description:    expectedMotorReturn.Descrption,
	}

	suite.mockMotorReturnRepository.On("Add", expectedCreateMotorReturn).Return(expectedCreateMotorReturn, nil)

	actual, err := suite.motorReturnUsecase.AddMotorReturn(expectedCreateMotorReturn)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCreateMotorReturn, actual)
}

// test add fail
func (suite *MotorReturnUsecaseTestSuite) TestAddMotorReturn_Fail() {
	expectedCreateMotorReturn := motorReturnDto.CreateMotorReturnRequest{
		ID:             expectedMotorReturn.ID,
		TransactionID:  expectedMotorReturn.TrasactionID,
		ExtraCharge:    expectedMotorReturn.ExtraCharge,
		ConditionMotor: expectedMotorReturn.ConditionMotor,
		Description:    expectedMotorReturn.Descrption,
	}
	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("Add", expectedCreateMotorReturn).Return(expectedCreateMotorReturn, expectedError)

	_, err := suite.motorReturnUsecase.AddMotorReturn(expectedCreateMotorReturn)

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// test get by id success
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnById_Success() {

	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockUserRepository.On("GetByID", expectedCustomer.Uuid).Return(expectedCustomer, nil)

	actual, err := suite.motorReturnUsecase.GetMotorReturnById(expectedMotorReturnResponse.ID)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedMotorReturnResponse, actual)
}

// test get by id fail mock motor return
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnById_FailMotorReturn() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnById(expectedMotorReturnResponse.ID)

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// test get by id fail mock transaction
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnById_FailTransaction() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnById(expectedMotorReturnResponse.ID)

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// test get by id fail mock User
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnById_FailUser() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockUserRepository.On("GetByID", expectedCustomer.Uuid).Return(expectedCustomer, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnById(expectedMotorReturnResponse.ID)

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// get all success
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnAll_Success() {

	suite.mockMotorReturnRepository.On("GetAll").Return(expectedAllMotorReturn, nil)
	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockUserRepository.On("GetByID", expectedCustomer.Uuid).Return(expectedCustomer, nil)

	actual, err := suite.motorReturnUsecase.GetMotorReturnAll()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAllMotorReturnResponse, actual)
}

// get all fail
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnAll_fail() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetAll").Return(expectedAllMotorReturn, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnAll()

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// get all fail get Motor Return
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnAll_failGetMotorReturn() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetAll").Return(expectedAllMotorReturn, nil)
	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnAll()

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// get all fail get transaction
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnAll_failGetTransaction() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetAll").Return(expectedAllMotorReturn, nil)
	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnAll()

	assert.EqualError(suite.T(), err, expectedError.Error())
}

// get all fail get User
func (suite *MotorReturnUsecaseTestSuite) TestGetMotorReturnAll_failGetUser() {

	expectedError := errors.New("mock error")

	suite.mockMotorReturnRepository.On("GetAll").Return(expectedAllMotorReturn, nil)
	suite.mockMotorReturnRepository.On("GetById", expectedMotorReturn.ID).Return(expectedMotorReturn, nil)
	suite.mockTransactionRepository.On("GetById", expectTransaction.ID).Return(expectTransaction, nil)
	suite.mockUserRepository.On("GetByID", expectedCustomer.Uuid).Return(expectedCustomer, expectedError)

	_, err := suite.motorReturnUsecase.GetMotorReturnAll()

	assert.EqualError(suite.T(), err, expectedError.Error())
}

func TestMotorReturnUsecase(t *testing.T) {
	suite.Run(t, new(MotorReturnUsecaseTestSuite))
}
