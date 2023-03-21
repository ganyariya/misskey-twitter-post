package strings

import (
	"golang.org/x/text/width"
)

const (
	URL_SPACE                = 1
	TWITTER_HALF_WIDTH_LIMIT = 280
	TWITTER_URL_LENGTH       = 23
)

func KindToSize(k width.Kind) int {
	switch k {
	case width.Neutral:
		return 1
	case width.EastAsianAmbiguous:
		return 2
	case width.EastAsianWide:
		return 2
	case width.EastAsianNarrow:
		return 1
	case width.EastAsianFullwidth:
		return 2
	case width.EastAsianHalfwidth:
		return 1
	}
	return 2
}

func TextSplitIntoTweets(text string, tweetLengthLimit int, urlSize int) []string {
	characters := []rune(text)
	splitTweetTexts := []string{}

	for i, j := 0, 0; i < len(characters); i = j {
		s := 0
		for j < len(characters) && s+KindToSize(width.LookupRune(characters[j]).Kind())+URL_SPACE+urlSize <= tweetLengthLimit {
			s += KindToSize(width.LookupRune(characters[j]).Kind())
			j++
		}
		splitTweetTexts = append(splitTweetTexts, string(characters[i:j]))
	}
	return splitTweetTexts
}
