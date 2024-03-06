package motorReturnDto

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
		ID string
		// Transaction
	}
)
