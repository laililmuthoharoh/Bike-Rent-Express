package usersRepository

import (
	"bike-rent-express/model/dto"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var expectUsers = dto.GetUsers{
	Uuid:       "1",
	Name:       "test",
	Username:   "test",
	Address:    "test",
	Role:       "USER",
	Can_rent:   true,
	Password:   "test",
	Created_at: "0000",
	Updated_at: "0000",
	Telp:       "0813123",
}

func TestGetById_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)

	query := "SELECT (.+) FROM users WHERE id = \\$1"
	row := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(expectUsers.Uuid, expectUsers.Name, expectUsers.Username, expectUsers.Password, expectUsers.Address, expectUsers.Role, expectUsers.Can_rent, expectUsers.Created_at, expectUsers.Updated_at, expectUsers.Telp)

	mock.ExpectQuery(query).WillReturnRows(row)
	actualUsers, err := userRepository.GetByID(expectUsers.Uuid)
	assert.Nil(t, err)
	assert.Equal(t, expectUsers, actualUsers)
}

func TestGetById_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)

	query := "SELECT (.+) FROM users WHERE id = \\$1"

	mock.ExpectQuery(query)
	actualUsers, err := userRepository.GetByID(expectUsers.Uuid)
	assert.NotEqual(t, expectUsers, actualUsers)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestGetAll_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)

	query := "SELECT (.+) FROM users WHERE role = 'USER'"
	values := [][]driver.Value{
		{
			expectUsers.Uuid,
			expectUsers.Name,
			expectUsers.Username,
			expectUsers.Address,
			expectUsers.Role,
			expectUsers.Can_rent,
			expectUsers.Created_at,
			expectUsers.Updated_at,
			expectUsers.Telp,
		},
	}
	expectAllUser := []dto.GetUsers{
		{
			Uuid:       expectUsers.Uuid,
			Name:       expectUsers.Name,
			Username:   expectUsers.Username,
			Address:    expectUsers.Address,
			Role:       expectUsers.Role,
			Can_rent:   expectUsers.Can_rent,
			Created_at: expectUsers.Created_at,
			Updated_at: expectUsers.Updated_at,
			Telp:       expectUsers.Telp,
		},
	}
	rows := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRows(values...)

	mock.ExpectQuery(query).WillReturnRows(rows)
	actualAllUser, err := userRepository.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expectAllUser, actualAllUser)
}

func TestGetAll_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)

	query := "SELECT (.+) FROM users WHERE role = 'USER'"
	expectAllUser := []dto.GetUsers{
		{
			Uuid:       expectUsers.Uuid,
			Name:       expectUsers.Name,
			Username:   expectUsers.Username,
			Address:    expectUsers.Address,
			Role:       expectUsers.Role,
			Can_rent:   expectUsers.Can_rent,
			Created_at: expectUsers.Created_at,
			Updated_at: expectUsers.Updated_at,
			Telp:       expectUsers.Telp,
		},
	}

	mock.ExpectQuery(query)
	actualAllUser, err := userRepository.GetAll()
	assert.NotNil(t, err)
	assert.Nil(t, actualAllUser)
	assert.NotEqual(t, expectAllUser, actualAllUser)
}

func TestUpdateUsers_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userItem := dto.Users{
		Name:    expectUsers.Name,
		Address: expectUsers.Address,
		CanRent: expectUsers.Can_rent,
		Telp:    expectUsers.Telp,
		ID:      expectUsers.Uuid,
	}
	userRepositroy := NewUsersRepository(dbMock)

	query := "UPDATE users"
	mock.ExpectExec(query).WithArgs(expectUsers.Name, expectUsers.Address, expectUsers.Can_rent, expectUsers.Telp, expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepositroy.UpdateUsers(userItem)
	assert.Nil(t, err)
}

