package repository

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/ariboss89/tickitz-services/internal/dto"
	"github.com/ariboss89/tickitz-services/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type OrdersRepository struct {
}

func NewOrdersRepository(db *pgxpool.Pool) *OrdersRepository {
	return &OrdersRepository{}
}

func (o OrdersRepository) GetSchedule(ctx context.Context, db DBTX, date, time, location string, movieId int) ([]model.Schedules, error) {
	sqlStr := `
	SELECT s.id,
       s.show_date,
       s.show_time,
       s.price,
       m.title,
       c.name,
			 c.address,
			 c.city,
			 st.studio_name
	FROM schedules s
	JOIN cinemas c ON c.id = s.cinema_id
	JOIN studios st ON st.id = c.id
	JOIN movies m ON m.id = s.movie_id
  WHERE s.movie_id = $1 AND c.city = $2 AND s.show_time = $3 AND s.show_date= $4
  ORDER BY id ASC
	`
	values := []any{}

	movieStr := strconv.Itoa(movieId)

	if movieId != 0 {
		values = append(values, movieStr)
	}
	if location != "" {
		values = append(values, location)
	}
	if time != "" {
		values = append(values, time)
	}
	if date != "" {
		values = append(values, date)
	}

	rows, err := db.Query(ctx, sqlStr, values...)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var schedules []model.Schedules
	for rows.Next() {
		var schedule model.Schedules
		err := rows.Scan(&schedule.Id, &schedule.Show_Date, &schedule.Show_Time, &schedule.Price,
			&schedule.Movie_Name, &schedule.Cinema_Name, &schedule.Cinema_Address, &schedule.Cinema_City, &schedule.Studio_Name)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func (o OrdersRepository) CreateOrder(ctx context.Context, db DBTX, userId int, newOrder dto.CreateOrder) (dto.OrderResponse, error) {
	var orderId string
	sqlCheck := "INSERT INTO orders (schedule_id, total_ticket, sub_total, tax, total_price,booking_code, user_id, payment_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	tax := (newOrder.Sub_Total * 10) / 100
	totalPrice := newOrder.Sub_Total + tax

	newOrder.Tax = tax
	newOrder.Total_Price = totalPrice

	bookingCode := fmt.Sprintf("CNM-%d", time.Now().UnixNano())

	values := []any{newOrder.Schedule_Id, newOrder.Total_Ticket, newOrder.Sub_Total, newOrder.Tax, newOrder.Total_Price, bookingCode, userId, newOrder.Payment_id}

	row := db.QueryRow(ctx, sqlCheck, values...)

	if err := row.Scan(&orderId); err != nil {
		log.Println(err.Error())
		return dto.OrderResponse{}, err
	}

	// if err != nil {
	// 	log.Println(err)
	// 	return dto.OrderResponse{}, err
	// }

	return dto.OrderResponse{
		Id: orderId,
	}, nil
}

func (o OrdersRepository) CreateDetailOrder(ctx context.Context, db DBTX, orderId string, seatId int) (pgconn.CommandTag, error) {

	sqlCheck := "INSERT INTO dt_orders (order_id, seat_id) VALUES ($1, $2)"

	values := []any{orderId, seatId}

	return db.Exec(ctx, sqlCheck, values...)
}
