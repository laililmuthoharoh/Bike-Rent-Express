package employee

import employeeDto "bike-rent-express/model/dto/employee"

type (
	EmployeeRepository interface {
		Add(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error)
		Get() ([]employeeDto.Employee, error)
		UsernameIsReady(username string) (bool, error)
		GetByUsername(username string) (employeeDto.Employee, error)
		GetById(id string) (employeeDto.Employee, error)
		Update(employeeUpdateRequest employeeDto.UpdateEmployeeRequest) (employeeDto.Employee, error)
		UpdatePassword(employee employeeDto.Employee) error
		Delete(id string) (string, error)
	}
	EmployeeUsecase interface {
		Register(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error)
		Get() ([]employeeDto.Employee, error)
		GetById(id string) (employeeDto.Employee, error)
		Update(employeUpdateRequest employeeDto.UpdateEmployeeRequest) (employeeDto.Employee, error)
		Delete(id string) (string, error)
		Login(loginRequest employeeDto.LoginRequest) (employeeDto.LoginResponse, error)
		ChangePassword(id string, changePasswordRequest employeeDto.ChangePasswordRequest) error
	}
)
