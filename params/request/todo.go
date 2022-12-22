package request

type CreateTodo struct {
	Title       string `example:"Title Todo" json:"title"`
	Description string `example:"Desc Todo" json:"desc"`
}

type UpdateTodo struct {
	Title       string `example:"Title Todo" json:"title"`
	Description string `example:"Desc Todo" json:"desc"`
}
