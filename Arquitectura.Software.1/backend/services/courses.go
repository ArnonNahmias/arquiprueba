package services

import (
	"backend/dao"
	"backend/domain"
)

func GetCourses() ([]domain.Course, error) {
	return dao.GetCourses()
}

func GetCourseByID(id string) (*domain.Course, error) {
	return dao.GetCourseByID(id)
}

func CreateCourse(course *domain.Course) error {
	return dao.CreateCourse(course)
}

func UpdateCourse(id string, course *domain.Course) error {
	return dao.UpdateCourse(id, course)
}

func DeleteCourse(id string) error {
	return dao.DeleteCourse(id)
}
