package repository

import (
	"database/sql"
	"go-db-demo/internal/model"
)

type UserRepository interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUserById(user *model.User, id string) (*model.User, error)
	DeleteUserById(id string) error
}

type PostgresUserRepository struct {
	db *sql.DB
}

// init a new PostgresUserRepository
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

// create a new user
func (r *PostgresUserRepository) CreateUser(user *model.User) (*model.User, error) {
	query := `INSERT INTO users ( name, email) VALUES ($1, $2) RETURNING id, name, email`
	row := r.db.QueryRow(query, user.Name, user.Email)
	var createdUser model.User
	if err := row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email); err != nil {
		return nil, err
	}
	return &createdUser, nil
}

// get user by id
func (r *PostgresUserRepository) GetUserByID(id string) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

// update user by id
func (r *PostgresUserRepository) UpdateUserById(user *model.User, id string) (*model.User, error) {
	
}
