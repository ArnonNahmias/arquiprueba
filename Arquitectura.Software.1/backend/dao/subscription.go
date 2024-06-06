package dao

import (
	"backend/clients"
	"errors"
)

// Subscription struct
type Subscription struct {
	ID       uint `gorm:"primaryKey"`
	UserID   int
	CourseID int
}

// SubscribeUserToCourse subscribes a user to a course.
func SubscribeUserToCourse(userID int, courseID int) error {
	// Check if the subscription already exists
	var existingSubscription Subscription
	if err := clients.DB.Where("user_id = ? AND course_id = ?", userID, courseID).First(&existingSubscription).Error; err == nil {
		return errors.New("subscription already exists")
	}

	// Create new subscription
	subscription := Subscription{
		UserID:   userID,
		CourseID: courseID,
	}
	result := clients.DB.Create(&subscription)
	return result.Error
}
