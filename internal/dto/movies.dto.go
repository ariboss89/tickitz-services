package dto

import (
	"mime/multipart"
	"time"
)

type MoviesParam struct {
	Id   int    `uri:"id"`
	Slug string `uri:"slug"`
}

type Movies struct {
	Id             int        `json:"Id" binding:"required"`
	Title          *string    `json:"title"`
	Synopsys       *string    `json:"synopsis"`
	Poster_Url     *string    `json:"poster_url"`
	Background_Url *string    `json:"background_url"`
	Release_date   *time.Time `json:"release_date"`
	Duration       *int       `json:"duration"`
	Status         *string    `json:"status"`
	Rating         *float32   `json:"rating"`
}

type NewMovies struct {
	Title          string    `json:"title"  binding:"required"`
	Synopsys       string    `json:"synopsis" binding:"required"`
	Poster_Url     string    `json:"poster_url" binding:"required"`
	Background_Url string    `json:"background_url" binding:"required"`
	Release_date   time.Time `json:"release_date" binding:"required"`
	Duration       int       `json:"duration" binding:"required"`
	Status         string    `json:"status" binding:"required"`
	Rating         float32   `json:"rating" binding:"required"`
}

type UpdateMovies struct {
	Title          string  `json:"title"`
	Synopsys       string  `json:"synopsis"`
	Poster_Url     string  `json:"poster_url"`
	Background_Url string  `json:"background_url"`
	Release_Date   string  `json:"release_date"`
	Duration       int     `json:"duration"`
	Status         string  `json:"status"`
	Rating         float32 `json:"rating"`
}

type DeleteMovie struct {
	Title string `json:"title"`
}

type MoviesQuery struct {
	Title string   `form:"title"`
	Genre []string `form:"genre"`
	Page  int      `form:"page"`
}

type SearchMovies struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Poster_Url string `json:"poster_url"`
	Genre      string `json:"genre_name"`
}

type MovieGenres struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

type MovieDetails struct {
	Id_Movie       int        `json:"id_movie"`
	Title          *string    `json:"title"`
	Synopsis       *string    `json:"synopsis"`
	Poster_Url     *string    `json:"poster_url"`
	Background_Url *string    `json:"background_url"`
	Release_Date   *time.Time `json:"release_date"`
	Duration       *int       `json:"duration"`
	Director_Name  *string    `json:"director_name"`
	Genres         *string    `json:"genres_name"`
	Actors         *string    `json:"actor_name"`
}

type PostMovies struct {
	Title          string                `form:"title,omitempty" json:"title"`
	Synopsis       string                `form:"synopsis,omitempty" json:"synopsis"`
	PosterFile     *multipart.FileHeader `form:"poster_file,omitempty" json:"poster_file"`
	Poster_Url     string                `form:"poster_url,omitempty" json:"poster_url"`
	BackgroundFile *multipart.FileHeader `form:"background_file,omitempty" json:"background_file"`
	Background_Url string                `form:"background_url,omitempty" json:"background_url"`
	Release_Date   string                `form:"release_date,omitempty" json:"release_date"`
	Duration       int                   `form:"duration,omitempty" json:"duration"`
	Status         string                `form:"status,omitempty" json:"status"`
	Rating         float32               `form:"rating,omitempty" json:"rating"`
}
