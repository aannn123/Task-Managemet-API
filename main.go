package main

import (
	"net/http"
	"tusk/config"
	"tusk/controllers"
	"tusk/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})

	config.CreateOwnerAccount(db)

	// Controller
	userController := controllers.UserController{DB: db}
	taskController := controllers.TaskController{DB: db}

	// Router
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Tusk API")
	})

	router.POST("/users/login", userController.Login)
	router.POST("/users", userController.CreateAccount)
	router.DELETE("/users/:id", userController.Delete)
	router.GET("/users/Employee", userController.GetEmployee)

	router.POST("/tasks", taskController.Create)
	router.DELETE("/tasks/:id", taskController.Delete)
	router.PATCH("/tasks/:id/submit", taskController.Submit)
	router.PATCH("/tasks/:id/reject", taskController.Reject)
	router.PATCH("/tasks/:id/fix", taskController.Fix)
	router.PATCH("/tasks/:id/approve", taskController.Approve)
	router.GET("/tasks/:id", taskController.FindById)

	router.Static("/attachments", "./attachments")
	router.Run("192.168.1.15:8080")
}
