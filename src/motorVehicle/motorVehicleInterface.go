package motorVehicle

import (
	"bike-rent-express/model/dto/motorVehicleDto"

	"github.com/google/uuid"
)

type (
	MotorVechileRepository interface {
		RetrieveAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error)
		RetrieveMotorVehicleById(id uuid.UUID) (motorVehicleDto.MotorVehicle, error)
		InsertMotorVehicle(motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error)
		ChangeMotorVehicle(id uuid.UUID, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error)
		DropMotorVehicle(id uuid.UUID, motor motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error)
	}

	MotorVechileUsecase interface {
		GetAllMotorVehicle() ([]motorVehicleDto.MotorVehicle, error)
		GetMotorVehicleById(id uuid.UUID) (motorVehicleDto.MotorVehicle, error)
		CreateMotorVehicle(motor motorVehicleDto.CreateMotorVehicle) (motorVehicleDto.MotorVehicle, error)
		UpdateMotorVehicle(id uuid.UUID, motor motorVehicleDto.UpdateMotorVehicle) (motorVehicleDto.MotorVehicle, error)
		DeleteMotorVehicle(id uuid.UUID, input motorVehicleDto.MotorVehicle) (motorVehicleDto.MotorVehicle, error)
	}
)
