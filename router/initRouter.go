package router

import (
	"bike-rent-express/src/motorVehicle/motorVehicleDelivery"
	"bike-rent-express/src/motorVehicle/motorVehicleRepository"
	"bike-rent-express/src/motorVehicle/motorVehicleUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	{
		motorVehicleRepo := motorVehicleRepository.NewMotorVehicleRepository(db)
		motorVehicleUC := motorVehicleUsecase.NewMotorVehicleUsecase(motorVehicleRepo)
		motorVehicleDelivery.NewMotorVehicleDelivery(v1Group, motorVehicleUC)
	}
}
