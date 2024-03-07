package motorReturn

import "bike-rent-express/model/dto/motorReturnDto"

type (
	MotorReturnRepository interface {
		Add(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error)
		GetById(id string) (motorReturnDto.MotorReturn, error)
		GetAll() ([]motorReturnDto.MotorReturn, error)
	}

	MotorReturnUsecase interface {
		AddMotorReturn(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error)
		GetMotorReturnById(id string) (motorReturnDto.MotorReturnResponse, error)
		GetMotorReturnAll() ([]motorReturnDto.MotorReturnResponse, error)
	}
)
