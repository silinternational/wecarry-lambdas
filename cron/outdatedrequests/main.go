package main

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/silinternational/wecarry-lambdas/domain"
)

const ApiTimeout = 10 * time.Second

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
	log.Println("WeCarry API Email Notification for Outdated Requests started at", now.Format(time.RFC1123Z))

	url := os.Getenv("SERVICE_INTEGRATION_URL") + "/service"
	task := "outdated_requests"

	log.Println("running task: " + task)
	if err := domain.RunTask(url, task, ApiTimeout); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
