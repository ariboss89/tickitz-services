package repository

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (a AdminRepository) GetAllMovies(ctx context.Context) ([]model.Movies, error) {
	sqlStr := "SELECT id, title, synopsis, poster_url, background_url, release_date, duration, status, rating FROM movies WHERE deleted_at is null ORDER BY id ASC"
	rows, err := a.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var movies []model.Movies
	for rows.Next() {
		var movie model.Movies
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Synopsys, &movie.Poster_Url, &movie.Background_Url, &movie.Release_Date, &movie.Duration, &movie.Status, &movie.Rating)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

// func (a AdminRepository) UpdateMovie(ctx context.Context, update dto.UpdateMovies) (model.Movies, error) {
// 	sqlStr := "UPDATE movies SET title = $1, synopsis = $2, poster_url=$3, background_url=$4, release_date = $5, duration = $6, status=$7, rating= $8,updated_at = NOW() WHERE id = $9 RETURNING id, title, synopsis, poster_url, background_url, release_date, duration, status, rating"

// 	if update.Title != "" {
// 		log.Println(errors.New("title can't be empty"))
// 		return model.Movies{}, errors.New("title can't be empty")
// 	} else if update.Duration == 0 {
// 		log.Println(errors.New("duration can't be empty"))
// 		return model.Movies{}, errors.New("duration can't be empty")
// 	} else if update.Status != "upcoming" && update.Status != "now_showing" {
// 		log.Println(errors.New("status must upcoming or now_showing"))
// 		return model.Movies{}, errors.New("status must upcoming or now_showing")
// 	} else if update.Rating == 0.0 {
// 		log.Println(errors.New("rating can't be empty"))
// 		return model.Movies{}, errors.New("rating can't be empty")
// 	} else if update.Id == 0 {
// 		log.Println(errors.New("id can't be empty"))
// 		return model.Movies{}, errors.New("id can't be empty")
// 	}

// 	values := []any{update.Title, update.Synopsys, update.Poster_Url, update.Background_Url, update.Release_Date, update.Duration, update.Status, update.Rating, update.Id}

// 	row := a.db.QueryRow(ctx, sqlStr, values...)

// 	var movie model.Movies

// 	if err := row.Scan(&movie.Id, &movie.Title, &movie.Synopsys, &movie.Poster_Url, &movie.Background_Url, &movie.Release_Date, &movie.Duration, &movie.Status, &movie.Rating); err != nil {
// 		log.Println(err.Error())
// 		return model.Movies{}, err
// 	}

// 	return movie, nil
// }

func (a AdminRepository) DeleteMovie(ctx context.Context, id int) (dto.DeleteMovie, error) {
	sqlStr := "UPDATE movies SET deleted_at = NOW() WHERE id = $1 RETURNING title"
	values := id
	row := a.db.QueryRow(ctx, sqlStr, values)

	var delete dto.DeleteMovie
	if err := row.Scan(&delete.Title); err != nil {
		log.Println(err.Error())
		return dto.DeleteMovie{}, err
	}
	return delete, nil
}

func (a AdminRepository) UpdateMovie(ctx context.Context, update dto.UpdateMovies, id int) (pgconn.CommandTag, error) {
	var sql strings.Builder
	values := []any{}
	valuesAll := []any{}

	//checkDate := pkg.IsValidDate(update.Release_Date.String())
	//const layout = "2006-01-02 15:04:05"

	//timex, errTime := time.Parse(layout, update.Release_Date.String())

	// if errTime != nil {
	// 	log.Println(errTime)
	// }

	sql.WriteString("UPDATE movies SET")
	if update.Title != "" {
		fmt.Fprintf(&sql, " title=$%d", len(values)+1)
		values = append(values, update.Title)
		valuesAll = append(valuesAll, &update.Title)
	}
	if update.Synopsys != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " synopsis=$%d", len(values)+1)
		values = append(values, update.Synopsys)
		valuesAll = append(valuesAll, &update.Synopsys)
	}
	if update.Poster_Url != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " poster_url=$%d", len(values)+1)
		values = append(values, update.Poster_Url)
		valuesAll = append(valuesAll, &update.Poster_Url)
	}
	if update.Background_Url != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " background_url=$%d", len(values)+1)
		values = append(values, update.Background_Url)
		valuesAll = append(valuesAll, &update.Background_Url)
	}
	// if errTime != nil {
	// 	log.Println(errTime)
	// } else {
	if update.Release_Date != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " release_date=$%d", len(values)+1)
		values = append(values, update.Release_Date)
		valuesAll = append(valuesAll, &update.Release_Date)
	}
	// }
	if update.Status != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " status=$%d", len(values)+1)
		values = append(values, update.Status)
		valuesAll = append(valuesAll, &update.Status)
	}
	if update.Rating != 0.0 {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " rating=$%d", len(values)+1)
		values = append(values, update.Rating)
		valuesAll = append(valuesAll, &update.Rating)
	}
	if update.Duration != 0 {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " duration=$%d", len(values)+1)
		//values = append(values, update.Duration)
		valuesAll = append(valuesAll, &update.Duration)
	}
	if update.Title != "" || update.Synopsys != "" || update.Poster_Url != "" || update.Background_Url != "" || update.Status != "" || update.Duration != 0 || update.Rating != 0.0 {
		sql.WriteString(" WHERE ")
		fmt.Fprintf(&sql, "Id=%d", id)
	}

	sqlStr := sql.String()

	// valuesAll := []any{update.Title, update.Synopsys, update.Poster_Url, update.Background_Url, update.Release_Date, update.Duration, update.Status, update.Rating, update.Id}
	return a.db.Exec(ctx, sqlStr, valuesAll...)
}

func (a AdminRepository) PostMovies(ctx context.Context, post dto.PostMovies) (pgconn.CommandTag, error) {

	sqlStr := "INSERT INTO movies (title, synopsis, poster_url, background_url, release_date, duration, status, rating) VALUES (($1), ($2), ($3), ($4), ($5), ($6), ($7), ($8))"
	values := []any{post.Title, post.Synopsis, post.Poster_Url, post.Background_Url, post.Release_Date, post.Duration, post.Status, post.Rating}

	return a.db.Exec(ctx, sqlStr, values...)
}

func (a AdminRepository) UpdateStatusByOrderId(ctx context.Context, db DBTX, updt dto.UpdateStatusOrder) (pgconn.CommandTag, error) {
	sqlStr := `
			UPDATE orders SET status = $1 WHERE id = $2
			`
	values := []any{strings.ToLower(updt.Status), updt.OrderId}
	return db.Exec(ctx, sqlStr, values...)
}
