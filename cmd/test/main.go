package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "だよ！でお勉強をしているよ ganyariya.net すごいね https://ok.google なるほど"
	urlPattern := regexp.MustCompile(`https?://[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]+|([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}`)
	fmt.Println(urlPattern.FindAllStringIndex(text, -1))

	for i, r := range text {
		fmt.Println(i, string(r))
	}

}
