package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	request, err := http.NewRequest("GET", os.Getenv("SERVICE_INTEGRATION_URL") + "/site/status", nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SERVICE_INTEGRATION_TOKEN")))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	if response.StatusCode >= 300 {
		log.Printf("unexpected HTTP response code: %s", response.Status)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error reading response body: %s", err.Error())
		return nil
	}
	log.Printf("response body: %s", string(responseBytes))
	
	return nil
}
