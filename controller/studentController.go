package controller

import (
	"net/http"
	"strconv"

	"project/dto"
	"project/models"
	"project/services"

	"github.com/gin-gonic/gin"
)

func GetStudents(c *gin.Context) {

	students, err := services.FetchStudents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context){
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"Invalid student id",
		})
		return
	}

	student, err := services.GetStudentById(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"Error":"Student not found",
		})
		return
	}

	c.JSON(http.StatusOK,student)

}

func AddStudent(c *gin.Context){
	var student models.Student

	err := c.ShouldBindJSON(&student)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Err_Msg":err.Error(),
		})
		return
	}

	newStudent, err := services.AddStudent(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errro": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated,newStudent)
}

func UpdateStudent(c *gin.Context){

	idParam := c.Param("id")
	
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error":err.Error(),
		})
		return
	}

	var req dto.UpdateStudentRequest

	err = c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"Error":err.Error(),
		})
		return
	}

	updateStudent, err := services.UpdateStudent(uint(id),req)

	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,&updateStudent)

}

func DeleteStudent(c *gin.Context){
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"Error": "Invalid student id",
		})
		return
	}

	err = services.DeleteStudent(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Message":"Student Deleted Successfully",
	})
}