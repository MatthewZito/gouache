package entities

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// User represents a user object.
type User struct {
	Id       string `json:"id" dynamodbav:"id"`
	Username string `json:"username" dynamodbav:"username"`
	Password string `json:"password" dynamodbav:"password"`
}

// GetKey retrieves the primary key from a `User` object.
func (user User) GetKey() map[string]types.AttributeValue {
	username, err := attributevalue.Marshal(user.Username)
	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{"username": username}
}
