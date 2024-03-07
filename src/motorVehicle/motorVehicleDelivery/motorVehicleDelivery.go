package motorVehicleDelivery

import (
	"bike-rent-express/model/dto/json"
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/src/motorVehicle"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type motorVehicleDelivery struct {
	motorVehicleUC motorVehicle.MotorVechileUsecase
}

func NewMotorVehicleDelivery(v1Group *gin.RouterGroup, motorVehicleUC motorVehicle.MotorVechileUsecase) {
	handler := motorVehicleDelivery{
		motorVehicleUC}
	motorVehicleGroup := v1Group.Group("/motorVehicles")
	//motorVehicleGroup.Use(middleware.BasicAuth)
	{
		motorVehicleGroup.GET("/", handler.getAllMotorVehicle)
		motorVehicleGroup.GET("/:id", handler.getMotorVehicleById)
		motorVehicleGroup.POST("/", handler.createMotorVehicle)
		motorVehicleGroup.PUT("/:id", handler.updateMotorVehicle)
		motorVehicleGroup.PUT("delete/:id", handler.deleteMotorVehicle)
	}
}

func (md motorVehicleDelivery) getAllMotorVehicle(ctx *gin.Context) {

	motor, err := md.motorVehicleUC.GetAllMotorVehicle()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, motor, "success", "01", "01")
}

func (md motorVehicleDelivery) getMotorVehicleById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := uuid.Parse(idParam)
	data, err := md.motorVehicleUC.GetMotorVehicleById(id)
	if err != nil && errors.Is(sql.ErrNoRows, err) {
		json.NewResponseBadRequest(ctx, nil, err.Error(), "01", "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success", "01", "01")
}

func (md motorVehicleDelivery) createMotorVehicle(ctx *gin.Context) {
	var input motorVehicleDto.CreateMotorVehicle

	if err := ctx.ShouldBindJSON(&input); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	data, err := md.motorVehicleUC.CreateMotorVehicle(input)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "created", "01", "01")

}

func (md motorVehicleDelivery) updateMotorVehicle(ctx *gin.Context) {
	var input motorVehicleDto.UpdateMotorVehicle

	idParam := ctx.Param("id")
	id, _ := uuid.Parse(idParam)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	data, err := md.motorVehicleUC.UpdateMotorVehicle(id, input)
	if err != nil && errors.Is(sql.ErrNoRows, err) {
		json.NewResponseBadRequest(ctx, nil, err.Error(), "01", "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "updated", "01", "01")
}

func (md motorVehicleDelivery) deleteMotorVehicle(ctx *gin.Context) {
	var input motorVehicleDto.MotorVehicle

	idParam := ctx.Param("id")
	id, _ := uuid.Parse(idParam)

	data, err := md.motorVehicleUC.DeleteMotorVehicle(id, input)
	if err != nil && errors.Is(sql.ErrNoRows, err) {
		json.NewResponseBadRequest(ctx, nil, err.Error(), "01", "01")
		return
	} else if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "soft deleted", "01", "01")
}
