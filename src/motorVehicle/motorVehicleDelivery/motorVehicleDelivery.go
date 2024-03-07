package motorVehicleDelivery

import (
	"bike-rent-express/model/dto/json"
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/pkg/utils"
	"bike-rent-express/src/motorVehicle"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
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
		motorVehicleGroup.DELETE("/:id", handler.deleteMotorVehicle)
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
	id := ctx.Param("id")

	data, err := md.motorVehicleUC.GetMotorVehicleById(id)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			json.NewResponseSuccess(ctx, nil, "Data not found", "02", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "02", "01")
		return
	}

	json.NewResponseSuccess(ctx, data, "success get data by id", "02", "02")
}

func (md motorVehicleDelivery) createMotorVehicle(ctx *gin.Context) {
	var input motorVehicleDto.CreateMotorVehicle

	ctx.ShouldBindJSON(&input)
	if err := utils.Validated(input); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad Request", "03", "01")
		return
	}

	data, err := md.motorVehicleUC.CreateMotorVehicle(input)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "01")
		return
	}

	json.NewResponseCreated(ctx, data, "motor vehicle created", "03", "01")

}

func (md motorVehicleDelivery) updateMotorVehicle(ctx *gin.Context) {
	var input motorVehicleDto.UpdateMotorVehicle

	id := ctx.Param("id")

	ctx.ShouldBindJSON(&input)
	if err := utils.Validated(input); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad Request", "04", "01")
		return
	}

	motor, err := md.motorVehicleUC.UpdateMotorVehicle(id, input)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "04", "01")
		return
	}

	json.NewResponseSuccess(ctx, motor, "motor vehicle updated", "04", "01")
}

func (md motorVehicleDelivery) deleteMotorVehicle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := md.motorVehicleUC.DeleteMotorVehicle(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "05", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "Sucessfully deleted motor vehicle", "05", "01")
}
