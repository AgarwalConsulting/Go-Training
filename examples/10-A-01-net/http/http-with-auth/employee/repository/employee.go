package repository

// Employee represents a single employee details
type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProjectID   int    `json:"project" validate:"gte=1000"`
	Designation string `json:"title"`
}
