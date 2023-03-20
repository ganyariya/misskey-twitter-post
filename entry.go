package misskey_twitter_post

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/ganyariya/misskey-twitter-post/pkg/entry"
)

func init() {
	functions.HTTP("twitter", entry.TwitterPostEntry)
}
