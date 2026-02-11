package service

import (
	"context"
	"errors"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/ariboss89/tickitz-services/pkg"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	authRepository *repository.AuthRepository
	redis          *redis.Client
}

func NewAuthService(authRepository *repository.AuthRepository, rdb *redis.Client) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		redis:          rdb,
	}
}

func (a AuthService) Register(ctx context.Context, newUser dto.NewUser) (dto.RegisterResponse, error) {
	hc := pkg.HashConfig{}
	hc.UseRecommended()

	hp, err := hc.GenHash(newUser.Password)
	if err != nil {
		return dto.RegisterResponse{}, err
	}
	newUser.Password = hp
	data, err := a.authRepository.Register(ctx, newUser)
	if err != nil {
		return dto.RegisterResponse{}, err
	}
	response := dto.RegisterResponse{
		Email: data.Email,
	}
	return response, nil
}

func (a AuthService) Login(ctx context.Context, newUser dto.Login) (dto.LoginResponse, error) {
	// w http.ResponseWriter
	var resp dto.LoginResponse
	hc := pkg.HashConfig{}

	data, err := a.authRepository.Login(ctx, newUser)

	if err != nil {
		return dto.LoginResponse{}, err
	}

	hp, err := hc.ComparePwdAndHash(newUser.Password, data.Password)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if hp {
		claims := pkg.NewJWTClaims(data.Id, data.Role, data.Email)
		token, err := claims.GenToken()

		if err != nil {
			return dto.LoginResponse{}, err
		}

		resp = dto.LoginResponse{
			Email: data.Email,
			Role:  data.Role,
			Token: token,
		}

	} else {

		return dto.LoginResponse{}, errors.New("username or password is wrong")
	}

	return resp, nil
}

func (a AuthService) UpdatePassword(ctx context.Context, update dto.UpdatePassword, email string) (dto.UpdatePasswordResponse, error) {

	//data := pkg.IsBlacklisted(, a.redis, token)

	var user dto.Login
	var resp dto.UpdatePasswordResponse

	user.Email = email
	user.Password = update.Old_Password

	hc := pkg.HashConfig{}

	dataLogin, err := a.authRepository.Login(ctx, user)
	if err != nil {
		return dto.UpdatePasswordResponse{}, err
	}

	hp, err := hc.ComparePwdAndHash(update.Old_Password, dataLogin.Password)
	if err != nil {
		return dto.UpdatePasswordResponse{}, err
	}

	if hp {
		hc.UseRecommended()

		hp, err := hc.GenHash(update.New_Password)
		if err != nil {
			return dto.UpdatePasswordResponse{}, err
		}

		update.New_Password = hp

		data, err := a.authRepository.UpdatePassword(ctx, update, email)
		if err != nil {
			return dto.UpdatePasswordResponse{}, err
		}
		resp = dto.UpdatePasswordResponse{
			Email: data.Email,
		}
	}
	return resp, nil
}

func (a AuthService) LogoutUser(ctx context.Context, token string) (bool, error) {
	rkey := "ari:tickitz:logout" + token
	//rsc := a.redis.Get(ctx, rkey)

	rsc, err := a.redis.Exists(ctx, rkey).Result()

	if err != nil {
		return false, err
	}

	if rsc == 0 {
		a.redis.Set(ctx, rkey, token, 0)
	}

	return rsc > 0, nil
	//return true, nil

	// if rsc.Err() == nil {
	// 	tokenStore := rsc.Val()
	// 	if token == tokenStore {
	// 		return errors.New("token obsoleted")
	// 	}
	// 	//return nil
	// }

	// if rsc.Err() == redis.Nil {
	// 	log.Println("users cache miss")
	// }

	// rdsStatus := a.redis.Set(ctx, rkey, token, 0)
	// if rdsStatus.Err() != nil {
	// 	log.Println("caching failed")
	// 	log.Println(rdsStatus.Err().Error())
	// }

	//return nil
}
