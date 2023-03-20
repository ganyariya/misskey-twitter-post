package entry

import (
	"log"
	"net/http"
	"os"

	"github.com/ganyariya/misskey-twitter-post/pkg/image"
	"github.com/ganyariya/misskey-twitter-post/pkg/misskey"
	"github.com/ganyariya/misskey-twitter-post/pkg/twitter"
)

func TwitterPostEntry(w http.ResponseWriter, r *http.Request) {
	if err := misskey.ValidateMisskeyHookSecret(r); err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	misskeyRequest, err := misskey.ParseMisskeyRequest(r)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	fileBase64Datas := []string{}
	for _, url := range misskeyRequest.GetFileUrls() {
		data, err := image.DLImageToBase64(url)
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}
		fileBase64Datas = append(fileBase64Datas, data)
	}

	tweetText := misskeyRequest.BuildTweetText(os.Getenv("MISSKEY_DOMAIN"))
	err = twitter.PostToTwitter(tweetText)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
}
