package domain

import "time"

type Course struct {
	IdCurso    int64     `json:"id"`          // Course ID
	Nombre     string    `json:"title"`       // Course title
	Dificultad string    `json:"description"` // Course description
	Precio     string    `json:"category"`    // Course Category. Allowed values: to be defined
	ImageURL   string    `json:"image_url"`
	CreatedAt  time.Time `json:"creation_date"` // Course creation date
	UpdatedAt  time.Time `json:"last_updated"`  // Course last updated date
}

type SearchResponse struct {
	Results []Course `json:"results"`
}

type SubscribeRequest struct {
	UserID   int64 `json:"user_id"`
	CourseID int64 `json:"course_id"`
}

type SearchRequest struct {
	IdCurso int `json:"IdCurso"`
}

