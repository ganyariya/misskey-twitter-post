package misskey_twitter_post

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

const (
	X_MISSKEY_HOOK_SECRET_HEADER_KEY = "X-Misskey-Hook-Secret"
)

func init() {
	functions.HTTP("twitter", TwitterPostEntry)
}

func TwitterPostEntry(w http.ResponseWriter, r *http.Request) {
	if err := ValidateMisskeyHookSecret(r); err != nil {
		log.Println("Error: ", err.Error())
		return
	}

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
}

func ValidateMisskeyHookSecret(r *http.Request) error {
	log.Printf("Header: %+v", r.Header)
	actualHookSecret := r.Header.Get(X_MISSKEY_HOOK_SECRET_HEADER_KEY)
	expectedHookSecret := os.Getenv("MISSKEY_HOOK_SECRET")
	if actualHookSecret != expectedHookSecret {
		return fmt.Errorf("misskey Hook Secret Error: %s is invalid", actualHookSecret)
	}
	return nil
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
