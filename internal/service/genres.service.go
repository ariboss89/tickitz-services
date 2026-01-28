package service

import (
	"context"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
)

type GenreService struct {
	genreRepository *repository.GenresRepository
}

func NewGenreService(genreRepository *repository.GenresRepository) *GenreService {
	return &GenreService{
		genreRepository: genreRepository,
	}
}

func (m GenreService) GetAllGenres(ctx context.Context) ([]dto.Genres, error) {
	data, err := m.genreRepository.GetAllGenres(ctx)
	if err != nil {
		return nil, err
	}
	var response []dto.Genres
	for _, v := range data {
		response = append(response, dto.Genres{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	return response, nil
}
