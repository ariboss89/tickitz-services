package service

import (
	"context"
	"errors"
	"strings"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/repository"
	"github.com/redis/go-redis/v9"
)

//var ErrInvalidGender = errors.New("invalid gender")

type MovieService struct {
	movieRepository *repository.MoviesRepository
	redis           *redis.Client
}

func NewMovieService(movieRepository *repository.MoviesRepository, rdb *redis.Client) *MovieService {
	return &MovieService{
		movieRepository: movieRepository,
		redis:           rdb,
	}
}

func (m MovieService) GetMoviesByStatus(ctx context.Context, status string) ([]dto.Movies, error) {
	//cek cache
	// rkey := "ari:tickitz:movies"
	// rsc := m.redis.Get(ctx, rkey)
	// if rsc.Err() == nil {
	// 	var result []dto.Movies
	// 	cache, err := rsc.Bytes()
	// 	if err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		if err := json.Unmarshal(cache, &result); err != nil {
	// 			log.Println(err.Error())
	// 		} else {
	// 			return result, nil
	// 		}
	// 	}
	// }

	// if rsc.Err() == redis.Nil {
	// 	log.Println("movies cache miss")
	// }

	var response []dto.Movies

	if strings.ToLower(status) == "upcoming" || strings.ToLower(status) == "now_showing" {
		data, err := m.movieRepository.GetMoviesByStatus(ctx, status)
		if err != nil {
			return []dto.Movies{}, err
		}
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
	} else if strings.ToLower(status) == "" || strings.ToLower(status) == "popular" {
		data, err := m.movieRepository.GetAllMovies(ctx)
		if err != nil {
			return []dto.Movies{}, err
		}
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
	} else {
		return nil, errors.New("status invalid")
	}

	// cacheStr, err := json.Marshal(response)
	// if err != nil {
	// 	log.Println(err)
	// 	log.Println("failed to marshal")
	// }

	// rdsStatus := m.redis.Set(ctx, rkey, string(cacheStr), time.Minute*10)
	// if rdsStatus.Err() != nil {
	// 	log.Println("caching failed")
	// 	log.Println(rdsStatus.Err().Error())
	// }

	return response, nil
}

func (m MovieService) GetMovieGenresById(ctx context.Context, id int) ([]dto.MovieGenres, error) {
	data, err := m.movieRepository.GetMovieGenresById(ctx, id)
	if err != nil {
		return nil, err
	}
	var response []dto.MovieGenres
	for _, v := range data {
		response = append(response, dto.MovieGenres{
			Id:    v.Id,
			Title: v.Title,
			Name:  v.Name,
		})
	}
	return response, nil
}

func (m MovieService) GetMovieDetailById(ctx context.Context, id int) ([]dto.MovieDetails, error) {
	data, err := m.movieRepository.GetMovieDetailById(ctx, id)
	if err != nil {
		return nil, err
	}
	var response []dto.MovieDetails
	for _, v := range data {
		response = append(response, dto.MovieDetails{
			Id_Movie:       id,
			Title:          v.Title,
			Synopsis:       v.Synopsis,
			Poster_Url:     v.Poster_Url,
			Background_Url: v.Background_Url,
			Release_Date:   v.Release_Date,
			Duration:       v.Duration,
			Director_Name:  v.Director_Name,
			Genres:         v.Genres,
			Actors:         v.Actors,
		})
	}
	return response, nil
}

func (m MovieService) SearchMoviesByTitleAndGenre(ctx context.Context, title string, genre []string, page int) ([]dto.SearchMovies, error) {
	data, err := m.movieRepository.SearchMoviesByTitleAndGenre(ctx, title, genre, page)
	if err != nil {
		return nil, err
	}
	var response []dto.SearchMovies
	for _, v := range data {
		response = append(response, dto.SearchMovies{
			Id:         v.Id,
			Title:      v.Title,
			Poster_Url: v.Poster_Url,
			Genre:      v.Genre,
		})
	}
	return response, nil
}
