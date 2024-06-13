package dao

import (
	"time"
)

type Subscription struct {
	ID           uint      `gorm:"primaryKey;column:Id_subscription;autoIncrement"`
	UserID       uint      `gorm:"column:User_id;not null"`
	CourseID     uint      `gorm:"column:Course_id;not null"`
	CreationDate time.Time `gorm:"column:Creation_date;autoCreateTime"`
	LastUpdated  time.Time `gorm:"column:Last_updated;autoUpdateTime"`
}
