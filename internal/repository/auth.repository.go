package repository

import (
	"context"
	"errors"
	"log"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (a AuthRepository) Register(ctx context.Context, newUser dto.NewUser) (model.Users, error) {
	var count int
	sqlCheck := "SELECT count(email) FROM users WHERE email = $1"
	value := newUser.Email

	err := a.db.QueryRow(ctx, sqlCheck, value).Scan(&count)

	if err == nil && count > 0 {
		err = errors.New("user exist")
		return model.Users{}, err
	}

	sqlStr := "INSERT INTO users (email, password, role) VALUES (($1), ($2), 'user') RETURNING email"
	values := []any{newUser.Email, newUser.Password}

	row := a.db.QueryRow(ctx, sqlStr, values...)

	var user model.Users

	if err := row.Scan(&user.Email); err != nil {
		log.Println(err.Error())
		return model.Users{}, err
	}

	return user, nil
}

func (a AuthRepository) Login(ctx context.Context, newUser dto.Login) (model.Users, error) {
	var user model.Users

	sqlStr := "SELECT id, email, password, role FROM users WHERE email = ($1)"
	values := []any{newUser.Email}

	err := a.db.QueryRow(ctx, sqlStr, values...).Scan(&user.Id, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func (a AuthRepository) UpdatePassword(ctx context.Context, update dto.UpdatePassword, email string) (dto.UpdatePasswordResponse, error) {

	sqlStr := "UPDATE users SET password = $1, updated_at = NOW() WHERE email = $2 RETURNING email"

	values := []any{update.New_Password, email}

	var updt dto.UpdatePasswordResponse

	updt.Email = email

	row := a.db.QueryRow(ctx, sqlStr, values...)

	if err := row.Scan(&updt.Email); err != nil {
		log.Println(err.Error())
		return dto.UpdatePasswordResponse{}, err
	}

	return updt, nil
}
