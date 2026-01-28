package dto

import (
	"mime/multipart"
	"time"
)

type Users struct {
	Id         int     `json:"id,omitempty"`
	First_Name string  `json:"first_name,omitempty"`
	Last_Name  string  `json:"last_name,omitempty"`
	Email      string  `json:"email"`
	Phone      *string `json:"phone,omitempty"`
	Image      *string `json:"image,omitempty"`
}

type NewUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfile struct {
	First_Name string `form:"first_name,omitempty" json:"first_name"`
	Last_Name  string `form:"last_name,omitempty" json:"last_name"`
	Phone      string `form:"phone,omitempty" json:"phone"`
	// Email      string                `form:"email,omitempty" json:"email"`
	ImageFile *multipart.FileHeader `form:"image,omitempty" json:"image_file"`
	ImageName string                `form:"image_name, omitempty" json:"image_name"`
}

type UpdatePassword struct {
	Old_Password string `json:"old_password,omitempty"`
	New_Password string `json:"new_password,omitempty"`
}

type History struct {
	Id             string    `json:"order_id,omitempty"`
	Show_Date      time.Time `json:"show_date,omitempty"`
	Show_Time      time.Time `json:"show_time,omitempty"`
	Movie_Title    string    `json:"movie_title,omitempty"`
	Total_Price    string    `json:"total_price,omitempty"`
	Status         string    `json:"status,omitempty"`
	Booking_Code   string    `json:"booking_code,omitempty"`
	Payment_Method string    `json:"payment_method,omitempty"`
	Seat_Number    string    `json:"seat_number,omitempty"`
	First_Name     string    `json:"first_name,omitempty"`
	Last_Name      string    `json:"last_name,omitempty"`
	Phone          string    `json:"phone,omitempty"`
}

type Logout struct {
	Token string `json:"token,omitempty"`
}
