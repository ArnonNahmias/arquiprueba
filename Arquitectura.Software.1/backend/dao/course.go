package dao

import (
	"backend/clients"
	"backend/domain"
)

func GetCourses() ([]domain.Course, error) {
	var courses []domain.Course
	result := clients.DB.Find(&courses)
	return courses, result.Error
}

func GetCourseByID(id string) (*domain.Course, error) {
	var course domain.Course
	result := clients.DB.First(&course, "id = ?", id)
	return &course, result.Error
}

func CreateCourse(course *domain.Course) error {
	result := clients.DB.Create(course)
	return result.Error
}

func UpdateCourse(id string, course *domain.Course) error {
	result := clients.DB.Model(&domain.Course{}).Where("id = ?", id).Updates(course)
	return result.Error
}

func DeleteCourse(id string) error {
	result := clients.DB.Delete(&domain.Course{}, id)
	return result.Error
}

func SearchCourses(query string) ([]domain.Course, error) {
	var courses []domain.Course
	result := clients.DB.Where("nombre LIKE ?", "%"+query+"%").Find(&courses)
	return courses, result.Error
}
