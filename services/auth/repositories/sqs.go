package repositories

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/exbotanical/gouache/utils"
)

// SQSSendMessageAPI defines the interface for the SendMessage function and
// affords the opportunity to test using a mocked service.
type SQSSendMessageAPI interface {
	SendMessage(ctx context.Context,
		params *sqs.SendMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

// NewSQSClient creates a new SQS Client connection object.
func NewSQSClient(host string, port string, region string) (SQSSendMessageAPI, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),

		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: utils.ToEndpoint(host, port)}, nil
			})),
	)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to sqs; see %v", err)
	}

	client := sqs.NewFromConfig(cfg)

	return client, nil
}
