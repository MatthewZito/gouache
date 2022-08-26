package repositories

import (
	"context"
	"fmt"
	"os"

	"github.com/exbotanical/gouache/utils"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// UserTable represents a dynamodb client connection to a given table `TableName`.
type UserTable struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

// InitUserTable initializes a new dynamodb client connection and `UserTable` object.
func InitUserTable() (*UserTable, error) {
	accessKey := os.Getenv("DYNAMO_ACCESS_KEY")
	secretKey := os.Getenv("DYNAMO_SECRET_KEY")
	host := os.Getenv("DYNAMO_HOST")
	port := os.Getenv("DYNAMO_PORT")
	region := os.Getenv("DYNAMO_REGION")
	tableName := os.Getenv("DYNAMO_TABLE_NAME")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),

		config.WithRegion(region),

		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: utils.ToEndpoint(host, port)}, nil
			})),
	)

	if err != nil {
		return nil, fmt.Errorf("cannot connect to dynamodb; see %v", err)
	}

	t := UserTable{TableName: tableName,
		DynamoDbClient: dynamodb.NewFromConfig(cfg)}

	return &t, nil
}
