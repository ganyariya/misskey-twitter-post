package misskey

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	X_MISSKEY_HOOK_SECRET_HEADER_KEY = "X-Misskey-Hook-Secret"
)

func ValidateMisskeyHookSecret(r *http.Request) error {
	log.Printf("Header: %+v", r.Header)
	actualHookSecret := r.Header.Get(X_MISSKEY_HOOK_SECRET_HEADER_KEY)
	expectedHookSecret := os.Getenv("MISSKEY_HOOK_SECRET")
	if actualHookSecret != expectedHookSecret {
		return fmt.Errorf("misskey Hook Secret Error: %s is invalid", actualHookSecret)
	}
	return nil
}
