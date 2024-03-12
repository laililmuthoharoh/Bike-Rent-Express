package motorReturnDelivery

import (
	"bike-rent-express/model/dto/json"
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/pkg/middleware"
	"bike-rent-express/pkg/utils"
	"bike-rent-express/src/motorReturn"

	"github.com/gin-gonic/gin"
)

type motorReturnDelivery struct {
	motorReturnUC motorReturn.MotorReturnUsecase
}

func NewMotorReturnDelivey(v1Group *gin.RouterGroup, motorReturnUC motorReturn.MotorReturnUsecase) {
	handler := motorReturnDelivery{motorReturnUC}

	motorReturnGroup := v1Group.Group("employee/:id/motor-return")
	{
		motorReturnGroup.POST("", middleware.JWTAuth("EMPLOYEE"), handler.CreateMotorReturn)
		motorReturnGroup.GET("/:motor-return-id", middleware.JWTAuth("EMPLOYEE", "ADMIN"), handler.GetMotorReturnById)
	}

	v1Group.GET("/users/motor-return", middleware.JWTAuth("ADMIN", "EMPLOYEE"), handler.GetAllMotorReturn)

}

func (m *motorReturnDelivery) CreateMotorReturn(c *gin.Context) {
	var createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest

	c.BindJSON(&createMotorReturnRequest)
	if err := utils.Validated(createMotorReturnRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "01", "01")
		return
	}

	motorReturnCreated, err := m.motorReturnUC.AddMotorReturn(createMotorReturnRequest)
	if err != nil {
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	json.NewResponseCreated(c, motorReturnCreated, "Motor return created", "01", "01")
}

func (m *motorReturnDelivery) GetMotorReturnById(c *gin.Context) {
	id := c.Param("motor-return-id")
	motorReturnDetail, err := m.motorReturnUC.GetMotorReturnById(id)
	if err != nil {
		json.NewResponseError(c, err.Error(), "02", "01")
		return
	}

	json.NewResponseSuccess(c, motorReturnDetail, "Success get motor return by id", "02", "01")
}

func (m *motorReturnDelivery) GetAllMotorReturn(c *gin.Context) {

	motorsReturn, err := m.motorReturnUC.GetMotorReturnAll()
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "01")
		return
	}

	json.NewResponseSuccess(c, motorsReturn, "Success get all motor return", "03", "01")
}
