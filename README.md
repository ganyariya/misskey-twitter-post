
# Misskey Twitter Post

Cloud function to post Misskey notes to twitter.

```bash
# install
go mod tidy

# local test
FUNCTION_TARGET=twitter go run ./cmd/main.go

# deploy
gcloud functions deploy twitter --entry-point twitter  --trigger-http --runtime go120
```

## Link

- Cloud Function
  - https://cloud.google.com/functions/docs/concepts/go-runtime?hl=ja
  - https://cloud.google.com/functions/docs/configuring/env-var?hl=ja#gcloud
  - https://cloud.google.com/functions/docs/deploy?hl=ja
- Twitter
  - https://github.com/g8rswimmer/go-twitter/tree/master/v2#tweets
  - https://developer.twitter.com/en/portal/dashboard 

