package misskey_twitter_post

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func init() {
	functions.HTTP("twitter", TwitterPostEntry)
}

func TwitterPostEntry(w http.ResponseWriter, r *http.Request) {

	misskeyRequest, err := ParseMisskeyRequest(r)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	err = PostToTwitter(misskeyRequest)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	log.Println("MisskeyRequest: ", misskeyRequest)
	log.Println("Host: ", r.Host)
	log.Printf("Header: %+v", r.Header)
}

func ParseMisskeyRequest(r *http.Request) (*MisskeyRequest, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Response Body Decode Error: ", err.Error())
		return nil, err
	}
	body := string(bodyBytes)
	log.Println("Response Body: ", body)

	misskeyRequest := &MisskeyRequest{}
	if err := json.Unmarshal([]byte(body), misskeyRequest); err != nil {
		log.Println("Response Body Decode Error: ", err.Error())
		return nil, err
	}
	log.Printf("MisskeyRequest: %+v", misskeyRequest)
	return misskeyRequest, nil
}

func PostToTwitter(misskeyRequest *MisskeyRequest) error {
	c, err := gotwi.NewClient(&gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("USER_TWITTER_OAUTH_ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("USER_TWITTER_OAUTH_ACCESS_TOKEN_SECRET"),
	})

	if err != nil {
		log.Println("Twitter Client Init Error: ", err.Error())
		return err
	}
	p := &types.CreateInput{
		Text: gotwi.String(misskeyRequest.Body.Note.Text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		log.Println("Twitter Post Error: ", err.Error())
		return err
	}
	log.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
	return nil
}
