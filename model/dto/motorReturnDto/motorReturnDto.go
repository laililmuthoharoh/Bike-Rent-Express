package motorReturnDto

import "bike-rent-express/model/dto"

type (
	MotorReturn struct {
		ID             string `json:"id"`
		TrasactionID   string `json:"transaction_id"`
		ReturnDate     string `json:"return_date"`
		ExtraCharge    int    `json:"extra_charge"`
		ConditionMotor string `json:"condition_motor"`
		Descrption     string `json:"description"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updatad_at"`
	}

	CreateMotorReturnRequest struct {
		ID             string `json:"id"`
		TransactionID  string `json:"transaction_id" validate:"required"`
		ExtraCharge    int    `json:"extra_charge" validate:"required"`
		ConditionMotor string `json:"condition_motor" validate:"required"`
		Description    string `json:"description" validate:"required"`
	}

	MotorReturnResponse struct {
		ID             string       `json:"id"`
		ReturnDate     string       `json:"return_date"`
		ExtraCharge    int          `json:"extra_charge"`
		ConditionMotor string       `json:"condition_motor"`
		Descrption     string       `json:"description"`
		Customer       dto.GetUsers `json:"customer"`
		CreatedAt      string       `json:"created_at"`
		UpdatedAt      string       `json:"updatad_at"`
	}
)
