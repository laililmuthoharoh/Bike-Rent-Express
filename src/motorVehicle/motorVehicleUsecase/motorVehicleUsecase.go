package motorVehicleUsecase

import (
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/src/motorVehicle"
	"errors"
	"strings"
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
func (mu motorVehicleUsecase) GetMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error) {
	motor, err := mu.motorVehicleRepo.RetrieveMotorVehicleById(id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return motor, errors.New("1")
		}
		return motor, err
	}

	return motor, nil
}

func (mu motorVehicleUsecase) CreateMotorVehicle(motor motorVehicleDto.CreateMotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	ready, err := mu.motorVehicleRepo.CheckPlatMotor(motor.Plat)
	if err != nil {
		return motorVehicleDto.MotorVehicle{}, err
	}

	if !ready {
		return motorVehicleDto.MotorVehicle{}, errors.New("1")
	}

	newMotor, err := mu.motorVehicleRepo.InsertMotorVehicle(motorVehicleDto.MotorVehicle{
		Name:           motor.Name,
		Type:           motor.Type,
		Price:          motor.Price,
		Plat:           motor.Plat,
		ProductionYear: motor.ProductionYear,
		Status:         motor.Status,
	})

	if err != nil {
		return newMotor, err
	}

	return newMotor, nil
}

func (mu motorVehicleUsecase) UpdateMotorVehicle(id string, input motorVehicleDto.UpdateMotorVehicle) (motorVehicleDto.MotorVehicle, error) {
	motor, err := mu.motorVehicleRepo.RetrieveMotorVehicleById(id)
	if err != nil {
		return motor, err
	}

	if input.Plat != motor.Plat {
		ready, err := mu.motorVehicleRepo.CheckPlatMotor(input.Plat)
		if err != nil {
			return motorVehicleDto.MotorVehicle{}, err
		}

		if !ready {
			return motorVehicleDto.MotorVehicle{}, errors.New("1")
		}

	}

	if input.Name != "" {
		motor.Name = input.Name
	}
	if input.Type != "" {
		motor.Type = input.Type
	}
	if input.Price != 0 {
		motor.Price = input.Price
	}
	if input.Plat != "" {
		motor.Plat = input.Plat
	}
	if input.ProductionYear != "" {
		motor.ProductionYear = input.ProductionYear
	}
	if input.Status != "" {
		motor.Status = input.Status
	}

	data, err := mu.motorVehicleRepo.ChangeMotorVehicle(id, motor)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (mu motorVehicleUsecase) DeleteMotorVehicle(id string) error {
	motor, err := mu.motorVehicleRepo.RetrieveMotorVehicleById(id)
	if err != nil {
		return err
	}
	if motor.Status == "NOT_AVAILABLE" {
		return errors.New("1")
	}
	err = mu.motorVehicleRepo.DropMotorVehicle(id)
	if err != nil {
		return err
	}

	return nil

}
