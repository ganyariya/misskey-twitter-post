package entry

import (
	"log"
	"net/http"

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

	err = twitter.PostToTwitter(misskeyRequest)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
}
