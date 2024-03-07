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

func (r *usersRepository) GetByID(uuid string) (*dto.GetUsers, error) {
	query := `SELECT id, name, username, address, role, can_rent,created_at, updated_at,telp FROM users WHERE id = $1`
	var usersItem dto.GetUsers
	if err := r.db.QueryRow(query, uuid).Scan(
		&usersItem.Uuid,
		&usersItem.Name,
		&usersItem.Username,
		&usersItem.Address,
		&usersItem.Role,
		&usersItem.Can_rent,
		&usersItem.Created_at,
		&usersItem.Updated_at,
		&usersItem.Telp,
	); err != nil {
		return nil, err
	}

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usersItem.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	// usersItem.Password = string(hashedPassword)

	return &usersItem, nil
}

func (r *usersRepository) GetAll() ([]*dto.GetUsers, error) {
	query := `SELECT id, name, username, address, role, can_rent, created_at, updated_at, telp FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*dto.GetUsers
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

		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *usersRepository) UpdateUsers(usersUpdate *dto.Users) error {
	query := `
        UPDATE users
        SET name = $1, 
            username = $2, 
            password = $3, 
            address = $4, 
            role = $5, 
            can_rent = $6, 
            telp = $7
        WHERE id = $8
    `
	result, err := r.db.Exec(query,
		usersUpdate.Name,
		usersUpdate.Username,
		usersUpdate.Password,
		usersUpdate.Address,
		usersUpdate.Role,
		usersUpdate.Can_rent,
		usersUpdate.Telp,
		usersUpdate.Uuid,
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

func (c *usersRepository) RegisterUsers(newUsers *dto.RegisterUsers) error {

	query := `INSERT INTO users (name, username, password, address, role, can_rent,created_at, telp)
	values ($1,$2,$3,$4,$5,$6,$7,$8)`
	result, err := c.db.Exec(query,
		newUsers.Name,
		newUsers.Username,
		newUsers.Password,
		newUsers.Address,
		newUsers.Role,
		newUsers.Can_rent,
		newUsers.Created_at,
		newUsers.Telp,
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
