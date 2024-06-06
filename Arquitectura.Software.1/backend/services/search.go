package services

import (
	"backend/dao"
	"backend/domain"
)

func Search(query string) ([]domain.Course, error) {
	return dao.SearchCourses(query)
}
