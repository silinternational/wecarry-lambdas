package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
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

	tasks := []string{"file_cleanup", "token_cleanup", "location_cleanup"}
	var errs []error
	for _, t := range tasks {
		url := os.Getenv("SERVICE_INTEGRATION_URL") + "/service"
		log.Println("running task: " + t)
		if err := runTask(url, t); err != nil {
			log.Println(err.Error())
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("error(s): %v", errs)
	}
	return nil
}

func runTask(url, task string) error {
	requestBody := `{"task":"` + task + `"}`
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(requestBody))
	if err != nil {
		return errors.New("failed to create new Request, " + err.Error())
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("SERVICE_INTEGRATION_TOKEN"))

	client := &http.Client{
		Timeout: ApiTimeout,
	}
	response, err := client.Do(request)
	if err != nil {
		return errors.New("error making HTTP request to " + url + ", " + err.Error())
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("error reading response body, " + err.Error())
	}

	if response.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP response, status code = %s, request url %s, request body %s, response body %s",
			response.Status, url, requestBody, responseBytes)
	}

	if len(responseBytes) > 0 {
		log.Printf("finished task %s, response body: %s", task, responseBytes)
	} else {
		log.Printf("finished task %s", task)
	}

	return nil
}
