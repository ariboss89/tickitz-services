package service

import (
	"context"
	"errors"
	"slices"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminService struct {
	adminRepository *repository.AdminRepository
	db              *pgxpool.Pool
}

func NewAdminService(adminRepository *repository.AdminRepository, db *pgxpool.Pool) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
		db:              db,
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

func (a AdminService) UpdateStatusByOrderId(ctx context.Context, sts dto.UpdateStatusOrder) error {
	status := []string{"pending", "done", "cancelled"}
	isAvailable := slices.Contains(status, sts.Status)

	if !isAvailable {
		return errors.New("status is not valid")
	}

	cmd, err := a.adminRepository.UpdateStatusByOrderId(ctx, a.db, sts)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("no data deleted")
	}
	return nil
}
