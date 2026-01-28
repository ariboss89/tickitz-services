package dto

type Response struct {
	Msg     string         `json:"msg"`
	Success bool           `json:"success"`
	Data    []any          `json:"data"`
	Error   string         `json:"error,omitempty"`
	Meta    PaginationMeta `json:"meta"`
}

type PaginationMeta struct {
	Page      int    `json:"page,omitempty"`
	TotalPage int    `json:"total_page,omitempty"`
	NextPage  string `json:"next_page,omitempty"`
	PrevPage  string `json:"prev_page,omitempty"`
}

type RegisterResponse struct {
	Email string `json:"email,omitempty"`
}

type LoginResponse struct {
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}

type ResponseError struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type UpdatePasswordResponse struct {
	Email string `json:"email,omitempty"`
}

type OrderResponse struct {
	Id string `json:"order_id"`
}

type LogoutResponse struct {
	Msg string `json:"msg"`
}
