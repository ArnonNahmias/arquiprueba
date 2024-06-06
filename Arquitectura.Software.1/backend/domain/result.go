package domain

type Result struct {
	Message string `json:"message"`
}

type SearchResponse struct {
	Results []Course `json:"results"`
}

type SubscribeRequest struct {
	UserID   int `json:"user_id"`
	CourseID int `json:"course_id"`
}
