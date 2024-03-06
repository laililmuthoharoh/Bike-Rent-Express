package transaction

import "bike-rent-express/model/dto/transactionDto"

type (
	TransactionRepository interface {
		Add(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.AddTransactionRequest, error)
	}

	TransactionUsecase interface {
		AddTransaction(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.AddTransactionRequest, error)
	}
)
