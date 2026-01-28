package model

import "time"

type Users struct {
	Id         int     `db:"id,omitempty"`
	First_Name *string `db:"first_name,omitempty"`
	Last_Name  *string `db:"last_name,omitempty"`
	Email      string  `db:"email"`
	Phone      *string `db:"phone,omitempty"`
	Password   string  `db:"password,omitempty"`
	Role       string  `db:"role, omitempty"`
	Image      *string `db:"image, omitempty"`
}

type History struct {
	Id             string    `db:"order_id,omitempty"`
	Show_Date      time.Time `db:"show_date,omitempty"`
	Show_Time      time.Time `db:"show_time,omitempty"`
	Movie_Title    string    `db:"movie_title,omitempty"`
	Total_Price    string    `db:"total_price,omitempty"`
	Status         string    `db:"status,omitempty"`
	Booking_Code   string    `db:"booking_code,omitempty"`
	Payment_Method string    `db:"payment_method,omitempty"`
	Seat_Number    string    `db:"seat_number,omitempty"`
	First_Name     string    `db:"first_name,omitempty"`
	Last_Name      string    `db:"last_name,omitempty"`
	Phone          string    `db:"phone,omitempty"`
}
