docker run -p 8000:8000 amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb

aws dynamodb --endpoint-url http://localhost:8000 --region us-east-1 create-table --table-name resource  --attribute-definitions AttributeName=Id,AttributeType=S --key-schema AttributeName=Id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
aws dynamodb --endpoint-url http://localhost:8000 --region us-east-1 create-table --table-name user  --attribute-definitions AttributeName=Username,AttributeType=S --key-schema AttributeName=Username,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

aws dynamodb --endpoint-url http://localhost:8000 --region us-east-1 list-tables

curl -d '{ }' -H "Content-Type: application/json" localhost:4000/resource -v

aws dynamodb --endpoint-url http://localhost:8000 create-table --table-name user --attribute-definitions AttributeName=Id,AttributeType=S --key-schema AttributeName=Id,KeyType=HASH --region us-east-1 --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

aws dynamodb --endpoint-url http://localhost:8000  put-item --table-name user --item file://u.json --region us-east-1

aws dynamodb --endpoint-url http://localhost:8000  scan --table-name user --region us-east-1

aws dynamodb --endpoint-url http://localhost:8000 query --region us-east-1 --table-name user --key-condition-expression "Username=:username" --expression-attribute-values file://expression_attributes.json
