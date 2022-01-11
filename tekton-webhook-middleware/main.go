package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type ExtractedWebhook struct {
	Commit     string `json:"commit"`
	Branch     string `json:"branch"`
	RepoName   string `json:"repo-name"`
	RepoRegion string `json:"repo-region"`
}

func handler(event events.CodeCommitEvent) (string, error) {

	extractedWebhook := &ExtractedWebhook{
		Commit:     event.Records[0].CodeCommit.References[0].Commit,
		Branch:     event.Records[0].CodeCommit.References[0].Ref,
		RepoName:   event.Records[0].CustomData,
		RepoRegion: event.Records[0].AWSRegion,
	}

	webhookPayload, err := json.Marshal(extractedWebhook)
	if err != nil {
		return "Can't parse embeded webhook", err
	}

	httpClient := &http.Client{Timeout: 10 * time.Second}

	fmt.Println(os.Getenv("TEKTON_WEBHOOK_URL"))

	httpRequest, err := http.NewRequest(http.MethodPost, os.Getenv("TEKTON_WEBHOOK_URL"), bytes.NewBuffer(webhookPayload))
	if err != nil {
		return "Can't construct request", err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return "Can't submit request", err
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("error")
	}
	fmt.Println("Content", string(respData))

	return "success", nil

}

func main() {
	lambda.Start(handler)
}
