package employeeDto

type (
	Employee struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Telp      string `json:"telp"`
		Username  string `json:"username"`
		Password  string `json:"-"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	UpdateEmployeeRequest struct {
		ID   string `json:"id"`
		Name string `json:"name" validate:"required"`
		Telp string `json:"telp" validate:"required"`
	}

	CreateEmployeeRequest struct {
		ID       string `json:"id"`
		Name     string `json:"name" validate:"required"`
		Telp     string `json:"telp" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginRequest struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		AccessToken string   `json:"access_token"`
		Employee    Employee `json:"employee"`
	}

	ChangePasswordRequest struct {
		PasswordOld string `json:"password_old" validate:"required"`
		NewPassword string `json:"new_password" validate:"required"`
	}
)
