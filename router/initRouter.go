package router

import (
	"bike-rent-express/src/Users/usersDelivery"
	"bike-rent-express/src/Users/usersRepository"
	"bike-rent-express/src/Users/usersUsecase"
	"bike-rent-express/src/motorVehicle/motorVehicleDelivery"
	"bike-rent-express/src/motorVehicle/motorVehicleRepository"
	"bike-rent-express/src/motorVehicle/motorVehicleUsecase"
	"bike-rent-express/src/employee/employeeDelivery"
	"bike-rent-express/src/employee/employeeRepository"
	"bike-rent-express/src/employee/employeeUsecase"
	"bike-rent-express/src/transaction/transactionDelivery"
	"bike-rent-express/src/transaction/transactionRepository"
	"bike-rent-express/src/transaction/transactionUsecase"

	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	usersRepo := usersRepository.NewUsersRepository(db)
	usersUC := usersUsecase.NewUsersUsecase(usersRepo)
	usersDelivery.NewUsersDelivery(v1Group, usersUC)

	motorVehicleRepo := motorVehicleRepository.NewMotorVehicleRepository(db)
	motorVehicleUC := motorVehicleUsecase.NewMotorVehicleUsecase(motorVehicleRepo)
	motorVehicleDelivery.NewMotorVehicleDelivery(v1Group, motorVehicleUC)
	
	transactionRepository := transactionRepository.NewTransactionRepository(db)
	transactionUC := transactionUsecase.NewTransactionRepository(transactionRepository)
	transactionDelivery.NewTransactionUsecase(v1Group, transactionUC)

	employeeRepository := employeeRepository.NewEmployeeRepository(db)
	employeeUC := employeeUsecase.NewEmployeeUsecase(employeeRepository)
	employeeDelivery.NewEmployeeDelivery(v1Group, employeeUC)
}
