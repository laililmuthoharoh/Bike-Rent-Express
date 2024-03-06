package router

import (
	"bike-rent-express/src/Users/usersDelivery"
	"bike-rent-express/src/Users/usersRepository"
	"bike-rent-express/src/Users/usersUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	usersRepo := usersRepository.NewUsersRepository(db)

	usersUC := usersUsecase.NewUsersUsecase(usersRepo)

	usersDelivery.NewUsersDelivery(v1Group, usersUC)
}
