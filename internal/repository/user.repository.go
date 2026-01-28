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

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) GetUserProfileByEmail(ctx context.Context, email string) ([]model.Users, error) {
	sqlStr := "SELECT id, first_name, last_name, email, phone, image FROM users WHERE email = $1"
	value := email
	rows, err := u.db.Query(ctx, sqlStr, value)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Phone, &user.Image)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u UserRepository) GetHistory(ctx context.Context, id int) ([]model.History, error) {
	sqlStr := `
	SELECT
    o.id, 
    s.show_date,
    s.show_time,
    m.title,
    o.total_price, 
    o.status, 
    o.booking_code, 
    p.payment_method,
    string_agg(se.seat_number, ',' ),
    u.first_name,
    u.last_name,
    u.phone
	FROM orders o 
	JOIN dt_orders dto ON dto.order_id = o.id
	JOIN users u ON o.user_id = u.id
	JOIN schedules s ON s.id = o.schedule_id 
	JOIN payments p ON p.id = o.payment_id
	JOIN movies m ON m.id = s.movie_id 
	JOIN seats se ON se.id = dto.seat_id
	WHERE o.status = 'paid' AND o.user_id = $1
	GROUP BY o.id, s.id, p.id, m.title, u.id;
	`
	value := id
	rows, err := u.db.Query(ctx, sqlStr, value)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var histories []model.History
	for rows.Next() {
		var history model.History
		err := rows.Scan(&history.Id, &history.Show_Date, &history.Show_Time, &history.Movie_Title, &history.Total_Price, &history.Status, &history.Booking_Code, &history.Payment_Method, &history.Seat_Number, &history.First_Name, &history.Last_Name, &history.Phone)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}

func (u UserRepository) UpdateProfile(ctx context.Context, update dto.UpdateProfile, email string) (pgconn.CommandTag, error) {
	var sql strings.Builder
	values := []any{}
	//valuesAll := []any{}

	sql.WriteString("UPDATE users SET")
	if update.First_Name != "" {
		fmt.Fprintf(&sql, " first_name=$%d", len(values)+1)
		values = append(values, update.First_Name)
		//valuesAll = append(valuesAll, &update.First_Name)
	}
	if update.Last_Name != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " last_name=$%d", len(values)+1)
		values = append(values, update.Last_Name)
		//valuesAll = append(valuesAll, &update.Last_Name)
	}
	if update.ImageName != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " image=$%d", len(values)+1)
		values = append(values, fmt.Sprintf("/profile/%s", update.ImageName))
		//valuesAll = append(valuesAll, fmt.Sprintf("/profile/%s", &update.ImageName))
	}
	if update.Phone != "" {
		if len(values) > 0 {
			sql.WriteString(",")
		}
		fmt.Fprintf(&sql, " phone=$%d", len(values)+1)
		values = append(values, update.Phone)
		//valuesAll = append(valuesAll, &update.Phone)
	}
	if update.First_Name != "" || update.Last_Name != "" || update.Phone != "" || email != "" || update.ImageName != "" {
		sql.WriteString(", updated_at= NOW() WHERE ")
		fmt.Fprintf(&sql, "email='%s'", email)
	}

	sqlStr := sql.String()

	return u.db.Exec(ctx, sqlStr, values...)
}
