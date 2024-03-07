package employeeRepository

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"bike-rent-express/src/employee"
	"database/sql"
	"errors"
	"time"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) employee.EmployeeRepository {
	return &employeeRepository{db}
}

func (e *employeeRepository) Add(employee employeeDto.CreateEmployeeRequest) (employeeDto.CreateEmployeeRequest, error) {
	usernameReady, err := e.UsernameIsReady(employee.Username)
	if err != nil {
		return employee, err
	}

	if !usernameReady {
		return employee, errors.New("1")
	}

	query := "INSERT INTO employee(name, telp, username, password) VALUES($1,$2,$3,$4) RETURNING id;"

	if err := e.db.QueryRow(query, employee.Name, employee.Telp, employee.Username, employee.Password).Scan(&employee.ID); err != nil {
		return employee, err
	}

	return employee, nil
}

func (e *employeeRepository) GetByUsername(username string) (employeeDto.Employee, error) {
	var employee employeeDto.Employee
	query := "SELECT SELECT id, name, telp, username, password FROM employee WHERE username = $1 AND deleted_at IS NULL;"

	if err := e.db.QueryRow(query, username).Scan(&employee.ID, &employee.Name, &employee.Telp, &employee.Password); err != nil {
		return employee, err
	}

	return employee, nil
}

func (e *employeeRepository) UsernameIsReady(username string) (bool, error) {
	query := "SELECT COUNT(username) FROM employee WHERE username = $1"

	var result int
	if err := e.db.QueryRow(query, username).Scan(&result); err != nil {
		return false, err
	}

	return result == 0, nil
}

func (e *employeeRepository) Get() ([]employeeDto.Employee, error) {
	query := "SELECT id, name, telp, username, created_at, updated_at FROM employee WHERE deleted_at IS NULL;"

	var employee []employeeDto.Employee
	row, err := e.db.Query(query)
	if err != nil {
		return employee, err
	}

	for row.Next() {
		var employe employeeDto.Employee
		err := row.Scan(&employe.ID, &employe.Name, &employe.Telp, &employe.Username, &employe.CreatedAt, &employe.UpdatedAt)
		if err != nil {
			return employee, err
		}
		employee = append(employee, employe)
	}

	return employee, nil
}

func (e *employeeRepository) GetById(id string) (employeeDto.Employee, error) {

	var employee employeeDto.Employee
	query := "SELECT id, name, telp, username, created_at, updated_at FROM employee WHERE id = $1 AND deleted_at IS NULL;"
	if err := e.db.QueryRow(query, id).Scan(&employee.ID, &employee.Name, &employee.Telp, &employee.Username, &employee.CreatedAt, &employee.UpdatedAt); err != nil {
		return employee, err
	}

	return employee, nil
}

func (e *employeeRepository) Update(employeeUpdateRequest employeeDto.UpdateEmployeeRequest) (employeeDto.Employee, error) {
	query := "UPDATE employee SET name = $1, telp = $2, password=$3, updated_at= $4 WHERE id = $5 AND deleted_at IS NULL;"
	now := time.Now()
	_, err := e.db.Exec(query, employeeUpdateRequest.Name, employeeUpdateRequest.Telp, employeeUpdateRequest.Password, now, employeeUpdateRequest.ID)
	if err != nil {
		return employeeDto.Employee{}, err
	}

	employee, err := e.GetById(employeeUpdateRequest.ID)
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (e *employeeRepository) Delete(id string) (string, error) {
	now := time.Now()

	query := "UPDATE employee SET deleted_at = $1 WHERE id = $2 AND deleted_at IS NULL;"
	_, err := e.db.Exec(query, now, id)
	if err != nil {
		return "", err
	}

	return "Sucessfully delete employee", nil
}
