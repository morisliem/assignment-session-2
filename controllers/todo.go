package controllers

import (
	"database/sql"
	"net/http"
	"session-2/params/request"
	"session-2/params/view"
	"session-2/repository"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1
// GetTodos godoc
// @Summary Get TODOs
// @Schemes
// @Description Get all todos
// @Tags TODOs
// @Accept json
// @Produce json
// @Success 200 {object} view.GetTodosSuccessSwag
// @Failure 500 {object} view.DefaultError
// @Router /todos [GET]
func (s *Server) GetAll(c *gin.Context) {
	todoRepo := repository.TodoRepository{
		DB: s.DB,
	}

	todos, err := todoRepo.GetAllTodos()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.DefaultError{
			Status:  http.StatusInternalServerError,
			Message: "failed to get todos",
			Error:   err.Error(),
		})
		return
	}

	resp := view.GetTodosSuccessSwag{
		Status:  http.StatusOK,
		Message: "Success",
		Payload: todos,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// CreateTodo godoc
// @Summary Create TODO
// @Schemes
// @Description create todo
// @Tags TODOs
// @Accept json
// @Produce json
// @Param request body request.CreateTodo true "Request Body"
// @Success 200 {object} view.CreateTodoSuccessRespSwag
// @Failure 500 {object} view.CreateTodoFailedRespSwag
// @Router /todos [POST]
func (s *Server) CreatedTodo(c *gin.Context) {
	var req request.CreateTodo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	todoRepo := repository.TodoRepository{
		DB: s.DB,
	}

	err = todoRepo.InsertTodo(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.CreateTodoFailedRespSwag{
			Status:  http.StatusInternalServerError,
			Message: "failed to created todo",
			Error:   err,
		})
		return
	}

	resp := view.CreateTodoSuccessRespSwag{
		Status:  http.StatusCreated,
		Message: "Created",
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// Get todo by ID godoc
// @Summary Get todo by ID
// @Schemes
// @Description Get todo by ID
// @Tags TODOs
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} view.GetTodoSuccessSwag
// @Failure 404 {object} view.DefaultError
// @Failure 500 {object} view.DefaultError
// @Router /todos/{id} [GET]
func (s *Server) GetByID(c *gin.Context) {
	id := c.Param("id")
	todoRepo := repository.TodoRepository{
		DB: s.DB,
	}

	todo, err := todoRepo.GetTodoByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, view.DefaultError{
				Status:  http.StatusNotFound,
				Message: "todo not exists",
				Error:   "",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.DefaultError{
			Status:  http.StatusInternalServerError,
			Message: "failed to get todo",
			Error:   err.Error(),
		})
		return
	}

	resp := view.GetTodoSuccessSwag{
		Status:  http.StatusOK,
		Message: "Success",
		Payload: todo,
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// Update todo by ID godoc
// @Summary Update todo by ID
// @Schemes
// @Description Update todo by ID
// @Tags TODOs
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param request body request.UpdateTodo true "Request Body"
// @Success 200 {object} view.DefaultSuccess
// @Failure 404 {object} view.DefaultError
// @Failure 500 {object} view.DefaultError
// @Router /todos/{id} [PUT]
func (s *Server) UpdateByID(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdateTodo
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, nil)
		return
	}

	todoRepo := repository.TodoRepository{
		DB: s.DB,
	}

	_, err = todoRepo.GetTodoByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, view.DefaultError{
				Status:  http.StatusNotFound,
				Message: "todo not exists",
				Error:   "",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.DefaultError{
			Status:  http.StatusInternalServerError,
			Message: "failed to update todo",
			Error:   err.Error(),
		})
		return
	}

	err = todoRepo.UpdateTodoByID(id, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.DefaultError{
			Status:  http.StatusInternalServerError,
			Message: "failed to update todo",
			Error:   err.Error(),
		})
		return
	}

	resp := view.DefaultSuccess{
		Status:  http.StatusOK,
		Message: "Success",
	}

	c.JSON(resp.Status, resp)
}

// @BasePath /api/v1
// Delete todo by ID godoc
// @Summary Delete todo by ID
// @Schemes
// @Description Delete todo by ID
// @Tags TODOs
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} view.DefaultSuccess
// @Failure 404 {object} view.DefaultError
// @Failure 500 {object} view.DefaultError
// @Router /todos/{id} [DELETE]
func (s *Server) DeleteByID(c *gin.Context) {
	id := c.Param("id")
	todoRepo := repository.TodoRepository{
		DB: s.DB,
	}

	_, err := todoRepo.GetTodoByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, view.DefaultError{
				Status:  http.StatusNotFound,
				Message: "todo not exists",
				Error:   "",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.DefaultError{
			Status:  http.StatusInternalServerError,
			Message: "failed to update todo",
			Error:   err.Error(),
		})
		return
	}

	err = todoRepo.DeleteTodoByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, view.DefaultError{
			Status:  http.StatusInternalServerError,
			Message: "failed to delete todo",
			Error:   err.Error(),
		})
		return
	}

	resp := view.DefaultSuccess{
		Status:  http.StatusOK,
		Message: "Success",
	}

	c.JSON(resp.Status, resp)
}
