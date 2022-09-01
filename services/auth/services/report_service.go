package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	entities "github.com/exbotanical/gouache/entities/reporting"
	"github.com/exbotanical/gouache/repositories"
)

// ReportService defines the contract for a `ReportServiceProvider` instance.
type ReportService interface {
	SendReport(c context.Context, name string, data interface{}) (*sqs.SendMessageOutput, error)
	SendControllerErrorReport(r *http.Request, internal string, friendly string)
}

// ReportServiceProvider is a service layer API for interacting with reporting data.
type ReportServiceProvider struct {
	*repositories.ReportRepository
}

// NewReportService creates a new `ReportService`.
func NewReportService() (ReportService, error) {
	r, err := repositories.NewReportRepository()
	if err != nil {
		return nil, err
	}

	s := &ReportServiceProvider{
		r,
	}

	return s, nil
}

// SendReport sends data to the SQS queue.
func (s *ReportServiceProvider) SendReport(c context.Context, name string, data interface{}) (*sqs.SendMessageOutput, error) {
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
		QueueUrl:    s.deriveQueueUrl(),
	}

	return s.Client.SendMessage(c, input)
}

// SendControllerErrorReport sends a RequestReport with metadata from the provided http.Request `r` to the gouache/reporting system.
func (s *ReportServiceProvider) SendControllerErrorReport(r *http.Request, internal string, friendly string) {
	rl := entities.RequestReport{
		Path:   r.RequestURI,
		Method: r.Method,
		Error:  fmt.Sprintf("internal: %s, friendly: %s", internal, friendly),
	}

	go func() {
		s.SendReport(context.TODO(), entities.HTTP_HANDLER_EX, rl)
	}()
}

// Derive from the ReportService `Url` and `QueueName` a Queue URL pointer, as required by the SQS client.
func (s *ReportServiceProvider) deriveQueueUrl() *string {
	queueUrl := s.Url + "/queue/" + s.QueueName

	return &queueUrl
}
