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
	UserID         string `json:"user_id"`
	MotorVehicleId string `json:"motor_vehicle_id"`
	EmployeeId     string `json:"employee_id"`
	StartDate      string `json:"start_date"`
	EndDate        string `json:"end_date"`
}
