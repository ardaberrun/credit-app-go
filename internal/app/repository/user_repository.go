package repository

import (
	"database/sql"
	"github/ardaberrun/credit-app-go/internal/app/model"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
	GetUsers() ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func InitializeUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	query :=
	"INSERT INTO USERS (role_name, email, hashed_password, account_number, balance, created_at) VALUES ($1, $2, $3, $4, $5, $6)";

	_, err := r.db.Query(query, "user", user.Email, user.HashedPassword, user.AccountNumber, user.Balance, user.CreatedAt);
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUsers() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM USERS")
	if err != nil {
		return nil, err
	}
	defer rows.Close();

	users := []*model.User{}
	for rows.Next() {
		user := new(model.User)

		err := rows.Scan(
			&user.Id,
			&user.RoleName,
			&user.Email,
			&user.HashedPassword,
			&user.AccountNumber,
			&user.Balance,
			&user.CreatedAt);

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(id int) (*model.User, error) {
	user := new(model.User);

	row := r.db.QueryRow("SELECT * FROM USERS WHERE id = $1", id);
	err := row.Scan(&user.Id, &user.RoleName, &user.Email, &user.HashedPassword, &user.AccountNumber, &user.Balance, &user.CreatedAt);
	if err != nil {
		return nil, err
	}

	return user, nil;
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := new(model.User);

	row := r.db.QueryRow("SELECT * FROM USERS WHERE email = $1", email);
	err := row.Scan(&user.Id, &user.RoleName, &user.Email, &user.HashedPassword, &user.AccountNumber, &user.Balance, &user.CreatedAt);
	if err != nil {
		return nil, err
	}

	return user, nil;
}