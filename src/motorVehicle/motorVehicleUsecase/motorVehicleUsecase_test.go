package motorVehicleUsecase

import (
	"bike-rent-express/model/dto/motorVehicleDto"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type mockMotorVehicleRepository struct {
	mock.Mock
}

func (m *mockMotorVehicleRepository) RetrieveAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error) {
	arg := m.Called()
	return arg.Get(0).([]motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleRepository) RetrieveMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error) {
	arg := m.Called(id)
	return arg.Get(0).(motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleRepository) InsertMotorVehicle(motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	arg := m.Called(expectedMotorVehicleById)
	return arg.Get(0).(motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleRepository) ChangeMotorVehicle(id string, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	arg := m.Called(id, expectedMotorVehicleById)
	return arg.Get(0).(motorVehicleDto.MotorVehicle), arg.Error(1)
}

func (m *mockMotorVehicleRepository) DropMotorVehicle(id string) error {
	arg := m.Called(id)
	return arg.Error(0)
}

func TestGetAllMotorVehicle_Success(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expected := []motorVehicleDto.MotorVehicle{expectedMotorVehicleById}

	mockRepo.On("RetrieveAllMotorVehicle").Return(expected, nil)

	usecase := NewMotorVehicleUsecase(mockRepo)

	result, err := usecase.GetAllMotorVehicle()

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestGetAllMotorVehicle_Fail(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expected := []motorVehicleDto.MotorVehicle{expectedMotorVehicleById}
	expectedError := errors.New("mock error")

	mockRepo.On("RetrieveAllMotorVehicle").Return(expected, expectedError)

	usecase := NewMotorVehicleUsecase(mockRepo)

	_, err := usecase.GetAllMotorVehicle()

	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetMotorVehicleById_Success(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	mockRepo.On("RetrieveMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, nil)

	usecase := NewMotorVehicleUsecase(mockRepo)

	result, err := usecase.GetMotorVehicleById(expectedMotorVehicleById.Id)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedMotorVehicleById, result)
}

func TestGetMotorVehicleById_Fail(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expectedError := errors.New("mock error")

	mockRepo.On("RetrieveMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, expectedError)

	usecase := NewMotorVehicleUsecase(mockRepo)

	_, err := usecase.GetMotorVehicleById(expectedMotorVehicleById.Id)

	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, expectedError.Error())
}

func TestCreateMotorVehicle_Success(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expected := motorVehicleDto.CreateMotorVehicle{
		Name:           "Vario",
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	}

	mockRepo.On("InsertMotorVehicle", expectedMotorVehicleById).Return(expectedMotorVehicleById, nil)

	usecase := NewMotorVehicleUsecase(mockRepo)

	result, err := usecase.CreateMotorVehicle(expected)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedMotorVehicleById, result)
}

func TestCreateMotorVehicle_Fail(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expected := motorVehicleDto.CreateMotorVehicle{
		Name:           "Vario",
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	}
	expectedError := errors.New("mock error")

	mockRepo.On("InsertMotorVehicle", expectedMotorVehicleById).Return(expectedMotorVehicleById, expectedError)

	usecase := NewMotorVehicleUsecase(mockRepo)

	_, err := usecase.CreateMotorVehicle(expected)

	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, expectedError.Error())
}

func TestUpdateMotorVehicle_Success(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expected := motorVehicleDto.UpdateMotorVehicle{
		Name:           "Vario",
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	}

	mockRepo.On("RetrieveMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, nil)

	mockRepo.On("ChangeMotorVehicle", expectedMotorVehicleById.Id, expectedMotorVehicleById).Return(expectedMotorVehicleById, nil)

	usecase := NewMotorVehicleUsecase(mockRepo)

	result, err := usecase.UpdateMotorVehicle(expectedMotorVehicleById.Id, expected)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedMotorVehicleById, result)
}

func TestUpdateMotorVehicle_fail(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expected := motorVehicleDto.UpdateMotorVehicle{
		Name:           "Vario",
		Type:           "MATIC",
		Price:          50000,
		Plat:           "BA1234I",
		ProductionYear: "2023",
		Status:         "AVAILABLE",
	}
	expectedError := errors.New("mock error")

	mockRepo.On("RetrieveMotorVehicleById", expectedMotorVehicleById.Id).Return(expectedMotorVehicleById, nil)

	mockRepo.On("ChangeMotorVehicle", expectedMotorVehicleById.Id, expectedMotorVehicleById).Return(expectedMotorVehicleById, expectedError)

	usecase := NewMotorVehicleUsecase(mockRepo)

	_, err := usecase.UpdateMotorVehicle(expectedMotorVehicleById.Id, expected)

	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, expectedError.Error())
}

func TestDeleteMotorVehicle_Success(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	mockRepo.On("DropMotorVehicle", expectedMotorVehicleById.Id).Return(nil)

	usecase := NewMotorVehicleUsecase(mockRepo)

	err := usecase.DeleteMotorVehicle(expectedMotorVehicleById.Id)

	mockRepo.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestDeleteMotorVehicle_Fail(t *testing.T) {
	mockRepo := new(mockMotorVehicleRepository)

	expectedError := errors.New("mock error")

	mockRepo.On("DropMotorVehicle", expectedMotorVehicleById.Id).Return(expectedError)

	usecase := NewMotorVehicleUsecase(mockRepo)

	err := usecase.DeleteMotorVehicle(expectedMotorVehicleById.Id)

	mockRepo.AssertExpectations(t)
	assert.EqualError(t, err, expectedError.Error())
}
