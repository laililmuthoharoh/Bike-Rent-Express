package usersUsecase

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto"
	"bike-rent-express/pkg/middleware"
	"bike-rent-express/src/Users"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type usersUC struct {
	usersRepo Users.UsersRepository
}

func (uc *usersUC) GetAllUsers() ([]dto.GetUsers, error) {
	// Call the repository method to fetch all users
	users, err := uc.usersRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *usersUC) UpdateUsers(updateUsers dto.Users) error {
	if updateUsers.Name == "" {
		return errors.New("transaction Type cannot be empty")
	}
	if updateUsers.Username == "" {
		return errors.New("transaction Type cannot be empty")
	}

	if updateUsers.Password == "" {
		return errors.New("description cannot be empty")
	}
	if updateUsers.Address == "" {
		return errors.New("transaction Type cannot be empty")
	}
	if updateUsers.Role == "" {
		return errors.New("transaction Type cannot be empty")
	}
	if updateUsers.Can_rent == "" {
		return errors.New("transaction Type cannot be empty")
	}
	if updateUsers.Telp == "" {
		return errors.New("transaction Type cannot be empty")
	}
	updateUsers.Updated_at = time.Now().Format("2006-01-02")

	fmt.Println(updateUsers)

	return uc.usersRepo.UpdateUsers(updateUsers)
}

func (uc *usersUC) GetByIDs(id string) (dto.GetUsers, error) {
	return uc.usersRepo.GetByID(id)
}

func NewUsersUsecase(usersRepo Users.UsersRepository) Users.UsersUsecase {
	return &usersUC{usersRepo}
}

func (c *usersUC) RegisterUsers(newUsers dto.RegisterUsers) error {

	if newUsers.Name == "" {
		return errors.New("name Type cannot be empty")
	}
	if newUsers.Username == "" {
		return errors.New("username Type cannot be empty")
	}

	if newUsers.Password == "" {
		return errors.New("password cannot be empty")
	}
	if newUsers.Address == "" {
		return errors.New("address Type cannot be empty")
	}
	if newUsers.Role == "" {
		return errors.New("role Type cannot be empty")
	}
	if newUsers.Can_rent == "" {
		return errors.New("can rent Type cannot be empty")
	}
	if newUsers.Telp == "" {
		return errors.New("phone Type cannot be empty")
	}
	newUsers.Created_at = time.Now().Format("2006-01-02")

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(newUsers.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUsers.Password = string(encryptPassword)

	return c.usersRepo.RegisterUsers(newUsers)
}

func (c *usersUC) LoginUsers(loginRequest model.LoginRequest) (dto.LoginResponse, error) {
	var loginResponse dto.LoginResponse
	user, err := c.usersRepo.GetByUsername(loginRequest.Username)
	if err != nil {
		return loginResponse, err
	}

	fmt.Println(user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return loginResponse, err
	}

	token, err := middleware.GenerateTokenJwt(user.Username, user.Role)
	if err != nil {
		return loginResponse, err
	}
	loginResponse.User = user
	loginResponse.AccesToken = token

	return loginResponse, nil
}
