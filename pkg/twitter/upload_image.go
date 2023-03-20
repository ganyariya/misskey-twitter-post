package twitter

import (
	"log"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func UploadMediasToTwitter(base64Datas []string) ([]string, error) {
	client := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("USER_TWITTER_OAUTH_ACCESS_TOKEN"),
		os.Getenv("USER_TWITTER_OAUTH_ACCESS_TOKEN_SECRET"),
		os.Getenv("GOTWI_API_KEY"),
		os.Getenv("GOTWI_API_KEY_SECRET"),
	)

	mediaIds := []string{}
	for _, data := range base64Datas {
		id, err := client.UploadMedia(data)
		if err != nil {
			log.Println("Anaconda Upload MediaError: ", err.Error())
			return []string{}, err
		}
		mediaIds = append(mediaIds, id.MediaIDString)
	}

	return mediaIds, nil
}
