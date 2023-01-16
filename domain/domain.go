package domain

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const defaultTimeout = time.Second * 10

func RunTask(url, task string) error {
	requestBody := `{"task":"` + task + `"}`
	request, err := http.NewRequest("POST", url, bytes.NewBufferString(requestBody))
	if err != nil {
		return errors.New("failed to create new Request, " + err.Error())
	}

	request.Header.Set("Authorization", "Bearer "+os.Getenv("SERVICE_INTEGRATION_TOKEN"))

	client := &http.Client{
		Timeout: defaultTimeout,
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
