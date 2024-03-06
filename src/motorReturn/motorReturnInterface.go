package motorReturn

import "bike-rent-express/model/dto/motorReturnDto"

type (
	MotorReturnRepository interface {
		Add(createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest) (motorReturnDto.CreateMotorReturnRequest, error)
		GetById(id string) (motorReturnDto.MotorReturn, error)
	}

	MotorReturnUsecase interface{}
)
