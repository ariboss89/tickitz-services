package service

import (
	"context"
	"errors"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
)

type UserRepo interface {
	GetUserProfileByEmail(ctx context.Context, email string)
}

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u UserService) GetUserProfileByEmail(ctx context.Context, email string) ([]dto.Users, error) {
	data, err := u.userRepository.GetUserProfileByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	var response []dto.Users
	for _, v := range data {
		response = append(response, dto.Users{
			First_Name: *v.First_Name,
			Last_Name:  *v.Last_Name,
			Email:      v.Email,
			Phone:      v.Phone,
			Image:      v.Image,
		})
	}
	return response, nil
}

func (u UserService) GetHistory(ctx context.Context, id int) ([]dto.History, error) {
	data, err := u.userRepository.GetHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	var response []dto.History
	for _, v := range data {
		response = append(response, dto.History{
			Id:             v.Id,
			Show_Date:      v.Show_Date,
			Show_Time:      v.Show_Time,
			Movie_Title:    v.Movie_Title,
			Total_Price:    v.Total_Price,
			Status:         v.Status,
			Booking_Code:   v.Booking_Code,
			Payment_Method: v.Payment_Method,
			Seat_Number:    v.Seat_Number,
			First_Name:     v.First_Name,
			Last_Name:      v.Last_Name,
			Phone:          v.Phone,
		})
	}
	return response, nil
}

func (u UserService) UpdateProfile(ctx context.Context, update dto.UpdateProfile, email string) error {
	cmd, err := u.userRepository.UpdateProfile(ctx, update, email)
	if err != nil {
		return nil
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no data updated")
	}
	// invalidasi cache
	return nil
}

// func (u UserService) UpdatePassword(ctx context.Context, update dto.UpdatePassword) (dto.UpdatePassword, error) {
// 	var user dto.Login
// 	var resp dto.UpdatePassword
// 	var auth AuthService

// 	user.Email = update.Email
// 	user.Password = update.Old_Password

// 	hc := pkg.HashConfig{}

// 	dataLogin, err := auth.authRepository.Login(ctx, user)
// 	if err != nil {
// 		return dto.UpdatePassword{}, err
// 	}

// 	hp, err := hc.ComparePwdAndHash(update.Old_Password, dataLogin.Password)
// 	if err != nil {
// 		return dto.UpdatePassword{}, err
// 	}

// 	if hp {
// 		hc.UseRecommended()

// 		hp, err := hc.GenHash(update.New_Password)
// 		if err != nil {
// 			return dto.UpdatePassword{}, err
// 		}

// 		update.New_Password = hp
// 		data, err := u.userRepository.UpdatePassword(ctx, update)
// 		if err != nil {
// 			return dto.UpdatePassword{}, err
// 		}
// 		resp = dto.UpdatePassword{
// 			Email: data.Email,
// 		}
// 	}

// 	return resp, nil
// }
