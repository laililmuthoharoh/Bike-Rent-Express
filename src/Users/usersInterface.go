package Users

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto"
)

type UsersRepository interface {
	RegisterUsers(newUsers dto.RegisterUsers) error
	GetByID(id string) (dto.GetUsers, error)
	GetAll() ([]dto.GetUsers, error)
	UpdateUsers(usersItem dto.Users) error
	GetByUsername(username string) (dto.Users, error)
	UpdateBalance(topUpRequest dto.TopUpRequest) error
	UpdatePassword(changePasswordRequest dto.ChangePassword) error
	UsernameIsReady(username string) (bool, error)
}

type UsersUsecase interface {
	RegisterUsers(newUsers dto.RegisterUsers) error
	GetByID(id string) (dto.GetUsers, error)
	GetAllUsers() ([]dto.GetUsers, error)
	UpdateUsers(usersItem dto.Users) error
	LoginUsers(loginRequest model.LoginRequest) (dto.LoginResponse, error)
	TopUp(topUpRequest dto.TopUpRequest) error
	ChangePassword(changePasswordRequest dto.ChangePassword) error
}
