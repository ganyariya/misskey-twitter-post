package misskey_twitter_post

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func init() {
	functions.HTTP("twitter", TwitterPostMain)
}

func ParseMisskeyRequest(r *http.Request) (*MisskeyRequest, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Response Body Decode Error")
		return nil, err
	}
	body := string(bodyBytes)
	fmt.Println("Request Body:", body)

	misskeyRequest := &MisskeyRequest{}
	if err := json.Unmarshal([]byte(body), misskeyRequest); err != nil {
		fmt.Println("Misskey Body Decode Error")
		return nil, err
	}
	fmt.Printf("MisskeyRequest: %+v\n", misskeyRequest)
	return misskeyRequest, nil
}

func PostToTwitter(misskeyRequest *MisskeyRequest) error {
	c, err := gotwi.NewClient(&gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           os.Getenv("USER_TWITTER_OAUTH_ACCESS_TOKEN"),
		OAuthTokenSecret:     os.Getenv("USER_TWITTER_OAUTH_ACCESS_TOKEN_SECRET"),
	})

	if err != nil {
		fmt.Println("Twitter Client Init Error")
		return err
	}
	p := &types.CreateInput{
		Text: gotwi.String(misskeyRequest.Body.Note.Text),
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println("Twitter Post Error")
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
	return nil
}

func TwitterPostMain(w http.ResponseWriter, r *http.Request) {

	misskeyRequest, err := ParseMisskeyRequest(r)
	if err != nil {
		fmt.Printf("error %v", err.Error())
		return
	}

	err = PostToTwitter(misskeyRequest)
	if err != nil {
		fmt.Printf("error %v", err.Error())
		return
	}

	fmt.Println("MisskeyRequest", misskeyRequest)
	fmt.Println("Host", r.Host)
	fmt.Println("Header", r.Header)
}
