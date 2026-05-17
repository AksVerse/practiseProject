package repositories

import (
	"project/config"
	"project/models"
)

func GetAllStudents() ([]models.Student, error) {

	var students []models.Student

	result := config.DB.Find(&students)

	return students, result.Error
}

func GetStudentById(id uint) (models.Student, error) {
	
	var student models.Student

	res := config.DB.First(&student, id)

	return student, res.Error
}

func AddStudent(student models.Student) (models.Student, error) {
	result := config.DB.Create(&student)

	return student, result.Error
}

func UpdateStudent(student models.Student) (models.Student, error) {
	result := config.DB.Save(&student)

	return student, result.Error
}

func DeleteStudent(id uint) error {
	result := config.DB.Delete(&models.Student{},id)

	return result.Error
}