package misskey

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

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
