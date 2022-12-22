package main

import (
	"database/sql"
	"fmt"
	"os"
	"session-2/controllers"
	_ "session-2/docs"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4444
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	router := gin.Default()

	connString := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(fmt.Errorf("error when open sql connection. error: %w", err))
	}

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error when open ping the database. error: %w", err))
	}

	server := controllers.Server{
		DB: db,
	}

	v1 := router.Group("/api/v1")
	v1.GET("/ping", controllers.Ping)

	todo := v1.Group("/todos")
	{
		todo.POST("/", server.CreatedTodo)
		todo.GET("/", server.GetAll)
		todo.GET("/:id", server.GetByID)
		todo.DELETE("/:id", server.DeleteByID)
		todo.PUT("/:id", server.UpdateByID)
	}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")))
}
