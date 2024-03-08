package usersDelivery

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto"
	"bike-rent-express/model/dto/json"
	"bike-rent-express/pkg/middleware"
	"bike-rent-express/pkg/utils"
	"bike-rent-express/src/Users"
	"fmt"

	"github.com/gin-gonic/gin"
)

type usersDelivery struct {
	usersUC Users.UsersUsecase
}

func NewUsersDelivery(v1Group *gin.RouterGroup, usersUC Users.UsersUsecase) {
	handler := usersDelivery{
		usersUC: usersUC,
	}
	usersGroup := v1Group.Group("/users")
	{
		usersGroup.GET("", middleware.JWTAuth("ADMIN"), handler.GetAllUsers)
		usersGroup.PUT("/:id", middleware.JWTAuth("ADMIN", "USER"), handler.UpdateUsers)

		usersGroup.GET("/:id", middleware.JWTAuth("ADMIN", "USER"), handler.getByID)
		usersGroup.PUT("/:id/change-password", middleware.JWTAuth("ADMIN", "USER"), handler.ChangePassword)
		usersGroup.PUT("/:id/top-up", middleware.JWTAuth("USER"), handler.TopUp)

		usersGroup.POST("/register", handler.RegisterUsers)
		usersGroup.POST("/login", handler.LoginUsers)

	}
}

func (h *usersDelivery) GetAllUsers(ctx *gin.Context) {
	users, err := h.usersUC.GetAllUsers()
	if err != nil {
		// Handle error response
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, users, "Success", "01", "01")
}

func (h *usersDelivery) UpdateUsers(ctx *gin.Context) {
	id := ctx.Param("id")

	var newUsers dto.Users
	newUsers.ID = id

	ctx.ShouldBindJSON(&newUsers)
	if err := utils.Validated(newUsers); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad Request", "02", "01")
		return
	}

	if err := h.usersUC.UpdateUsers(newUsers); err != nil {
		json.NewResponseError(ctx, err.Error(), "02", "01")
		fmt.Println(err)
		return

	}

	// Respond with success
	json.NewResponseSuccess(ctx, nil, "Account Updated", "02", "01")
}

func (d *usersDelivery) getByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	usersItem, err := d.usersUC.GetByID(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "03", "01")
		return
	}

	json.NewResponseSuccess(ctx, usersItem, "Success", "03", "01")
}

func (c *usersDelivery) RegisterUsers(ctx *gin.Context) {
	var newUsers dto.RegisterUsers

	ctx.ShouldBindJSON(&newUsers)
	if err := utils.Validated(newUsers); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad Request", "04", "01")
		return
	}

	if err := c.usersUC.RegisterUsers(newUsers); err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(ctx, nil, "username already in use", "04", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "04", "02")
		fmt.Println(err)
		return

	}

	// Respond with success
	json.NewResponseSuccess(ctx, newUsers, "Account Created", "04", "02")
}

func (c *usersDelivery) LoginUsers(ctx *gin.Context) {
	var loginRequest model.LoginRequest
	ctx.BindJSON(&loginRequest)
	if err := utils.Validated(loginRequest); err != nil {
		json.NewResponseBadRequest(ctx, err, "bad request", "05", "01")
		return
	}

	loginResponse, err := c.usersUC.LoginUsers(loginRequest)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(ctx, nil, "Incorrect username or password", "05", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "05", "01")
		return
	}

	json.NewResponseSuccess(ctx, loginResponse, "login succes", "05", "02")

}

func (c *usersDelivery) TopUp(ctx *gin.Context) {
	id := ctx.Param("id")
	var topupRequest dto.TopUpRequest

	ctx.BindJSON(&topupRequest)
	if err := utils.Validated(topupRequest); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad Request", "06", "01")
		return
	}

	topupRequest.UserID = id
	err := c.usersUC.TopUp(topupRequest)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "06", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "Success Top Up", "06", "01")

}

func (c *usersDelivery) ChangePassword(ctx *gin.Context) {
	id := ctx.Param("id")
	var changePasswordRequest dto.ChangePassword

	ctx.BindJSON(&changePasswordRequest)
	if err := utils.Validated(changePasswordRequest); err != nil {
		json.NewResponseBadRequest(ctx, err, "Bad Request", "07", "01")
		return
	}

	changePasswordRequest.ID = id
	err := c.usersUC.ChangePassword(changePasswordRequest)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(ctx, nil, "Data not found", "07", "01")
			return
		}
		if err.Error() == "2" {
			json.NewResponseSuccess(ctx, nil, "password does not match", "07", "02")
			return
		}
		json.NewResponseError(ctx, err.Error(), "07", "02")
		return
	}

	json.NewResponseSuccess(ctx, nil, "Success change password", "07", "03")
}