func TestUpdateUsers_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userItem := dto.Users{
		Name:    expectUsers.Name,
		Address: expectUsers.Address,
		CanRent: expectUsers.Can_rent,
		Telp:    expectUsers.Telp,
		ID:      expectUsers.Uuid,
	}
	userRepositroy := NewUsersRepository(dbMock)

	query := "UPDATE usersw"
	mock.ExpectExec(query).WithArgs(expectUsers.Name, expectUsers.Address, expectUsers.Can_rent, expectUsers.Telp, expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))

	err = userRepositroy.UpdateUsers(userItem)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestUpdateUsers_FailedAffectedRow0(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	defer dbMock.Close()

	userItem := dto.Users{
		Name:    expectUsers.Name,
		Address: expectUsers.Address,
		CanRent: expectUsers.Can_rent,
		Telp:    expectUsers.Telp,
		ID:      expectUsers.Uuid,
	}
	userRepositroy := NewUsersRepository(dbMock)

	query := "UPDATE users"
	mock.ExpectExec(query).WithArgs(expectUsers.Name, expectUsers.Address, expectUsers.Can_rent, expectUsers.Telp, expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(1, 0))

	err = userRepositroy.UpdateUsers(userItem)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestRegisterUser_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()
	registerUser := dto.RegisterUsers{
		ID:       expectUsers.Uuid,
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Role,
	}

	userRepository := NewUsersRepository(dbMock)

	mock.ExpectBegin()
	query := "INSERT INTO users(.+) RETURNING .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(expectUsers.Uuid)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "INSERT INTO balance(.+)"
	mock.ExpectExec(query).WithArgs(expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = userRepository.RegisterUsers(registerUser)
	assert.Nil(t, err)
}

func TestRegisterUser_FailedInsertUser(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()
	registerUser := dto.RegisterUsers{
		ID:       expectUsers.Uuid,
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Role,
	}

	userRepository := NewUsersRepository(dbMock)

	mock.ExpectBegin()
	query := "INSERT INTO users(.+) RETURNING .+"
	mock.ExpectQuery(query)

	query = "INSERT INTO balance(.+)"
	mock.ExpectExec(query).WithArgs(expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = userRepository.RegisterUsers(registerUser)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestRegisterUser_FailedInsertBalance(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()
	registerUser := dto.RegisterUsers{
		ID:       expectUsers.Uuid,
		Name:     expectUsers.Name,
		Username: expectUsers.Username,
		Password: expectUsers.Password,
		Address:  expectUsers.Address,
		Role:     expectUsers.Role,
		Telp:     expectUsers.Role,
	}

	userRepository := NewUsersRepository(dbMock)

	mock.ExpectBegin()
	query := "INSERT INTO users(.+) RETURNING .+"
	rows := sqlmock.NewRows([]string{".+"}).AddRow(expectUsers.Uuid)
	mock.ExpectQuery(query).WillReturnRows(rows)

	query = "INSERT INTO balances(.+)"
	mock.ExpectExec(query).WithArgs(expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	err = userRepository.RegisterUsers(registerUser)
	assert.NotNil(t, err)
	assert.Error(t, err)
}

func TestGetByUsername_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	user := dto.Users{
		ID:         expectUsers.Uuid,
		Name:       expectUsers.Name,
		Username:   expectUsers.Username,
		Password:   expectUsers.Password,
		Address:    expectUsers.Address,
		Role:       expectUsers.Role,
		CanRent:    expectUsers.Can_rent,
		Updated_at: expectUsers.Telp,
		Telp:       expectUsers.Telp,
	}

	userRepository := NewUsersRepository(dbMock)

	query := "SELECT (.+) FROM users WHERE .+"
	row := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(user.ID, user.Name, user.Username, user.Password, user.Address, user.Role, user.CanRent, user.Updated_at, user.Telp)
	mock.ExpectQuery(query).WillReturnRows(row)

	actualUser, err := userRepository.GetByUsername(user.Username)
	assert.Nil(t, err)
	assert.Equal(t, user, actualUser)
}

func TestGetByUsername_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error DB:", err.Error())
	}
	user := dto.Users{
		ID:         expectUsers.Uuid,
		Name:       expectUsers.Name,
		Username:   expectUsers.Username,
		Password:   expectUsers.Password,
		Address:    expectUsers.Address,
		Role:       expectUsers.Role,
		CanRent:    expectUsers.Can_rent,
		Updated_at: expectUsers.Telp,
		Telp:       expectUsers.Telp,
	}

	userRepository := NewUsersRepository(dbMock)

	query := "SELECT (.+) FROM userss WHERE .+"
	row := sqlmock.NewRows([]string{".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+", ".+"}).AddRow(user.ID, user.Name, user.Username, user.Password, user.Address, user.Role, user.CanRent, user.Updated_at, user.Telp)
	mock.ExpectQuery(query).WillReturnRows(row)

	actualUser, err := userRepository.GetByUsername(user.Username)
	assert.NotNil(t, err)
	assert.NotEqual(t, user, actualUser)
}

func TestUpdateBalance_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)
	topUpRequest := dto.TopUpRequest{
		Amount: 0,
		UserID: expectUsers.Uuid,
	}

	query := "UPDATE balance"
	mock.ExpectExec(query).WithArgs(0, expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepository.UpdateBalance(topUpRequest)
	assert.Nil(t, err)
}

func TestUpdateBalance_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)
	topUpRequest := dto.TopUpRequest{
		Amount: 0,
		UserID: expectUsers.Uuid,
	}

	query := "UPDATE balance"
	mock.ExpectExec(query).WithArgs(0, expectUsers.Uuid)

	err = userRepository.UpdateBalance(topUpRequest)
	assert.NotNil(t, err)
}

func TestUpdatePassword_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: expectUsers.Password,
		NewPassword: expectUsers.Password,
	}

	query := "UPDATE users"
	mock.ExpectExec(query).WithArgs(expectUsers.Password, expectUsers.Uuid).WillReturnResult(sqlmock.NewResult(1, 1))

	err = userRepository.UpdatePassword(changePassword)
	assert.Nil(t, err)
}

func TestUpdatePassword_Failed(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)
	changePassword := dto.ChangePassword{
		ID:          expectUsers.Uuid,
		OldPassword: expectUsers.Password,
		NewPassword: expectUsers.Password,
	}

	query := "UPDATE users"
	mock.ExpectExec(query).WithArgs(expectUsers.Password, expectUsers.Uuid)

	err = userRepository.UpdatePassword(changePassword)
	assert.NotNil(t, err)
}

func TestUsernameIsReady_Success(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error DB:", err.Error())
	}
	defer dbMock.Close()

	userRepository := NewUsersRepository(dbMock)
	expectUsernameIsReady := true
	query := "SELECT COUNT(.+) FROM users WHERE .+"
	row := sqlmock.NewRows([]string{".+"}).AddRow(0)
	mock.ExpectQuery(query).WillReturnRows(row)

	actualUsernameIsReady, err := userRepository.UsernameIsReady(expectUsers.Username)
	assert.Nil(t, err)
	assert.Equal(t, expectUsernameIsReady, actualUsernameIsReady)
}
