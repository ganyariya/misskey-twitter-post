
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

