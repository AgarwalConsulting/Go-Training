package blog

import "time"

// Post represents a blog post
type Post struct {
	ID          int
	Title       string
	Body        string
	PublishTime time.Time
}
