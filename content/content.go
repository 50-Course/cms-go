package content

import _ "github.com/golang/protobuf/ptypes/timestamp"

type ImageContent struct {
	Author  string `json:"author"`
	Url     string `json:"url"`
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	AltText string `json:"alttext"`
}

type VideoContent struct {
	Author string `json:"author"`
	Url    string `json:"url"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Format string `json:"format"`
}

type TextContent struct {
	Author      string `json:"author"`
	Url         string `json:"url"`
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"alttext"`
}

type Author struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	email string `json:"email"`
}
