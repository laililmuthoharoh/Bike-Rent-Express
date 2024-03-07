package motorReturnUsecase

import (
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/src/Users"
	"bike-rent-express/src/motorReturn"
	"bike-rent-express/src/transaction"
)

type motorReturnUsecase struct {
	motorReturnRepo motorReturn.MotorReturnRepository
	transactionRepo transaction.TransactionRepository
	userRepo        Users.UsersRepository
}

func NewMotorReturnUseCase(motorReturnRepo motorReturn.MotorReturnRepository, transactionRepo transaction.TransactionRepository, userRepo Users.UsersRepository) motorReturn.MotorReturnUsecase {
	return &motorReturnUsecase{motorReturnRepo, transactionRepo, userRepo}
}

func (m *motorReturnUsecase) AddMotorReturn(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error) {
	motorReturnCreated, err := m.motorReturnRepo.Add(createMotorReturnRequest)
	if err != nil{
		return motorReturnCreated, err
	}
	
	return motorReturnCreated, nil
}
func (m *motorReturnUsecase) GetMotorReturnById(id string) (motorReturnDto.MotorReturnResponse, error) {
	var motorReturnDetail motorReturnDto.MotorReturnResponse
	
	motorReturn, err := m.motorReturnRepo.GetById(id)
	if err != nil{
		return motorReturnDetail, err
	}

	transaction, err := m.transactionRepo.GetById(motorReturn.TrasactionID)
	if err != nil{
		return motorReturnDetail, err
	}

	user, err := m.userRepo.GetByID(transaction.)

	
	return motorReturnDto.MotorReturnResponse{}, nil
}
func (m *motorReturnUsecase) GetMotorReturnAll() ([]motorReturnDto.MotorReturnResponse, error) {
	return []motorReturnDto.MotorReturnResponse{}, nil
}
