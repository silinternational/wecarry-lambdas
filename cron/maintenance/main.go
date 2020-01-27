package main

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaConfig struct {
	ConfigPath string
}

func main() {
	lambda.Start(handler)
}

func handler(lambdaConfig LambdaConfig) error {
	// Log to stdout and remove leading date/time stamps from each log entry (Cloudwatch Logs will add these)
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	now := time.Now().UTC()
	log.Printf("WeCarry API Maintenance started at %s", now.Format(time.RFC1123Z))

	return nil
}
