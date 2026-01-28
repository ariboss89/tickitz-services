package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ariboss89/tickitz-services/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MoviesRepository struct {
	db *pgxpool.Pool
}

func NewMoviesRepository(db *pgxpool.Pool) *MoviesRepository {
	return &MoviesRepository{
		db: db,
	}
}

func (m MoviesRepository) GetAllMovies(ctx context.Context) ([]model.Movies, error) {
	sqlStr := "SELECT id, title, synopsis, poster_url, background_url, release_date, duration, status, rating FROM movies WHERE status = 'upcoming' AND rating > 7.5 AND deleted_at IS NULL ORDER BY id DESC"

	rows, err := m.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var movies []model.Movies
	for rows.Next() {
		var movie model.Movies
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Synopsys, &movie.Poster_Url,
			&movie.Background_Url, &movie.Release_Date, &movie.Duration, &movie.Status,
			&movie.Rating)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (m MoviesRepository) GetMoviesByStatus(ctx context.Context, status string) ([]model.Movies, error) {
	sqlStr := "SELECT id, title, synopsis, poster_url, background_url, release_date, duration, status, rating FROM movies WHERE status = $1 ORDER BY id ASC"
	value := status

	rows, err := m.db.Query(ctx, sqlStr, value)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var movies []model.Movies
	for rows.Next() {
		var movie model.Movies
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Synopsys, &movie.Poster_Url,
			&movie.Background_Url, &movie.Release_Date, &movie.Duration, &movie.Status,
			&movie.Rating)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (m MoviesRepository) GetMovieGenresById(ctx context.Context, id int) ([]model.MovieGenres, error) {
	sqlStr := "SELECT m.id, g.name, m.title FROM movie_genres mg JOIN movies m ON m.id = mg.movie_id JOIN genres g ON g.id = mg.genre_id WHERE mg.movie_id = $1"
	// row := m.db.QueryRow(ctx, sqlStr, value)
	rows, err := m.db.Query(ctx, sqlStr, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var mgs []model.MovieGenres
	for rows.Next() {
		var mg model.MovieGenres
		err := rows.Scan(&mg.Id, &mg.Name, &mg.Title)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		mgs = append(mgs, mg)
	}
	return mgs, nil
}

func (m MoviesRepository) GetMovieDetailById(ctx context.Context, id int) ([]model.MovieDetails, error) {
	sqlStr :=
		`SELECT 
				m.title, 
				m.synopsis, 
				m.poster_url, 
				m.background_url, 
				m.release_date, 
				m.duration, d.name AS "director name", 
				string_agg(DISTINCT g.name , ', ' ) AS "genres", 
				string_agg(DISTINCT a.name, ', ' ) AS "actors" 
		FROM movie_genres mg 
		JOIN directors d ON d.movie_id = mg.movie_id 
		JOIN genres g ON g.id = mg.genre_id 
		JOIN movies m ON m.id = mg.movie_id 
		JOIN movie_actors ma ON ma.movie_id = mg.movie_id 
		JOIN actors a ON a.id = ma.actor_id 
		WHERE m.id = $1 
		GROUP BY m.id, d.name, ma.movie_id;`

	rows, err := m.db.Query(ctx, sqlStr, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var mds []model.MovieDetails
	for rows.Next() {
		var md model.MovieDetails
		err := rows.Scan(&md.Title, &md.Synopsis, &md.Poster_Url, &md.Background_Url, &md.Release_Date, &md.Duration, &md.Director_Name, &md.Genres, &md.Actors)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		mds = append(mds, md)
	}
	return mds, nil
}

func (m MoviesRepository) SearchMoviesByTitleAndGenre(ctx context.Context, title string, genre []string, page int) ([]model.SearchMovies, error) {
	var sql strings.Builder
	values := []any{}

	sql.WriteString("SELECT m.id, m.title, m.poster_url, STRING_AGG(g.name, ', ') FROM movie_genres mg JOIN movies m ON m.id = mg.movie_id JOIN genres g ON g.id = mg.genre_id")

	if title != "" {
		fmt.Fprintf(&sql, " WHERE m.title ILIKE '%%%s%%' ", title)
	}

	if len(genre) > 0 {
		for i := range genre {
			if title == "" && i == 0 {
				fmt.Fprintf(&sql, " WHERE g.name=$%d ", len(values)+1)
				values = append(values, genre[i])
			}
			if title != "" && i == 0 {
				sql.WriteString("AND")
				fmt.Fprintf(&sql, " g.name=$%d ", len(values)+1)
				values = append(values, genre[i])
			}
			if len(values) > 0 && len(genre) > 1 {
				sql.WriteString("OR")
				fmt.Fprintf(&sql, " g.name=$%d ", len(values)+1)
				values = append(values, genre[i])
			}
		}
	}

	sql.WriteString(" GROUP BY m.id ORDER BY m.id ASC LIMIT 5 OFFSET ")

	offset := (page * 5) - 5

	fmt.Fprintf(&sql, "%d", offset)

	sqlStr := sql.String()

	rows, err := m.db.Query(ctx, sqlStr, values...)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var mds []model.SearchMovies
	for rows.Next() {
		var md model.SearchMovies
		err := rows.Scan(&md.Id, &md.Title, &md.Poster_Url, &md.Genre)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		mds = append(mds, md)
	}

	return mds, nil
}
