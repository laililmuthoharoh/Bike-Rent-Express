package motorReturnDelivery

import (
	"bike-rent-express/model/dto/json"
	"bike-rent-express/model/dto/motorReturnDto"
	"bike-rent-express/pkg/utils"
	"bike-rent-express/src/motorReturn"

	"github.com/gin-gonic/gin"
)

type motorReturnDelivery struct {
	motorReturnUC motorReturn.MotorReturnUsecase
}

func NewMotorReturnDelivey(v1Group *gin.RouterGroup, motorReturnUC motorReturn.MotorReturnUsecase) {
	// handler := motorReturnDelivery{motorReturnUC}

	// motorReturnGroup := v1Group.Group("/employee/")
}

func (m *motorReturnDelivery) CreateMotorReturn(c *gin.Context) {
	var createMotorReturnRequest motorReturnDto.CreateMotorReturnRequest

	c.BindJSON(&createMotorReturnRequest)
	if err := utils.Validated(createMotorReturnRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "01", "01")
		return
	}

}
