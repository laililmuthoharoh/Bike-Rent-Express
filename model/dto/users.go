package dto

type (
	RegisterUsers struct {
		ID       string `json:"-"`
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Address  string `json:"address" validate:"required"`
		Role     string `json:"role" validate:"required"`
		Telp     string `json:"telp" validate:"required"`
	}

	GetUsers struct {
		Uuid       string `json:"id"`
		Name       string `json:"nama"`
		Username   string `json:"username"`
		Address    string `json:"alamat"`
		Role       string `json:"role"`
		Can_rent   bool   `json:"cant_rent"`
		Password   string `json:"-"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
		Telp       string `json:"telepon"`
	}

	Users struct {
		ID         string `json:"id"`
		Name       string `json:"name" validate:"required"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Address    string `json:"address" validate:"required"`
		Role       string `json:"role"`
		CanRent    bool   `json:"can_rent" validate:"required"`
		Updated_at string
		Telp       string `json:"telp" validate:"required"`
	}

	LoginResponse struct {
		AccesToken string `json:"acces_token"`
		User       Users  `json:"user"`
	}

	TopUpRequest struct {
		Amount int `json:"amount" validate:"required"`
		UserID string
	}

	ChangePassword struct {
		ID          string
		OldPassword string `json:"old_password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required"`
	}
)
