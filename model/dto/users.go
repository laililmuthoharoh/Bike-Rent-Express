package dto

type (
	RegisterUsers struct {
		Uuid       string `json:"id"`
		Name       string `json:"nama"`
		Username   string `json:"username"`
		Password   string `json:"pass"`
		Address    string `json:"alamat"`
		Role       string `json:"role"`
		Can_rent   string `json:"sewa"`
		Created_at string
		Telp       string `json:"telepon"`
	}

	GetUsers struct {
		Uuid       string `json:"id"`
		Name       string `json:"nama"`
		Username   string `json:"username"`
		Address    string `json:"alamat"`
		Role       string `json:"role"`
		Can_rent   string `json:"sewa"`
		Password   string `json:"-"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
		Telp       string `json:"telepon"`
	}

	Users struct {
		Uuid       string `json:"id"`
		Name       string `json:"nama"`
		Username   string `json:"username"`
		Password   string `json:"pass"`
		Address    string `json:"alamat"`
		Role       string `json:"role"`
		Can_rent   string `json:"sewa"`
		Updated_at string
		Telp       string `json:"telepon"`
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
