package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	entities "github.com/exbotanical/gouache/entities/reporting"
	"github.com/exbotanical/gouache/utils"
)

type ReportRepository struct {
	client   SQSSendMessageAPI
	queueUrl string
}

// NewReportRepository creates a new SQSApi object bearing the queue URL and underlying connection client.
func NewReportRepository() *ReportRepository {
	host := os.Getenv("SQS_HOST")
	port := os.Getenv("SQS_PORT")
	region := os.Getenv("SQS_REGION")
	queueName := os.Getenv("SQS_QUEUE_NAME")

	client, err := NewSQSClient(host, port, region)
	if err != nil {
		return nil
	}

	url := utils.ToEndpoint(host, port) + "/queue/" + queueName

	api := &ReportRepository{
		client:   client,
		queueUrl: url,
	}

	return api
}

// SendReport sends a Report to the SQS queue.
func (api *ReportRepository) SendReport(c context.Context, name string, data interface{}) (*sqs.SendMessageOutput, error) {
	// @todo handle err
	jsonData, _ := json.Marshal(data)

	input := &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{
			"name": {
				DataType:    aws.String("String"),
				StringValue: aws.String(name),
			},
			"caller": {
				DataType:    aws.String("String"),
				StringValue: aws.String("gouache/auth"),
			},
		},
		MessageBody: aws.String(string(jsonData)),
		QueueUrl:    &api.queueUrl,
	}

	return api.client.SendMessage(c, input)
}

func (api *ReportRepository) SendControllerErrorReport(r *http.Request, internal string, friendly string) {
	rl := entities.RequestReport{
		Path:   r.RequestURI,
		Method: r.Method,
		Error:  fmt.Sprintf("internal: %s, friendly: %s", internal, friendly),
	}

	api.SendReport(context.TODO(), entities.HTTP_HANDLER_EX, rl)
}
