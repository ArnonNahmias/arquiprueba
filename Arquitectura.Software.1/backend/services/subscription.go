package services

import (
	"backend/dao"
)

// Subscribe subscribes a user to a course.
func Subscribe(userID int, courseID int) error {
	return dao.SubscribeUserToCourse(userID, courseID)
}
