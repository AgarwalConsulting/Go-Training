package blog

import (
	"encoding/json"
	"net/http"
	"time"
)

var posts = []Post{
	{1, "Hello World!", "This is our first post.", time.Now()},
	{2, "Another World!", "This is our second post.", time.Now()},
}

// GetPostsHandler returns all blog posts
func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}
