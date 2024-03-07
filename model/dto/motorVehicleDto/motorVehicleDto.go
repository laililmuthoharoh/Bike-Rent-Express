package motorVehicleDto

type (
	MotorVehicle struct {
		Id             string `json:"id,omitempty"`
		Name           string `json:"name,omitempty"`
		Type           string `json:"type,omitempty"`
		Price          int    `json:"price,omitempty"`
		Plat           string `json:"plat,omitempty"`
		CreatedAt      string `json:"created_at"`
		UpdatedAt      string `json:"updated_at"`
		ProductionYear string `json:"production_year,omitempty"`
		Status         string `json:"status,omitempty"`
	}

	CreateMotorVehicle struct {
		Name           string `json:"name" validate:"required"`
		Type           string `json:"type" validate:"required"`
		Price          int    `json:"price" validate:"required"`
		Plat           string `json:"plat" validate:"required"`
		ProductionYear string `json:"production_year" validate:"required"`
		Status         string `json:"status" validate:"required,status-valid"`
	}

	UpdateMotorVehicle struct {
		Name           string `json:"name" validate:"required"`
		Type           string `json:"type" validate:"required"`
		Price          int    `json:"price" validate:"required"`
		Plat           string `json:"plat" validate:"required"`
		ProductionYear string `json:"production_year" validate:"required"`
		Status         string `json:"status" validate:"required,status-valid"`
	}
)
