package twitter

import (
	"context"
	"log"
	"os"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

func PostToTwitter(tweetText string, base64MediaIds []string) error {
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
		Text: gotwi.String(tweetText),
		Media: &types.CreateInputMedia{
			MediaIDs: base64MediaIds,
		},
	}

	res, err := managetweet.Create(context.Background(), c, p)
	if err != nil {
		log.Println("Twitter Post Error: ", err.Error())
		return err
	}
	log.Printf("[%s] %s\n", gotwi.StringValue(res.Data.ID), gotwi.StringValue(res.Data.Text))
	return nil
}
