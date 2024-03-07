package transaction

import "bike-rent-express/model/dto/transactionDto"

type (
	TransactionRepository interface {
		Add(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.AddTransactionRequest, error)
		GetById(id string) (transactionDto.Transaction, error)
		GetAll() ([]transactionDto.Transaction, error)
	}

	TransactionUsecase interface {
		AddTransaction(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.Transaction, error)
		GetTransactionById(id string) (transactionDto.ResponseTransaction, error)
		GetTransactionAll() ([]transactionDto.ResponseTransaction, error)
	}
)
