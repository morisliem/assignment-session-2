package view

type DefaultSuccess struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"success"`
}

type DefaultError struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"error"`
	Error   string `json:"error"`
}
