package repositories

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/exbotanical/gouache/entities"
	"github.com/exbotanical/gouache/models"
	"github.com/google/uuid"
)

// GetUser - given a primary key `username` - retrieves a `User` from dynamodb.
func (t UserTable) GetUser(username string) (*entities.User, error) {
	response, err := t.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(t.TableName),
		Key: map[string]types.AttributeValue{
			"username": &types.AttributeValueMemberS{Value: username},
		},
	})

	if err != nil {
		return nil, err
	}

	user := &entities.User{}

	if err = attributevalue.UnmarshalMap(response.Item, &user); err != nil {
		return nil, err
	}

	if user.Username == "" {
		return nil, fmt.Errorf("user with username %s not found", username)
	}

	return user, err
}

// CreateUser creates a new `User` in dynamodb.
func (t UserTable) CreateUser(userModel models.NewUserModel) error {
	user := entities.User{
		Id:       uuid.New().String(),
		Username: userModel.Username,
		Password: userModel.Password,
	}

	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	// @todo return id
	if _, err = t.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(t.TableName),
		Item:      item,
	}); err != nil {
		return err
	}

	return nil
}
