package dto

import "time"

type OrdersQuery struct {
	Id       int    `json:"id"`
	Date     string `json:"date"`
	Time     string `json:"time"`
	Location string `json:"location"`
}

type Schedules struct {
	Id             int       `json:"movie_id" binding:"required"`
	Show_Date      time.Time `json:"show_date" binding:"required"`
	Show_Time      time.Time `json:"show_time" binding:"required"`
	Price          int       `json:"price" binding:"required"`
	Movie_Name     string    `json:"movie_name" binding:"required"`
	Cinema_Name    string    `json:"cinema_name" binding:"required"`
	Cinema_City    string    `json:"cinema_city" binding:"required"`
	Cinema_Address string    `json:"cinema_address" binding:"required"`
	Studio_Name    string    `json:"studio_name" binding:"required"`
}

type CreateOrder struct {
	// Id           string  `json:"order_id,omitempty"`
	Schedule_Id  int     `json:"schedule_id,omitempty"`
	Seats        []int   `json:"seat_id,omitempty"`
	Total_Ticket int     `json:"total_ticket,omitempty"`
	Sub_Total    float32 `json:"sub_total,omitempty"`
	Tax          float32 `json:"tax,omitempty"`
	Total_Price  float32 `json:"total_price,omitempty"`
	Payment_id   int     `json:"payment_id,omitempty"`
}

// type CreateOrder struct {
// 	Id           string  `json:"order_id,omitempty"`
// 	Schedule_Id  int     `json:"schedule_id,omitempty"`
// 	Total_Ticket int     `json:"total_ticket,omitempty"`
// 	Sub_Total    float32 `json:"sub_total,omitempty"`
// 	Tax          float32 `json:"tax,omitempty"`
// 	Total_Price  float32 `json:"total_price,omitempty"`
// 	Booking_Code string  `json:"booking_code,omitempty"`
// 	Status       string  `json:"status,omitempty"`
// 	Payment_id   int     `json:"payment_id,omitempty"`
// }

type CreateDetailOrder struct {
	Order_Id    string `json:"order_id,omitempty"`
	Schedule_Id int    `json:"schedule_id,omitempty"`
	Seat_Id     int    `json:"seat_id,omitempty"`
}
