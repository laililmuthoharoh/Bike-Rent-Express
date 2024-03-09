package employeeRepository

import (
	employeeDto "bike-rent-express/model/dto/employee"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/assert"
)

var expectEmployee = employeeDto.Employee{
	ID:        "ec22df4c-1c1c-4012-9395-bc0994807e35",
	Name:      "dino",
	Telp:      "0812321412312",
	Username:  "dino123",
	Password:  "dino12345",
	CreatedAt: "2024-03-07T00:00:00Z",
	UpdatedAt: "2024-03-07T00:00:00Z",
}

func TestGetById_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB :", err.Error())
	}
	defer dbMock.Close()

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT (.+) FROM employee WHERE .+ = \\$1 AND deleted_at IS NULL;"

	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Telp, expectEmployee.Username, expectEmployee.Password, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	employee, err := employeeRepository.GetById(expectEmployee.ID)
	assert.Nil(t, err)
	assert.Equal(t, expectEmployee, employee)

}

func TestGetById_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB :", err.Error())
	}
	defer dbMock.Close()

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT (.+) FROM employee WHERE .+ = \\$1 AND deleted_at;"

	mock.ExpectQuery(query)

	employee, err := employeeRepository.GetById(expectEmployee.ID)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectEmployee, employee)
}

func TestGetByUsername_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}

	defer dbMock.Close()
	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT (.+) FROM employee WHERE .+ = \\$1 AND deleted_at IS NULL;"

	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Telp, expectEmployee.Username, expectEmployee.Password, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)
	mock.ExpectQuery(query).WillReturnRows(rows)

	employee, err := employeeRepository.GetByUsername(expectEmployee.ID)
	assert.Nil(t, err)
	assert.Equal(t, expectEmployee, employee)

}

func TestGetByUsername_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}

	defer dbMock.Close()
	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT (.+) FROM employee WHERE .+ = \\$1 AND deleted_at IS NULL;"
	mock.ExpectQuery(query)

	employee, err := employeeRepository.GetByUsername(expectEmployee.ID)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectEmployee, employee)
}

func TestUsernameIsReady_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	expectResult := false
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}

	defer dbMock.Close()
	employeeRepository := NewEmployeeRepository(dbMock)
	query := "SELECT COUNT(.+) FROM employee WHERE .+ = \\$1;"

	rows := sqlmock.NewRows([]string{".+"}).AddRow("1")
	mock.ExpectQuery(query).WillReturnRows(rows)

	resultActual, err := employeeRepository.UsernameIsReady(expectEmployee.ID)
	assert.Nil(t, err)
	assert.Equal(t, expectResult, resultActual)
}

func TestUsernameIsReady_False(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	expectActual := false

	defer dbMock.Close()
	employeeRepository := NewEmployeeRepository(dbMock)
	query := "SELECT COUNT(.+) FROM employee WHERE .+ = \\$1;"
	mock.ExpectQuery(query)

	resultActual, err := employeeRepository.UsernameIsReady(expectEmployee.Name)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, expectActual, resultActual)
}

func TestGet_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT (.+) FROM employee WHERE .+ IS NULL;"
	values := [][]driver.Value{
		{
			expectEmployee.ID,
			expectEmployee.Name,
			expectEmployee.Telp,
			expectEmployee.Username,
			expectEmployee.CreatedAt,
			expectEmployee.UpdatedAt,
		},
	}

	expectedAllEmployee := []employeeDto.Employee{
		{
			ID:        expectEmployee.ID,
			Name:      expectEmployee.Name,
			Telp:      expectEmployee.Telp,
			Username:  expectEmployee.Username,
			CreatedAt: expectEmployee.CreatedAt,
			UpdatedAt: expectEmployee.UpdatedAt,
		},
	}

	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+"}).AddRows(values...)
	mock.ExpectQuery(query).WillReturnRows(rows)

	actualAllEmployee, err := employeeRepository.Get()
	assert.Nil(t, err)
	assert.Equal(t, expectedAllEmployee, actualAllEmployee)
}

func TestGet_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT .+, .+, .+, .+, .+, .+ FROM employee WHERE .+ ;"
	values := [][]driver.Value{}

	expectedAllEmployee := []employeeDto.Employee{
		{
			ID:        expectEmployee.ID,
			Name:      expectEmployee.Name,
			Telp:      expectEmployee.Telp,
			Username:  expectEmployee.Username,
			CreatedAt: expectEmployee.CreatedAt,
			UpdatedAt: expectEmployee.UpdatedAt,
		},
	}

	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+"}).AddRows(values...)
	mock.ExpectQuery(query).WillReturnRows(rows)

	actualAllEmployee, err := employeeRepository.Get()
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectedAllEmployee, actualAllEmployee)
}

func TestAdd_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	expectCreatedEmployee := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT COUNT(.+) FROM employee WHERE .+ = \\$1;"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(0)

	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "INSERT INTO employee(.+) RETURNING .+;"
	rows = sqlmock.NewRows([]string{"id"}).AddRow(expectCreatedEmployee.ID)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualCreatedEmployee, err := employeeRepository.Add(expectCreatedEmployee)
	assert.Nil(t, err)
	assert.Equal(t, expectCreatedEmployee, actualCreatedEmployee)
}

