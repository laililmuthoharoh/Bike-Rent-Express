package employeeUsecase

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/pkg/middleware"
	"bike-rent-express/src/employee"
	"database/sql"
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

func (e *employeeUsecase) Register(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error) {
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
	_, err := e.employeeRepository.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return "", errors.New("1")
		}
		return "", err
	}

	resultDelete, err := e.employeeRepository.Delete(id)
	if err != nil {
		return resultDelete, err
	}

	return resultDelete, nil
}

func (e *employeeUsecase) Login(loginRequest employeeDto.LoginRequest) (employeeDto.LoginResponse, error) {
	employee, err := e.employeeRepository.GetByUsername(loginRequest.Username)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "invalid input syntax for type uuid" {
			return employeeDto.LoginResponse{}, errors.New("1")
		}
		return employeeDto.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(loginRequest.Password))
	if err != nil {
		return employeeDto.LoginResponse{}, errors.New("2")
	}

	token, err := middleware.GenerateTokenJwt(employee.Username, "EMPLOYEE")
	if err != nil {
		return employeeDto.LoginResponse{}, err
	}

	loginResponse := employeeDto.LoginResponse{
		AccessToken: token,
		Employee:    employee,
	}

	return loginResponse, nil
}

func (e *employeeUsecase) ChangePassword(id string, changePasswordRequest employeeDto.ChangePasswordRequest) error {
	employee, err := e.employeeRepository.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return errors.New("1")
		}
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(changePasswordRequest.PasswordOld))
	if err != nil {
		return errors.New("2")
	}

	encryptPass, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	employee.Password = string(encryptPass)
	err = e.employeeRepository.UpdatePassword(employee)

	return err
}
