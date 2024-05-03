package main

import (
	"fmt"
	"time"

	"github.com/50-Course/cms-go/content"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	painter := figure.NewFigure("Gance CMS Server", "", true)
	painter.Print()

	fmt.Println()
	fmt.Println("Hello, World! Starting server in a fews....")
	time.Sleep(5 * time.Second)
	fmt.Println("Server started!")
	fmt.Println("Server is running on port 8080")
	fmt.Println("Press Ctrl+C to stop the server")
	fmt.Println("Please open your terminal to see our juicy server running on your termnal")

	c := content.TextContent{
		Author:      "Idan!",
		Url:         "https://50-Course.showwcase.com/",
		Description: "My sweety portfolo pie",
		Metadata: content.BaseMetadata{
			ID:    1,
			Title: "My Portfolio",
		},
	}

	fmt.Printf("Got a response from the server: %v\n", c)
}
