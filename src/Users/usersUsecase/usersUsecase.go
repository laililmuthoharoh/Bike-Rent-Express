package usersUsecase

import (
	"bike-rent-express/model"
	"bike-rent-express/model/dto"
	"bike-rent-express/pkg/middleware"
	"bike-rent-express/src/Users"
	"database/sql"
	"errors"
	"strings"

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

	return uc.usersRepo.UpdateUsers(updateUsers)
}

func (uc *usersUC) GetByID(id string) (dto.GetUsers, error) {
	user, err := uc.usersRepo.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return user, errors.New("1")
		}
		return user, err
	}

	return user, nil
}

func NewUsersUsecase(usersRepo Users.UsersRepository) Users.UsersUsecase {
	return &usersUC{usersRepo}
}

func (c *usersUC) RegisterUsers(newUsers dto.RegisterUsers) error {
	usernameReady, err := c.usersRepo.UsernameIsReady(newUsers.Username)

	if err != nil {
		return err
	}

	if !usernameReady {
		return errors.New("1")
	}

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
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return loginResponse, errors.New("1")
		}
		return loginResponse, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return loginResponse, errors.New("1")
	}
	token, err := middleware.GenerateTokenJwt(user.Username, user.Role)
	if err != nil {
		return loginResponse, err
	}
	loginResponse.User = user
	loginResponse.AccesToken = token

	return loginResponse, nil
}

func (c *usersUC) TopUp(topUpRequest dto.TopUpRequest) error {
	balance, err := c.usersRepo.GetBalance(topUpRequest.UserID)
	if err != nil {
		if err == sql.ErrNoRows || strings.Contains(err.Error(), "invalid input syntax for type uuid") {
			return errors.New("1")
		}
		return err
	}
	topUpRequest.Amount += balance.Amount
	err = c.usersRepo.UpdateBalance(topUpRequest)
	return err
}

func (c *usersUC) ChangePassword(changePasswordRequest dto.ChangePassword) error {
	user, err := c.usersRepo.GetByID(changePasswordRequest.ID)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return errors.New("1")
		}
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(changePasswordRequest.OldPassword))
	if err != nil {
		return errors.New("2")
	}
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(changePasswordRequest.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("2")
	}

	changePasswordRequest.NewPassword = string(encryptPass)

	err = c.usersRepo.UpdatePassword(changePasswordRequest)
	return err
}

func (c *usersUC) GetBalanceCustomer(id string) (dto.Balance, error) {
	balance, err := c.usersRepo.GetBalance(id)
	if err != nil {
		if strings.Contains(err.Error(), "invalid input syntax for type uuid") || err == sql.ErrNoRows {
			return balance, errors.New("1")
		}
	}
	return balance, err
}
