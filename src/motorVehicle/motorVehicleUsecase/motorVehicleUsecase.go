package motorVehicleUsecase

import (
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/src/motorVehicle"

	"github.com/google/uuid"
)

type motorVehicleUsecase struct {
	motorVehicleRepo motorVehicle.MotorVechileRepository
}

func NewMotorVehicleUsecase(motorVehicleRepo motorVehicle.MotorVechileRepository) motorVehicle.MotorVechileUsecase {
	return &motorVehicleUsecase{motorVehicleRepo}
}

// get all motors
func (mu motorVehicleUsecase) GetAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error) {

	motors, err := mu.motorVehicleRepo.RetrieveAllMotorVehicle()
	if err != nil {
		return motors, err
	}
	return motors, nil
}

// get by id
func (mu motorVehicleUsecase) GetMotorVehicleById(id uuid.UUID) (motorVehicleDto.MotorVehicle, error) {
	motor, err := mu.motorVehicleRepo.RetrieveMotorVehicleById(id)
	if err != nil {
		return motor, err
	}

	return motor, nil
}

func (mu motorVehicleUsecase) CreateMotorVehicle(motor motorVehicleDto.CreateMotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	dt, err := mu.motorVehicleRepo.InsertMotorVehicle(motorVehicleDto.MotorVehicle{
		Name:           motor.Name,
		Type:           motor.Type,
		Price:          motor.Price,
		Plat:           motor.Plat,
		ProductionYear: motor.ProductionYear,
		Status:         motor.Status,
	})
	if err != nil {
		return dt, err
	}

	return dt, nil
}

func (mu motorVehicleUsecase) UpdateMotorVehicle(id uuid.UUID, input motorVehicleDto.UpdateMotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	motor, err := mu.motorVehicleRepo.RetrieveMotorVehicleById(id)
	if err != nil {
		return motor, err
	}

	if input.Name != "" {
		motor.Name = input.Name
	} else if input.Type != "" {
		motor.Type = input.Type
	} else if input.Price != 0 {
		motor.Price = input.Price
	} else if input.Plat != "" {
		motor.Plat = input.Plat
	} else if input.ProductionYear != "" {
		motor.ProductionYear = input.ProductionYear
	} else if input.Status != "" {
		motor.Status = input.Status
	}

	data, err := mu.motorVehicleRepo.ChangeMotorVehicle(id, motor)
	if err != nil {
		return data, nil
	}
	return data, nil
}

func (mu motorVehicleUsecase) DeleteMotorVehicle(id uuid.UUID) (motorVehicleDto.MotorVehicle, error) {
	data, err := mu.motorVehicleRepo.RetrieveMotorVehicleById(id)
	if err != nil {
		return data, err
	}

	err = mu.motorVehicleRepo.DropMotorVehicle(id)
	if err != nil {
		return data, err
	}

	return data, nil

}
