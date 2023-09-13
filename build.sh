# Remember to build your handler executable for Linux!
# When using the `provided.al2` runtime, the handler executable should be named `bootstrap`
GOOS=linux GOARCH=arm64 go build -o bootstrap main.go
zip lambda-handler.zip bootstrap
