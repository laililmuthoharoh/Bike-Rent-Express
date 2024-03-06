package motorVehicleDto

import (
	"time"

	"github.com/google/uuid"
)

type (
	MotorVehicle struct {
		Id             uuid.UUID `json:"id,omitempty"`
		Name           string    `json:"name,omitempty"`
		Type           string    `json:"type,omitempty"`
		Price          int       `json:"price,omitempty"`
		Plat           string    `json:"plate,omitempty"`
		CreatedAt      time.Time `json:"created_at,omitempty"`
		UpdatedAt      time.Time `json:"updated_at,omitempty"`
		ProductionYear string    `json:"production_year,omitempty"`
		Status         string    `json:"status,omitempty"`
		DeletedAt      time.Time `json:"deleted_at,omitempty"`
	}

	CreateMotorVehicle struct {
		Name           string `json:"name" binding:"required"`
		Type           string `json:"type" binding:"required"`
		Price          int    `json:"price" binding:"required"`
		Plat           string `json:"plat" binding:"required"`
		ProductionYear string `json:"production_year" binding:"required"`
		Status         string `json:"status" binding:"required"`
		//status belum auto available: status vehicle_status DEFAULT 'AVAILABLE' NOT NULL
	}

	UpdateMotorVehicle struct {
		Name           string `json:"name"`
		Type           string `json:"type"`
		Price          int    `json:"price"`
		Plat           string `json:"plat"`
		ProductionYear string `json:"production_year"`
		Status         string `json:"status"`
	}
)
