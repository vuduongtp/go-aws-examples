package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vuduongtp/go-sns-sqs-serverless/util/print"
	"github.com/vuduongtp/go-sns-sqs-serverless/util/sns"
)

// Event defines your lambda input/output data structure,
type Event struct {
	Payload string `json:"payload"`
}

// HandleRequest handles the incomming StepFunction request
func HandleRequest(e Event) (*Event, error) {
	topicARN := os.Getenv("MY_TOPIC_NAME")
	snsSvc := sns.New()
	output, err := snsSvc.SendToSNS(topicARN, e)
	if err != nil {
		return nil, err
	}
	print.PrettyPrint(output)

	return &Event{
		Payload: fmt.Sprintf("(Message:%s) is handled by Lambda function", e.Payload),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}

func runLocal() {
	topicARN := "PASTE_YOUR_TOPIC_ARN_HERE"
	event := Event{
		Payload: "123456",
	}
	snsSvc := sns.New()
	ouput, err := snsSvc.SendToSNS(topicARN, event)
	fmt.Println(err)
	fmt.Println(ouput)
}
