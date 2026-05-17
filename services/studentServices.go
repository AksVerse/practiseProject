package services

import (
	"errors"
	"project/dto"
	"project/models"
	"project/repositories"
)

func FetchStudents() ([]models.Student, error) {

	return repositories.GetAllStudents()
}

func GetStudentById(id uint) (models.Student, error) {
	return repositories.GetStudentById(id)
}

func AddStudent(student models.Student) (models.Student, error) {
	if student.Age <= 16 {
		return models.Student{}, errors.New("Age should be greater than 16")
	}

	return repositories.AddStudent(student)
}

func UpdateStudent(id uint, req dto.UpdateStudentRequest) (models.Student, error) {
	student, err := repositories.GetStudentById(id)

	if err != nil {
		return models.Student{}, errors.New("Student not found")
	}

	if req.Name != nil {
		student.Name = *req.Name
	}

	if req.Age != nil {
		student.Age = *req.Age
	}

	if req.Email != nil {
		student.Email = *req.Email
	}

	return repositories.UpdateStudent(student)
}

func DeleteStudent(id uint) error {
	_, err := repositories.GetStudentById(id)

	if err != nil {
		return errors.New("Student not found")
	}

	return repositories.DeleteStudent(id)
}