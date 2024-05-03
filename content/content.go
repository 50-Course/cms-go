package content

import _ "github.com/golang/protobuf/ptypes/timestamp"

// ImageContent is a content type that represents an image contentj
type ImageContent struct {
	Author  string `json:"author"`
	Url     string `json:"url"`
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	AltText string `json:"alttext"`
}

// VideoContent is a content type that represents a video content
type VideoContent struct {
	Author string `json:"author"`
	Url    string `json:"url"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Format string `json:"format"`
}

// TextContent is a content type that represents a text content
type TextContent struct {
	Author      string `json:"author"`
	Url         string `json:"url"`
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"alttext"`
}

// Author is essentually a user in the system that access to publish or moderate
// contents.
type Author struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	email string `json:"email"`
}
