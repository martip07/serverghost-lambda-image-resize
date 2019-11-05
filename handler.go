package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() (string, error) {
	return "Howdy from serverghost lambda image resize!", nil
}

func main() {
	lambda.Start(handleRequest)
}
