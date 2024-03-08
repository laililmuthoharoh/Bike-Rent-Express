package employeeDelivery

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/model/dto/json"
	"bike-rent-express/pkg/middleware"
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
		employeeGroup.POST("/login", handler.LoginEmployee)
		employeeGroup.PUT("/:id/change-password", middleware.JWTAuth("ADMIN", "EMPLOYEE"), handler.ChangePassword)
		employeeGroup.GET("/:id", middleware.JWTAuth("ADMIN", "EMPLOYEE"), handler.GetEmployeById)
		employeeGroup.GET("", middleware.JWTAuth("ADMIN"), handler.GetEmployeeAll)
		employeeGroup.PUT("/:id", middleware.JWTAuth("ADMIN", "EMPLOYEE"), handler.UpdateEmployeeById)
		employeeGroup.DELETE("/:id", middleware.JWTAuth("ADMIN"), handler.DeleteEmployeeById)
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

func (e *employeeDelivery) LoginEmployee(c *gin.Context) {
	var loginRequest employeeDto.LoginRequest

	c.BindJSON(&loginRequest)

	if err := utils.Validated(loginRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "06", "01")
		return
	}

	loginResponse, err := e.employeeUC.Login(loginRequest)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(c, nil, "Incorrect username or password", "06", "01")
			return
		} else if err.Error() == "2" {
			json.NewResponseSuccess(c, nil, "Incorrect username or password", "06", "02")
			return
		}

		json.NewResponseError(c, err.Error(), "06", "01")
		return
	}

	json.NewResponseSuccess(c, loginResponse, "Login successfully", "06", "03")
}

func (e *employeeDelivery) ChangePassword(c *gin.Context) {
	var changePasswordRequest employeeDto.ChangePasswordRequest

	c.BindJSON(&changePasswordRequest)
	id := c.Param("id")

	if err := utils.Validated(changePasswordRequest); err != nil {
		json.NewResponseBadRequest(c, err, "Bad Request", "07", "01")
		return
	}

	err := e.employeeUC.ChangePassword(id, changePasswordRequest)
	if err != nil {
		if err.Error() == "1" {
			json.NewResponseSuccess(c, nil, "Incorrect username or password", "07", "01")
			return
		} else if err.Error() == "2" {
			json.NewResponseSuccess(c, nil, "Incorrect username or password", "07", "02")
			return
		}
		json.NewResponseError(c, err.Error(), "07", "01")
		return
	}

	json.NewResponseSuccess(c, nil, "Success updated password", "07", "03")
}
