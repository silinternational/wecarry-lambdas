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
	log.Println("WeCarry API Maintenance started at", now.Format(time.RFC1123Z))

	url := os.Getenv("SERVICE_INTEGRATION_URL") + "/service"
	log.Println("SERVICE_INTEGRATION_URL =", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("failed to create new Request,", err)
		return nil
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("SERVICE_INTEGRATION_TOKEN")))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println("error making HTTP request,", err)
		return nil
	}

	if response.StatusCode >= 300 {
		log.Println("unexpected HTTP response code,", response.Status)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("error reading response body,", err)
		return nil
	}
	log.Println("response body:", string(responseBytes))

	return nil
}
