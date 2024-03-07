package transactionUsecase

import (
	"bike-rent-express/model/dto/transactionDto"
	"bike-rent-express/src/Users"
	"bike-rent-express/src/employee"
	"bike-rent-express/src/motorVehicle"
	"bike-rent-express/src/transaction"
	"database/sql"
	"errors"
	"strings"
)

type transactionUsecase struct {
	transactionRepository transaction.TransactionRepository
	userRepository        Users.UsersRepository
	employeeRepository    employee.EmployeeRepository
	vehicleRepository     motorVehicle.MotorVechileRepository
}

func NewTransactionRepository(transactionRepository transaction.TransactionRepository, userRepository Users.UsersRepository, employeeRepository employee.EmployeeRepository, motorVehicleRepository motorVehicle.MotorVechileRepository) transaction.TransactionUsecase {
	return &transactionUsecase{transactionRepository, userRepository, employeeRepository, motorVehicleRepository}
}

func (t *transactionUsecase) AddTransaction(transactionRequest transactionDto.AddTransactionRequest) (transactionDto.Transaction, error) {
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

func (t *transactionUsecase) GetTransactionById(id string) (transactionDto.ResponseTransaction, error) {
	var transactionDetail transactionDto.ResponseTransaction

	transaction, err := t.transactionRepository.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows || strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return transactionDetail, errors.New("1")
		}
		return transactionDetail, err
	}

	motorVehicle, err := t.vehicleRepository.RetrieveMotorVehicleById(transaction.MotorVehicleId)
	if err != nil {
		return transactionDetail, err
	}

	employee, err := t.employeeRepository.GetById(transaction.EmployeeId)
	if err != nil {
		return transactionDetail, err
	}

	customer, err := t.userRepository.GetByID(transaction.UserID)
	if err != nil {
		return transactionDetail, err
	}

	transactionDetail.ID = transaction.ID
	transactionDetail.StartDate = transaction.StartDate
	transactionDetail.EndDate = transaction.EndDate
	transactionDetail.Price = transaction.Price
	transactionDetail.MotorVehicle = motorVehicle
	transactionDetail.Employee = employee
	transactionDetail.Customer = customer
	transactionDetail.CreatedAt = transaction.CreatedAt
	transactionDetail.UpdatedAt = transaction.UpdatedAt

	return transactionDetail, nil
}

func (t *transactionUsecase) GetTransactionAll() ([]transactionDto.ResponseTransaction, error) {
	var transactionsDetail []transactionDto.ResponseTransaction

	transactions, err := t.transactionRepository.GetAll()
	if err != nil {
		return transactionsDetail, err
	}

	for _, transaction := range transactions {
		var transactionDetail transactionDto.ResponseTransaction

		transaction, err := t.transactionRepository.GetById(transaction.ID)
		if err != nil {
			return transactionsDetail, err
		}

		motorVehicle, err := t.vehicleRepository.RetrieveMotorVehicleById(transaction.MotorVehicleId)
		if err != nil {
			return transactionsDetail, err
		}

		employee, err := t.employeeRepository.GetById(transaction.EmployeeId)
		if err != nil {
			return transactionsDetail, err
		}

		customer, err := t.userRepository.GetByID(transaction.UserID)
		if err != nil {
			return transactionsDetail, err
		}

		transactionDetail.ID = transaction.ID
		transactionDetail.StartDate = transaction.StartDate
		transactionDetail.EndDate = transaction.EndDate
		transactionDetail.Price = transaction.Price
		transactionDetail.MotorVehicle = motorVehicle
		transactionDetail.Employee = employee
		transactionDetail.Customer = customer
		transactionDetail.CreatedAt = transaction.CreatedAt
		transactionDetail.UpdatedAt = transaction.UpdatedAt

		transactionsDetail = append(transactionsDetail, transactionDetail)
	}

	return transactionsDetail, nil
}
