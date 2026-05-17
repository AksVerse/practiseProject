package routes

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/students", controller.GetStudents)

	router.GET("/student/:id",controller.GetStudentById)

	router.POST("/add-student",controller.AddStudent)

	router.PUT("/student-update/:id",controller.UpdateStudent)

	router.DELETE("/delete-student/:id",controller.DeleteStudent)
}