package employeeDto

type (
	Employee struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Telp     string `json:"telp"`
		Username string `json:"username"`
	}

	UpdateEmployeeRequest struct {
		ID       string `json:"id"`
		Name     string `json:"name" validate:"required"`
		Telp     string `json:"telp" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	CreateEmployeeRequest struct {
		ID       string `json:"id"`
		Name     string `json:"name" validate:"required"`
		Telp     string `json:"telp" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)
