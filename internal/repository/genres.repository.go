package repository

import (
	"context"
	"log"

	"github.com/ariboss89/tickitz-services/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GenresRepository struct {
	db *pgxpool.Pool
}

func NewGenresRepository(db *pgxpool.Pool) *GenresRepository {
	return &GenresRepository{
		db: db,
	}
}

func (m GenresRepository) GetAllGenres(ctx context.Context) ([]model.Genres, error) {
	sqlStr := "SELECT id, name FROM genres ORDER BY id ASC"
	rows, err := m.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var genres []model.Genres
	for rows.Next() {
		var genre model.Genres
		err := rows.Scan(&genre.Id, &genre.Name)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}
