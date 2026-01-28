package model

import "time"

type Movies struct {
	Id             int        `db:"Id"`
	Title          *string    `db:"title"`
	Synopsys       *string    `db:"synopsis"`
	Poster_Url     *string    `db:"poster_url"`
	Background_Url *string    `db:"background_url"`
	Release_Date   *time.Time `db:"release_date"`
	Duration       *int       `db:"duration"`
	Status         *string    `db:"status"`
	Rating         *float32   `db:"rating"`
}

type MovieDetails struct {
	Id_Movie       int        `db:"-"`
	Title          *string    `db:"title"`
	Synopsis       *string    `db:"synopsis"`
	Poster_Url     *string    `db:"poster_url"`
	Background_Url *string    `db:"background_url"`
	Release_Date   *time.Time `db:"release_date"`
	Duration       *int       `db:"duration"`
	Director_Name  *string    `db:"name"`
	Genres         *string    `db:"genre_name"`
	Actors         *string    `db:"actor_name"`
}

type MovieGenres struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Title string `db:"title"`
}

type SearchMovies struct {
	Id         int    `db:"id"`
	Title      string `db:"title"`
	Poster_Url string `db:"poster_url"`
	Duration   int    `db:"duration"`
	Genre      string `db:"genre_name"`
}
