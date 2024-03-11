package usersRepository

import (
	"bike-rent-express/model/dto"
	"bike-rent-express/src/Users"
	"database/sql"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) Users.UsersRepository {
	return &usersRepository{db}
}

func (r *usersRepository) GetByID(uuid string) (dto.GetUsers, error) {
	query := `SELECT id, name, username, password,address, role, can_rent,created_at, updated_at,telp FROM users WHERE id = $1`
	var usersItem dto.GetUsers
	if err := r.db.QueryRow(query, uuid).Scan(
		&usersItem.Uuid,
		&usersItem.Name,
		&usersItem.Username,
		&usersItem.Password,
		&usersItem.Address,
		&usersItem.Role,
		&usersItem.Can_rent,
		&usersItem.Created_at,
		&usersItem.Updated_at,
		&usersItem.Telp,
	); err != nil {
		return usersItem, err
	}

	return usersItem, nil
}

func (r *usersRepository) GetAll() ([]dto.GetUsers, error) {
	query := `SELECT id, name, username, address, role, can_rent, created_at, updated_at, telp FROM users WHERE role = 'USER'`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []dto.GetUsers
	for rows.Next() {
		var user dto.GetUsers
		if err := rows.Scan(
			&user.Uuid,
			&user.Name,
			&user.Username,
			&user.Address,
			&user.Role,
			&user.Can_rent,
			&user.Created_at,
			&user.Updated_at,
			&user.Telp,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *usersRepository) UpdateUsers(usersUpdate dto.Users) error {
	query := `
        UPDATE users
        SET name = $1, 
            address = $2, 
            can_rent = $3, 
            telp = $4,
			updated_at = CURRENT_TIMESTAMP
        WHERE id = $5
    `
	result, err := r.db.Exec(query,
		usersUpdate.Name,
		usersUpdate.Address,
		usersUpdate.CanRent,
		usersUpdate.Telp,
		usersUpdate.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (c *usersRepository) RegisterUsers(newUsers dto.RegisterUsers) error {
	tx, err := c.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	canRent := newUsers.Role == "USER"

	query := `INSERT INTO users (name, username, password, address, role, can_rent, telp) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id;`
	if err := tx.QueryRow(query,
		newUsers.Name,
		newUsers.Username,
		newUsers.Password,
		newUsers.Address,
		newUsers.Role,
		canRent,
		newUsers.Telp).Scan(&newUsers.ID); err != nil {
		tx.Rollback()
		return err
	}

	if newUsers.Role == "USER" {
		query = "INSERT INTO balance (amount, user_id) VALUES(0, $1);"
		_, err = tx.Exec(query, newUsers.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()

	return nil
}

func (c *usersRepository) GetByUsername(username string) (dto.Users, error) {
	var user dto.Users
	query := `SELECT id, name, username, password, address, role, can_rent,updated_at,telp FROM users WHERE username = $1`
	if err := c.db.QueryRow(query, username).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Address, &user.Role, &user.CanRent, &user.Updated_at, &user.Telp); err != nil {
		return user, err
	}
	return user, nil
}

func (c *usersRepository) UpdateBalance(topUpRequest dto.TopUpRequest) error {
	query := "UPDATE balance SET amount =$1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $2;"
	_, err := c.db.Exec(query, topUpRequest.Amount, topUpRequest.UserID)
	return err
}

func (c *usersRepository) UpdatePassword(changePasswordRequest dto.ChangePassword) error {
	query := "UPDATE users SET password = $1 WHERE id = $2"
	_, err := c.db.Exec(query, changePasswordRequest.NewPassword, changePasswordRequest.ID)
	return err
}

func (c *usersRepository) UsernameIsReady(username string) (bool, error) {
	query := "SELECT COUNT(username) FROM users WHERE username = $1;"
	var result int

	err := c.db.QueryRow(query, username).Scan(&result)
	return result == 0, err
}
