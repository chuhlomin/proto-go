package main

import (
	"log"
	"github.com/golang/protobuf/proto"
	"github.com/chuhlomin/proto-go/assets"
)

func main() {
	article := &assets.Article{
		Title: "TITLE",
		PrintData: &assets.Article_PrintData{
			Page: 0,
			Section: "NEWS",
		},
		Kicker: "BOOM",
	}

	bytes, err := proto.Marshal(article)
	if err != nil {
		log.Fatalf("Marshaling error: %v", err)
	}

	image := &assets.Image{};
	err = proto.Unmarshal(bytes, image);
	if err != nil {
		log.Fatalf("Unmarshaling error: %v", err)
		// Go error: proto: bad wiretype for field assets.Image.Width: got wiretype 2, want 0
	}

	log.Printf("Caption: %v", image.Caption)
	log.Printf("Author: %v", image.Author)
	log.Printf("Width: %v", image.Width)
}
