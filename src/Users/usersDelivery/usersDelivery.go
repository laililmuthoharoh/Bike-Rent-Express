package usersDelivery

import (
	"bike-rent-express/model/dto"
	"bike-rent-express/model/dto/json"
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
		usersGroup.GET("", handler.GetAllUsers)
		usersGroup.PUT("/update", handler.UpdateUsers)

		usersGroup.GET("/:id", handler.getByID)

		usersGroup.POST("/regist", handler.RegisterUsers)

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
	var newUsers dto.Users
	if err := ctx.ShouldBindJSON(newUsers); err != nil {
		json.NewResponseError(ctx, "Invalid request body", "02", "01")
		fmt.Println(err)
		return

	}

	// Call usecase to update the expense
	if err := h.usersUC.UpdateUsers(newUsers); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		fmt.Println(err)
		return

	}

	// Respond with success
	json.NewResponseSuccess(ctx, "", "Account Updated", "01", "01")
}

func (d *usersDelivery) getByID(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	usersItem, err := d.usersUC.GetByIDs(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, usersItem, "Success", "01", "01")
}

func (c *usersDelivery) RegisterUsers(ctx *gin.Context) {
	var newUsers dto.RegisterUsers
	if err := ctx.ShouldBindJSON(&newUsers); err != nil {
		json.NewResponseError(ctx, "Invalid request body", "02", "01")
		fmt.Println(err)
		return

	}

	// Call usecase to create the expense
	if err := c.usersUC.RegisterUsers(newUsers); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		fmt.Println(err)
		return

	}

	// Respond with success
	json.NewResponseSuccess(ctx, "", "Account Created", "01", "01")
}
