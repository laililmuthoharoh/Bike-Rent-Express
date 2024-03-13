package motorReturnUsecase

import (
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/src/Users"
	"bike-rent-express/src/motorReturn"
	"bike-rent-express/src/transaction"
	"database/sql"
	"errors"
	"strings"
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
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return motorReturnCreated, errors.New("3")
		}
		return motorReturnCreated, err
	}

	return motorReturnCreated, nil
}
func (m *motorReturnUsecase) GetMotorReturnById(id string) (motorReturnDto.MotorReturnResponse, error) {
	var motorReturnDetail motorReturnDto.MotorReturnResponse

	motorReturn, err := m.motorReturnRepo.GetById(id)
	if err != nil {
		return motorReturnDetail, err
	}

	transaction, err := m.transactionRepo.GetById(motorReturn.TrasactionID)
	if err != nil {
		return motorReturnDetail, err
	}

	user, err := m.userRepo.GetByID(transaction.UserID)
	if err != nil {
		return motorReturnDetail, err
	}

	motorReturnDetail.ID = motorReturn.ID
	motorReturnDetail.ReturnDate = motorReturn.ReturnDate
	motorReturnDetail.ExtraCharge = motorReturn.ExtraCharge
	motorReturnDetail.ConditionMotor = motorReturn.ConditionMotor
	motorReturnDetail.Descrption = motorReturn.Descrption
	motorReturnDetail.CreatedAt = motorReturn.CreatedAt
	motorReturnDetail.UpdatedAt = motorReturn.UpdatedAt
	motorReturnDetail.Customer = user

	return motorReturnDetail, nil
}
func (m *motorReturnUsecase) GetMotorReturnAll() ([]motorReturnDto.MotorReturnResponse, error) {
	var motorsReturnDetail []motorReturnDto.MotorReturnResponse

	motorsReturn, err := m.motorReturnRepo.GetAll()
	if err != nil {
		return motorsReturnDetail, err
	}

	for _, motorReturn := range motorsReturn {
		var motorReturnDetail motorReturnDto.MotorReturnResponse

		motorReturn, err := m.motorReturnRepo.GetById(motorReturn.ID)
		if err != nil {
			return motorsReturnDetail, err
		}

		transaction, err := m.transactionRepo.GetById(motorReturn.TrasactionID)
		if err != nil {
			return motorsReturnDetail, err
		}

		user, err := m.userRepo.GetByID(transaction.UserID)
		if err != nil {
			return motorsReturnDetail, err
		}
		motorReturnDetail.ID = motorReturn.ID
		motorReturnDetail.ReturnDate = motorReturn.ReturnDate
		motorReturnDetail.ExtraCharge = motorReturn.ExtraCharge
		motorReturnDetail.ConditionMotor = motorReturn.ConditionMotor
		motorReturnDetail.Descrption = motorReturn.Descrption
		motorReturnDetail.CreatedAt = motorReturn.CreatedAt
		motorReturnDetail.UpdatedAt = motorReturn.UpdatedAt
		motorReturnDetail.Customer = user

		motorsReturnDetail = append(motorsReturnDetail, motorReturnDetail)
	}

	return motorsReturnDetail, nil
}
