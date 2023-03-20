package image

import (
	"encoding/base64"
	"io"
	"log"
	"net/http"
)

func DLImageToBase64(imageUrl string) (string, error) {
	log.Println("try decode: ", imageUrl)
	if imageUrl == "" {
		return "", nil
	}

	response, err := http.Get(imageUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	imageBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	base64Data := base64.StdEncoding.EncodeToString(imageBytes)
	log.Println("decoded: ", base64Data)
	return base64Data, nil
}
