package service

import (
	"context"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
)

type ActorService struct {
	actorRepository *repository.ActorsRepository
}

func NewActorService(actorRepository *repository.ActorsRepository) *ActorService {
	return &ActorService{
		actorRepository: actorRepository,
	}
}

func (m ActorService) GetAllActors(ctx context.Context) ([]dto.Actors, error) {
	data, err := m.actorRepository.GetAllActors(ctx)
	if err != nil {
		return nil, err
	}
	var response []dto.Actors
	for _, v := range data {
		response = append(response, dto.Actors{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	return response, nil
}
