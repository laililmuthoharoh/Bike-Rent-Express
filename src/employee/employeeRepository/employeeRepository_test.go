package employeeRepository

import (
	employeeDto "bike-rent-express/model/dto/employee"
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

	query := "SELECT .+, .+, .+, .+, .+, .+, .+ FROM employee WHERE .+ = \\$1 AND deleted_at IS NULL;"

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

	query := "SELECT .+,  .+, .+, .+, .+, .+, .+ FROM employee WHERE .+ = \\$1 AND deleted_at;"

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

	query := "SELECT .+, .+, .+, .+, .+, .+, .+ FROM employee WHERE .+ = \\$1 AND deleted_at IS NULL;"

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

	query := "SELECT .+, .+, .+, .+, .+, .+, .+ FROM employee WHERE .+ = \\$1 AND deleted_at IS NULL;"
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

// func TestAdd_Success(t *testing.T){
// 	dbMock, mock, err := sqlmock.New()
// 	if err != nil{
// 		t.Fatal("error DB:", err.Error())
// 	}

// 	defer dbMock.Close()
// 	employeeRepository := NewEmployeeRepository(dbMock)

// }
