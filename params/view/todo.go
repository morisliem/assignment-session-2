package view

import "time"

type CreateTodoSuccess struct {
	ID          int       `example:"1" json:"id"`
	Title       string    `example:"Title Todo" json:"title"`
	Description string    `example:"Desc Todo" json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateTodoSuccessRespSwag struct {
	Status  int    `json:"status" example:"201"`
	Message string `json:"message" example:"created"`
}

type CreateTodoFailedRespSwag struct {
	Status  int    `json:"status" example:"500"`
	Message string `json:"message" example:"failed"`
	Error   error  `json:"error"`
}

type GetTodo struct {
	ID          int       `example:"1" json:"id"`
	Title       string    `example:"Title Todo" json:"title"`
	Description string    `example:"Desc Todo" json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetTodosSuccessSwag struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Payload []GetTodo `json:"payload"`
}

type GetTodoSuccessSwag struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Payload GetTodo `json:"payload"`
}
