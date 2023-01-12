package main

import (
	"fmt"
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
	log.Println("WeCarry API Maintenance started at", now.Format(time.RFC1123Z))

	url := os.Getenv("SERVICE_INTEGRATION_URL") + "/service"
	tasks := []string{"file_cleanup", "token_cleanup", "location_cleanup"}
	var errs []error
	for _, t := range tasks {
		log.Println("running task: " + t)
		if err := domain.RunTask(url, t, ApiTimeout); err != nil {
			log.Println(err.Error())
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("error(s): %v", errs)
	}
	return nil
}
