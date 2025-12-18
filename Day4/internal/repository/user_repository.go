package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-db-demo/internal/model"
	"time"
)

var baseKey string = "user:"

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	UpdateUserById(ctx context.Context, user *model.User, id string) (*model.User, error)
	DeleteUserById(ctx context.Context, id string) error
}

type PostgresUserRepository struct {
	db    *sql.DB
	redis *redis.Client
}

// init a new PostgresUserRepository
func NewPostgresUserRepository(db *sql.DB, redis *redis.Client) UserRepository {
	return &PostgresUserRepository{
		db:    db,
		redis: redis,
	}
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
	// redis
	key := baseKey + id
	cached, errCache := r.redis.Get(ctx, key).Bytes()
	if errCache == nil {
		var user model.User
		if err := json.Unmarshal(cached, &user); err == nil {
			fmt.Println("INFO : Using user in redis")
			return &user, nil
		}
	}

	// query
	query := "SELECT id, name, email FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	// check err
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("User not found")
	}

	// cache
	b, _ := json.Marshal(user)
	_ = r.redis.Set(ctx, key, b, 60*time.Second).Err()
	// return
	return user, err
}

// update user by id
func (r *PostgresUserRepository) UpdateUserById(ctx context.Context, user *model.User, id string) (*model.User, error) {

	// query
	query := "SELECT id, name, email from update_user_by_id($1, $2, $3)"
	row := r.db.QueryRowContext(ctx, query, id, user.Name, user.Email)
	userSaved := &model.User{}
	err := row.Scan(&userSaved.ID, &userSaved.Name, &userSaved.Email)
	// check error
	if err != nil {
		return nil, err
	}

	// cache
	fmt.Println("Delete info in redis")
	key := baseKey + id
	_ = r.redis.Del(ctx, key).Err()
	// response
	return userSaved, nil
}

// delete user by id
func (r *PostgresUserRepository) DeleteUserById(ctx context.Context, id string) error {
	query := "SELECT delete_user($1)"
	_, err := r.db.ExecContext(ctx, query, id)

	// cache
	fmt.Println("Delete info in redis")
	key := baseKey + id
	_ = r.redis.Del(ctx, key).Err()
	return err
}
