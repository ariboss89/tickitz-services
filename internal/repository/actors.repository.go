package repository

import (
	"context"
	"log"

	"github.com/ariboss89/tickitz-services/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActorsRepository struct {
	db *pgxpool.Pool
}

func NewActorsRepository(db *pgxpool.Pool) *ActorsRepository {
	return &ActorsRepository{
		db: db,
	}
}

func (m ActorsRepository) GetAllActors(ctx context.Context) ([]model.Actors, error) {
	sqlStr := "SELECT id, name FROM actors ORDER BY id ASC"
	rows, err := m.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var actors []model.Actors
	for rows.Next() {
		var actor model.Actors
		err := rows.Scan(&actor.Id, &actor.Name)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		actors = append(actors, actor)
	}
	return actors, nil
}
