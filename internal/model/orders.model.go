package model

import "time"

type Schedules struct {
	Id             int       `db:"id"`
	Show_Date      time.Time `db:"show_date"`
	Show_Time      time.Time `db:"show_time"`
	Price          int       `db:"price"`
	Movie_Name     string    `db:"movie_name"`
	Cinema_Name    string    `db:"cinema_name"`
	Cinema_City    string    `db:"cinema_city"`
	Cinema_Address string    `db:"cinema_address"`
	Studio_Name    string    `db:"studio_name"`
}

type Order struct {
	Id           string  `db:"order_id,omitempty"`
	Schedule_Id  int     `db:"schedule_id,omitempty"`
	Total_Ticket int     `db:"total_ticket,omitempty"`
	Sub_Total    float32 `db:"sub_total,omitempty"`
	Tax          float32 `db:"tax,omitempty"`
	Total_Price  float32 `db:"total_price,omitempty"`
	Status       string  `db:"status,omitempty"`
	Booking_Code *string `db:"booking_code,omitempty"`
	Point        *int    `db:"point,omitempty"`
	User_Id      int     `db:"user_id,omitempty"`
	Payment_id   int     `db:"payment_id,omitempty"`
}

type DetailOrder struct {
	Order_Id    string `db:"order_id,omitempty"`
	Schedule_Id int    `db:"schedule_id,omitempty"`
	Seat_Id     int    `db:"seat_id,omitempty"`
}
