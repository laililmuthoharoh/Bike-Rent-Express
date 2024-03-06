package router

import (
	transactiondelivery "bike-rent-express/src/transaction/transactionDelivery"
	"bike-rent-express/src/transaction/transactionRepository"
	"bike-rent-express/src/transaction/transactionUsecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	transactionRepository := transactionRepository.NewTransactionRepository(db)
	transactionUC := transactionUsecase.NewTransactionRepository(transactionRepository)
	transactiondelivery.NewTransactionUsecase(v1Group, transactionUC)
}
