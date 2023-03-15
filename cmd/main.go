package main

import (
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	_ "github.com/ganyariya/misskey-twitter-post"
)

func main() {
	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
