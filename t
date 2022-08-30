#!/usr/bin/env bash
SESSION='123'


# add report
add_report() {
  local caller="${1:-'x'}"
  local data="${2:-'y'}"
  local name="${3:-'z'}"

  curl localhost:6000/api/report -d "{\"caller\": \"$caller\", \"data\": \"$data\", \"name\": \"$name\"}" -H 'Content-Type: application/json' --cookie "gouache_session=$SESSION" -v
}

# get report
get_report() {
  local id="$1"

  curl localhost:6000/api/report/$1 --cookie "gouache_session=$SESSION" -v
}

# get all reports
get_all_reports() {
  local last_page_key="$1"

  if [[ key -ne "" ]]; then
    curl localhost:6000/api/report?last_page_key=$last_page_key --cookie "gouache_session=$SESSION" -v
  else
    curl localhost:6000/api/report --cookie "gouache_session=$SESSION" -v
  fi
}

# register user
register() {
  local username="${1:-'user'}"
  local password="${2:-'password'}"

  curl localhost:5000/auth/register -d "{\"username\":\"$username\",\"password\":\"$password\"}" -H "Content-Type: application/json"
}

# login user
login() {
  local username="${1:-'user'}"
  local password="${2:-'password'}"

  curl localhost:5000/auth/login -d "{\"username\":\"$username\",\"password\":\"$password\"}" -H "Content-Type: application/json"
}

# cleanup resources and services
cleanup() {
  local container_id="${1:-$(docker container ls | awk 'NR > 1 {print $1; exit}')}"

  docker container stop $container_id
  sudo service redis-server stop
}

# initialize resources and services needed for local targeting
main () {
  # start redis
  sudo service redis-server start

  # set session
  echo "SET "$SESSION" '{\"username\":\"user\",\"expiry\":\"9999-08-29 21:59:59.999999\"}'" | redis-cli

  # start local db
  docker run -p 8000:8000 amazon/dynamodb-local   -jar DynamoDBLocal.jar -sharedDb &

  # create tables
  aws dynamodb --endpoint-url http://localhost:8000 --region us-east-1 create-table --table-name resource  --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
  aws dynamodb --endpoint-url http://localhost:8000 --region us-east-1 create-table --table-name report  --attribute-definitions AttributeName=id,AttributeType=S --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
  aws dynamodb --endpoint-url http://localhost:8000 --region us-east-1 create-table --table-name user --attribute-definitions AttributeName=username,AttributeType=S --key-schema AttributeName=username,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
}

# stop here if being sourced
return 2>/dev/null

# stop on errors and unset variable refs
set -o errexit
set -o nounset

main $*
