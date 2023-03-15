package misskey_twitter_post

import (
	"fmt"
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("hello", Hello)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	io.WriteString(w, "Hello!")
}
