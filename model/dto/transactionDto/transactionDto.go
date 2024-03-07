package transactionDto

type Transaction struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	MotorVehicleId string `json:"motor_vehicle_id"`
	EmployeeId     string `json:"employee_id"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
	Price          string `json:"price"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

type AddTransactionRequest struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id" validate:"required"`
	MotorVehicleId string `json:"motor_vehicle_id" validate:"required"`
	EmployeeId     string `json:"employee_id" validate:"required"`
	StartDate      string `json:"start_date" validate:"required,format-date"`
	EndDate        string `json:"end_date" validate:"required,format-date"`
}

type ResponseTransaction struct {
	ID        string `json:"id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Price     int    `json:"price"`
}
