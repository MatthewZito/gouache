package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/exbotanical/gouache/utils"
)

// ReportRepository holds a SQS client connection object and associated queue connection metadata.
type ReportRepository struct {
	Client    *sqs.Client
	Url       string
	QueueName string
}

// NewReportRepository initializes a new `ReportRepository` with internal SQS client.
func NewReportRepository() (*ReportRepository, error) {
	accessKey := os.Getenv("AWS_FAKE_ACCESS_KEY")
	secretKey := os.Getenv("AWS_FAKE_SECRET_KEY")
	host := os.Getenv("SQS_HOST")
	port := os.Getenv("SQS_PORT")
	region := os.Getenv("SQS_REGION")
	queueName := os.Getenv("SQS_QUEUE_NAME")

	url := utils.ToEndpoint(host, port)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),

		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),

		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: url}, nil
			})),
	)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to sqs; see %v", err)
	}

	client := sqs.NewFromConfig(cfg)

	return &ReportRepository{
		Client:    client,
		Url:       url,
		QueueName: queueName,
	}, nil
}
