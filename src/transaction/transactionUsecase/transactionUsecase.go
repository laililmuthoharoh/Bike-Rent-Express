package transactionUsecase

import (
	"bike-rent-express/model/dto/transactionDto"
	"bike-rent-express/src/transaction"
	"fmt"
)

type transactionUsecase struct {
	transactionRepository transaction.TransactionRepository
}

func NewTransactionRepository(transactionRepository transaction.TransactionRepository) transaction.TransactionUsecase {
	return &transactionUsecase{transactionRepository}
}

func (t *transactionUsecase) AddTransaction(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.Transaction, error) {
	fmt.Println("ceks")
	resultTransactionCreated, err := t.transactionRepository.Add(transactionRequest)
	if err != nil {
		return transactionDto.Transaction{}, err
	}

	transaction, err := t.transactionRepository.GetById(resultTransactionCreated.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
