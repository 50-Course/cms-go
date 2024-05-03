package content

import _ "github.com/golang/protobuf/ptypes/timestamp"

// Base Metadata
//
// Contains shared data relatable across all content types
type BaseMetadata struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// ImageContent is a content type that represents an image contentj
type ImageContent struct {
	Author   string       `json:"author"`
	Url      string       `json:"url"`
	AltText  string       `json:"alttext"`
	Metadata BaseMetadata `json:"metadata"`
}

// VideoContent is a content type that represents a video content
type VideoContent struct {
	Author   string       `json:"author"`
	Url      string       `json:"url"`
	Format   string       `json:"format"`
	Metadata BaseMetadata `json:"metadata"`
}

// TextContent is a content type that represents a text content
type TextContent struct {
	Author      string       `json:"author"`
	Url         string       `json:"url"`
	Description string       `json:"alttext"`
	Metadata    BaseMetadata `json:"metadata"`
}

// Author is essentually a user in the system that access to publish or moderate
// contents.
type Author struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	email string `json:"email"`
}
