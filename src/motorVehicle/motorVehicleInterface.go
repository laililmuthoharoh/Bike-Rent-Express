package motorVehicle

import (
	"bike-rent-express/model/dto/motorVehicleDto"
)

type (
	MotorVechileRepository interface {
		RetrieveAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error)
		RetrieveMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error)
		InsertMotorVehicle(motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error)
		ChangeMotorVehicle(id string, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error)
		DropMotorVehicle(id string) error
		CheckPlatMotor(plat string) (bool, error)
	}

	MotorVechileUsecase interface {
		GetAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error)
		GetMotorVehicleById(id string) (motorVehicleDto.MotorVehicle, error)
		CreateMotorVehicle(motor motorVehicleDto.CreateMotorVehicle) (motorVehicleDto.MotorVehicle, error)
		UpdateMotorVehicle(id string, motor motorVehicleDto.UpdateMotorVehicle) (motorVehicleDto.MotorVehicle, error)
		DeleteMotorVehicle(id string) error
	}
)
