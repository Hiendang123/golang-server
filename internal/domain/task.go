package domain

// Task model
type Task struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
