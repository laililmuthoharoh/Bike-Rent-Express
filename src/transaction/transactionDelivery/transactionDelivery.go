package transactionDelivery

import (
	"bike-rent-express/model/dto/json"
	"bike-rent-express/model/dto/transactionDto"
	"bike-rent-express/pkg/utils"
	"bike-rent-express/src/transaction"

	"github.com/gin-gonic/gin"
)

type transactionDelivery struct {
	transactionUC transaction.TransactionUsecase
}

func NewTransactionUsecase(v1Group *gin.RouterGroup, transactionUC transaction.TransactionUsecase) {
	handler := transactionDelivery{transactionUC}

	transactionGroup := v1Group.Group("/user/transaction")
	{
		transactionGroup.POST("", handler.CreateTransaction)
		transactionGroup.GET("/:id", handler.CreateTransaction)
		transactionGroup.GET("", handler.CreateTransaction)
	}
}

func (t *transactionDelivery) CreateTransaction(c *gin.Context) {
	var transactionRequest transactionDto.AddTransactionRequest

	c.BindJSON(&transactionRequest)

	if err := utils.Validated(transactionRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "01", "01")
		return
	}

	resultTransaction, err := t.transactionUC.AddTransaction(transactionRequest)

	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	json.NewResponseCreated(c, resultTransaction, "Transaction Created", "01", "01")
}

func (t *transactionDelivery) GetTransactionById(c *gin.Context) {
	id := c.Param("id")

	transactionDetail, err := t.transactionUC.GetTransactionById(id)
	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "01")
		return
	}

	json.NewResponseSuccess(c, transactionDetail, "Success get transaction by id", "02", "01")
}

func (t *transactionDelivery) GetTransactionAll(c *gin.Context) {
	transactionsDetail, err := t.transactionUC.GetTransactionAll()
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "01")
		return
	}

	json.NewResponseSuccess(c, transactionsDetail, "Success get all transaction", "02", "01")
}
