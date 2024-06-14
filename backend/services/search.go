package services

import (
	"backend/clients"
	"backend/dao"
)

func Search(query string) ([]dao.Course, error) {
	var courses []dao.Course
	result := clients.DB.Where("title LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses)
	return courses, result.Error
}

func SearchByID(id string) (dao.Course, error) {
	var course dao.Course
	result := clients.DB.First(&course, id)
	return course, result.Error
}
