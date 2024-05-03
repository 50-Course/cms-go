package content

import (
	"fmt"

	_ "github.com/golang/protobuf/ptypes/timestamp"
)

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

func (c TextContent) GetContentID() int64 {
	return c.Metadata.ID
}

func (c ImageContent) GetContentID() int64 {
	return c.Metadata.ID
}

func (c VideoContent) GetContentID() int64 {
	return c.Metadata.ID
}

// Author is essentually a user in the system that access to publish or moderate
// contents.
type Author struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	email string `json:"email"`
}

// =======================================================
// 	Content Service
// TODO:  Implement the ContentService in the api/ package

type Content interface {
	// GetContentType returns the type of content
	GetContentType() string
	// GetContent returns the content of the given id
	GetContent(id int64) (interface{}, error)
	// CreateContent creates a new content and returns the id of the created content
	CreateContent(content interface{}) (int64, error)
	// UpdateContent updates the content of the given id
	UpdateContent(content interface{}) error
	// Removes a given content from the platform given the id
	DeleteContent(id int64) error
}

// TODO: Move to api/ package
// ContentService is a service that provides content related operations
//
// It implements the Content interface, therefore providing us a way to
// utlize this service in API Contracts and Testing layers (Such as integration testing)
type ContentService struct {
	Content
}

func ProcessContent(c Content) {
	fmt.Printf("Please wait! Analysing  content...")
}
