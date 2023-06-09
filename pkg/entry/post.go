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

	base64FileDatas := []string{}
	for _, url := range misskeyRequest.GetFileUrls() {
		data, err := image.DLImageToBase64(url)
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}
		base64FileDatas = append(base64FileDatas, data)
	}

	base64MediaIds, err := twitter.UploadMediasToTwitter(base64FileDatas)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	tweetText := misskeyRequest.BuildTweetText(os.Getenv("MISSKEY_DOMAIN"))
	err = twitter.PostToTwitter(tweetText, base64MediaIds)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
}
