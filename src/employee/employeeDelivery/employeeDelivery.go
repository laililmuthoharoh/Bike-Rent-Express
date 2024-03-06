package employeeDelivery

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/model/dto/json"
	"bike-rent-express/pkg/utils"
	"bike-rent-express/src/employee"
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type employeeDelivery struct {
	employeeUC employee.EmployeeUsecase
}

func NewEmployeeDelivery(v1Group *gin.RouterGroup, employeeUC employee.EmployeeUsecase) {
	handler := employeeDelivery{employeeUC}

	employeeGroup := v1Group.Group("employee")
	{
		employeeGroup.POST("/register", handler.AddEmployee)
		employeeGroup.GET("/:id", handler.GetEmployeById)
		employeeGroup.GET("", handler.GetEmployeeAll)
		employeeGroup.PUT("/:id", handler.UpdateEmployeeById)
		employeeGroup.DELETE("/:id", handler.DeleteEmployeeById)
	}
}

func (e *employeeDelivery) AddEmployee(c *gin.Context) {
	var addEmployeeRequest employeeDto.CreateEmployeeRequest

	c.BindJSON(&addEmployeeRequest)
	if err := utils.Validated(addEmployeeRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "01", "01")
		return
	}

	resultEmployee, err := e.employeeUC.Register(addEmployeeRequest)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(c, nil, "username already in use", "01", "01")
			return
		}
		json.NewResponseError(c, err.Error(), "01", "01")
		return
	}

	json.NewResponseCreated(c, resultEmployee, "Employee Created", "01", "01")
}

func (e *employeeDelivery) GetEmployeById(c *gin.Context) {
	id := c.Param("id")

	resultEmployee, err := e.employeeUC.GetById(id)
	if err != nil {
		if err.Error() == "1" || errors.Is(sql.ErrNoRows, err) {
			json.NewResponseSuccess(c, nil, "Data not found", "02", "01")
			return
		}
		json.NewResponseError(c, err.Error(), "02", "01")
		return
	}

	json.NewResponseSuccess(c, resultEmployee, "", "02", "02")
}

func (e *employeeDelivery) GetEmployeeAll(c *gin.Context) {
	resultEmployee, err := e.employeeUC.Get()
	if err != nil {
		json.NewResponseError(c, err.Error(), "03", "01")
		return
	}

	json.NewResponseSuccess(c, resultEmployee, "Success Get All Employee", "03", "01")
}

func (e *employeeDelivery) UpdateEmployeeById(c *gin.Context) {
	var employeUpdateRequest employeeDto.UpdateEmployeeRequest

	id := c.Param("id")
	employeUpdateRequest.ID = id

	c.BindJSON(&employeUpdateRequest)
	if err := utils.Validated(employeUpdateRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "04", "01")
		return
	}

	employee, err := e.employeeUC.Update(employeUpdateRequest)
	if err != nil {
		if err.Error() == "1" || errors.Is(sql.ErrNoRows, err) {
			json.NewResponseSuccess(c, nil, "Data not found", "04", "01")
			return
		}
		json.NewResponseError(c, err.Error(), "04", "01")
		return
	}

	json.NewResponseSuccess(c, employee, "Success update employee", "04", "01")
}

func (e *employeeDelivery) DeleteEmployeeById(c *gin.Context) {
	id := c.Param("id")

	msg, err := e.employeeUC.Delete(id)
	if err != nil {
		json.NewResponseError(c, err.Error(), "05", "01")
		return
	}

	json.NewResponseSuccess(c, nil, msg, "05", "01")
}