func TestAddUsernameIsReady_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	expectCreatedEmployee := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT COUNT(.+) FROM employees WHERE .+ = \\$1;"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(0)

	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "INSERT INTO employee(.+) RETURNING .+;"
	rows = sqlmock.NewRows([]string{"id"}).AddRow(expectCreatedEmployee.ID)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualCreatedEmployee, err := employeeRepository.Add(expectCreatedEmployee)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, expectCreatedEmployee, actualCreatedEmployee)
}

func TestAddUsernameIsNotReady_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	expectCreatedEmployee := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT COUNT(.+) FROM employee WHERE .+ = \\$1;"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(1)

	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "INSERT INTO employee(.+) RETURNING .+;"
	rows = sqlmock.NewRows([]string{"id"}).AddRow(expectCreatedEmployee.ID)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualCreatedEmployee, err := employeeRepository.Add(expectCreatedEmployee)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, expectCreatedEmployee, actualCreatedEmployee)
}

func TestAdd_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	expectCreatedEmployee := employeeDto.CreateEmployeeRequest{
		ID:       expectEmployee.ID,
		Name:     expectEmployee.Name,
		Telp:     expectEmployee.Telp,
		Username: expectEmployee.Username,
		Password: expectEmployee.Password,
	}

	employeeRepository := NewEmployeeRepository(dbMock)

	query := "SELECT COUNT(.+) FROM employee WHERE .+ = \\$1;"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(0)

	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "INSERT INTO employeesss(.+) RETURNING .+;"
	rows = sqlmock.NewRows([]string{"id"}).AddRow(expectCreatedEmployee.ID)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualCreatedEmployee, err := employeeRepository.Add(expectCreatedEmployee)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Equal(t, expectCreatedEmployee, actualCreatedEmployee)
}

func TestUpdate_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()
	employeeRepo := NewEmployeeRepository(dbMock)
	expectEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
		ID:   expectEmployee.ID,
	}

	query := "UPDATE employee"
	mock.ExpectExec(query).WithArgs(expectEmployeeRequest.Name, expectEmployeeRequest.Telp, expectEmployeeRequest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	query = "SELECT (.+) FROM employee .+"
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Telp, expectEmployee.Username, expectEmployee.Password, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualEmployeeRequest, err := employeeRepo.Update(expectEmployeeRequest)
	assert.Nil(t, err)
	assert.Equal(t, expectEmployee, actualEmployeeRequest)
}

func TestUpdate_FailedQuery(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()
	employeeRepo := NewEmployeeRepository(dbMock)
	expectEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
		ID:   expectEmployee.ID,
	}

	query := "UPDATE employees"
	mock.ExpectExec(query).WithArgs(expectEmployeeRequest.Name, expectEmployeeRequest.Telp, expectEmployeeRequest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	query = "SELECT (.+) FROM employee .+"
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Telp, expectEmployee.Username, expectEmployee.Password, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualEmployeeRequest, err := employeeRepo.Update(expectEmployeeRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectEmployee, actualEmployeeRequest)
}

func TestUpdate_FailedGetById(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()
	employeeRepo := NewEmployeeRepository(dbMock)
	expectEmployeeRequest := employeeDto.UpdateEmployeeRequest{
		Name: expectEmployee.Name,
		Telp: expectEmployee.Telp,
		ID:   expectEmployee.ID,
	}

	query := "UPDATE employee"
	mock.ExpectExec(query).WithArgs(expectEmployeeRequest.Name, expectEmployeeRequest.Telp, expectEmployeeRequest.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	query = "SELECT (.+) FROM employees .+"
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Telp, expectEmployee.Username, expectEmployee.Password, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actualEmployeeRequest, err := employeeRepo.Update(expectEmployeeRequest)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectEmployee, actualEmployeeRequest)
}

func TestUpdatePassword_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	employeeRepo := NewEmployeeRepository(dbMock)

	query := "UPDATE employee"
	mock.ExpectExec(query).WithArgs(expectEmployee.Password, expectEmployee.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	err = employeeRepo.UpdatePassword(expectEmployee)
	assert.Nil(t, err)
}

func TestDeleteSuccess_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	employeeRepo := NewEmployeeRepository(dbMock)
	expectResult := "Sucessfully delete employee"

	query := "UPDATE employee"
	mock.ExpectExec(query).WithArgs(expectEmployee.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	actualResult, err := employeeRepo.Delete(expectEmployee.ID)

	assert.Nil(t, err)
	assert.Equal(t, expectResult, actualResult)

}

func TestDeleteSuccess_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	employeeRepo := NewEmployeeRepository(dbMock)
	expectResult := "Sucessfully delete employee"

	query := "UPDATE employees"
	mock.ExpectExec(query).WithArgs(expectEmployee.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	actualResult, err := employeeRepo.Delete(expectEmployee.ID)

	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.NotEqual(t, expectResult, actualResult)

}
