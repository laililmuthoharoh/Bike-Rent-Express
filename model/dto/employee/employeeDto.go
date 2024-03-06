package employeeDto

type (
	Employee struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Telp     string `json:"telp"`
		Username string `json:"username"`
		Password string `json:"-"`
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

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken string   `json:"access_token"`
		Employee    Employee `json:"employee"`
	}
)
