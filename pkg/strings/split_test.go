package strings_test

import (
	"testing"

	"github.com/ganyariya/misskey-twitter-post/pkg/strings"
	"github.com/stretchr/testify/assert"
)

func TestTextSplitIntoTweets(t *testing.T) {
	tests := []struct {
		Text                    string
		TweetLengthLimit        int
		UrlSize                 int
		ExpectedSplitTweetTexts []string
	}{
		{
			"aaaa",
			5,
			0,
			[]string{"aaaa"},
		},
		{
			"",
			5,
			0,
			[]string{},
		},
		{
			"こんにちは",
			5,
			0,
			[]string{"こん", "にち", "は"},
		},
		{
			"こんにちは",
			5,
			1,
			[]string{"こ", "ん", "に", "ち", "は"},
		},
		{
			"こaんaにaちaは\n",
			5,
			1,
			[]string{"こa", "んa", "にa", "ちa", "は\n"},
		},
	}

	for _, tt := range tests {
		actualSplitTweetTexts := strings.TextSplitIntoTweets(tt.Text, tt.TweetLengthLimit, tt.UrlSize)
		assert.Equal(t, tt.ExpectedSplitTweetTexts, actualSplitTweetTexts)
	}
}
