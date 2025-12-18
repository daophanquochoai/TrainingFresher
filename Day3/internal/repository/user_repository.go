package repository

import (
	"context"
	"database/sql"
	"go-db-demo/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	UpdateUserById(ctx context.Context, user *model.User, id string) (*model.User, error)
	DeleteUserById(ctx context.Context, id string) error
}

type PostgresUserRepository struct {
	db *sql.DB
}

// init a new PostgresUserRepository
func NewPostgresUserRepository(db *sql.DB) UserRepository {
	return &PostgresUserRepository{db: db}
}

// create a new user
func (r *PostgresUserRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {

	query := `INSERT INTO users ( name, email) VALUES ($1, $2) RETURNING id, name, email`
	row := r.db.QueryRowContext(ctx, query, user.Name, user.Email)
	var createdUser model.User
	if err := row.Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email); err != nil {
		return nil, err
	}
	return &createdUser, nil
}

// get user by id
func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, err
}

// update user by id
func (r *PostgresUserRepository) UpdateUserById(ctx context.Context, user *model.User, id string) (*model.User, error) {
	query := "SELECT id, name, email from update_user_by_id($1, $2, $3)"
	row := r.db.QueryRowContext(ctx, query, id, user.Name, user.Email)
	userSaved := &model.User{}
	err := row.Scan(&userSaved.ID, &userSaved.Name, &userSaved.Email)
	if err != nil {
		return nil, err
	}
	return userSaved, nil
}

// delete user by id
func (r *PostgresUserRepository) DeleteUserById(ctx context.Context, id string) error {
	query := "SELECT delete_user($1)"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
