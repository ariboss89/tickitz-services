package service

import (
	"context"
	"errors"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
)

type AdminService struct {
	adminRepository *repository.AdminRepository
}

func NewAdminService(adminRepository *repository.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
	}
}

func (a AdminService) GetAllMovies(ctx context.Context) ([]dto.Movies, error) {
	data, err := a.adminRepository.GetAllMovies(ctx)
	if err != nil {
		return nil, err
	}
	var response []dto.Movies
	for _, v := range data {
		response = append(response, dto.Movies{
			Id:             v.Id,
			Title:          v.Title,
			Synopsys:       v.Synopsys,
			Poster_Url:     v.Poster_Url,
			Background_Url: v.Background_Url,
			Release_date:   v.Release_Date,
			Duration:       v.Duration,
			Status:         v.Status,
			Rating:         v.Rating,
		})
	}
	return response, nil
}

func (a AdminService) PostMovie(ctx context.Context, post dto.PostMovies) error {
	cmd, err := a.adminRepository.PostMovies(ctx, post)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no data updated")
	}
	// invalidasi cache
	return nil
}

func (a AdminService) UpdateMovie(ctx context.Context, update dto.UpdateMovies, id int) error {
	cmd, err := a.adminRepository.UpdateMovie(ctx, update, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no data updated")
	}
	// invalidasi cache
	return nil
}

func (a AdminService) DeleteMovie(ctx context.Context, id int) (dto.DeleteMovie, error) {
	data, err := a.adminRepository.DeleteMovie(ctx, id)
	if err != nil {
		return dto.DeleteMovie{}, err
	}

	response := dto.DeleteMovie{
		Title: data.Title,
	}

	return response, nil
}
