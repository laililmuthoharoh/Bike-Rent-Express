package motorVehicleDelivery

import (
	"bike-rent-express/model/dto/json"
	"bike-rent-express/model/dto/motorVehicleDto"
	"bike-rent-express/pkg/middleware"
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
	motorVehicleGroup := v1Group.Group("/motor-vehicles")
	{
		motorVehicleGroup.GET("/", middleware.JWTAuth("ADMIN", "USER"), handler.getAllMotorVehicle)
		motorVehicleGroup.GET("/:id", middleware.JWTAuth("ADMIN", "USER"), handler.getMotorVehicleById)
		motorVehicleGroup.POST("/", middleware.JWTAuth("ADMIN"), handler.createMotorVehicle)
		motorVehicleGroup.PUT("/:id", middleware.JWTAuth("ADMIN"), handler.updateMotorVehicle)
		motorVehicleGroup.DELETE("/:id", middleware.JWTAuth("ADMIN"), handler.deleteMotorVehicle)
	}
}

func (md motorVehicleDelivery) getAllMotorVehicle(ctx *gin.Context) {

	motor, err := md.motorVehicleUC.GetAllMotorVehicle()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	if len(motor) == 0 {
		json.NewResponseSuccess(ctx, nil, "Empty data", "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, motor, "success", "01", "02")
}

func (md motorVehicleDelivery) getMotorVehicleById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := md.motorVehicleUC.GetMotorVehicleById(id)
	if err != nil {
		if err.Error() == "1" || errors.Is(sql.ErrNoRows, err) {
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
		if err.Error() == "1" {
			json.NewResponseBadRequest(ctx, nil, "the license plate is already registered to another vehicle", "03", "01")
			return
		}
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
		if err.Error() == "1" {
			json.NewResponseBadRequest(ctx, nil, "the license plate is already registered to another vehicle", "03", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "04", "01")
		return
	}

	json.NewResponseSuccess(ctx, motor, "motor vehicle updated", "04", "01")
}

func (md motorVehicleDelivery) deleteMotorVehicle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := md.motorVehicleUC.DeleteMotorVehicle(id)
	if err != nil {
		if err == sql.ErrNoRows {
			json.NewResponseBadRequest(ctx, nil, "Data not found", "05", "01")
			return
		}
		if err.Error() == "1" {
			json.NewResponseBadRequest(ctx, nil, "the motor cannot be deleted because the motor still has a not available status.", "05", "02")
			return
		}
		json.NewResponseError(ctx, err.Error(), "05", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "Sucessfully deleted motor vehicle", "05", "01")
}
