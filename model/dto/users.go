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
		Created_at string
		Updated_at string
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
)
