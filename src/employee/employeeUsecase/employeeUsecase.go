package employeeUsecase

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/src/employee"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type employeeUsecase struct {
	employeeRepository employee.EmployeeRepository
}

func NewEmployeeUsecase(employeRepository employee.EmployeeRepository) employee.EmployeeUsecase {
	return &employeeUsecase{employeRepository}
}

func (e *employeeUsecase) AddEmployee(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)

	if err != nil {
		return employee, err
	}

	employee.Password = string(password)

	returnEmployee, err := e.employeeRepository.Add(employee)
	if err != nil {
		return returnEmployee, err
	}
	return returnEmployee, nil
}

func (e *employeeUsecase) Get() ([]employeeDto.Employee, error) {
	employee, err := e.employeeRepository.Get()

	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (e *employeeUsecase) GetById(id string) (employeeDto.Employee, error) {
	employee, err := e.employeeRepository.GetById(id)

	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return employee, errors.New("1")
		}
		return employee, err
	}

	return employee, nil
}

func (e *employeeUsecase) Update(employeUpdateRequest employeeDto.UpdateEmployeeRequest) (employeeDto.Employee, error) {
	employee, err := e.employeeRepository.Update(employeUpdateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return employee, errors.New("1")
		}
		return employee, err
	}

	return employee, err
}

func (e *employeeUsecase) Delete(id string) (string, error) {
	resultDelete, err := e.employeeRepository.Delete(id)
	if err != nil {
		return resultDelete, err
	}

	return resultDelete, nil
}
