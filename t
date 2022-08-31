#!/usr/bin/env bash
SESSION='123'
AUTH_PORT=5000
REPORTING_PORT=6000
DYNAMO_PORT=8000

# add report
add_report() {
  local caller="${1:-'x'}"
  local data="${2:-'y'}"
  local name="${3:-'z'}"

  curl localhost:$REPORTING_PORT/api/report -d "{\"caller\": \"$caller\", \"data\": \"$data\", \"name\": \"$name\"}" -H 'Content-Type: application/json' --cookie "gouache_session=$SESSION" -v
}

# get report
get_report() {
  local id="$1"

  curl localhost:$REPORTING_PORT/api/report/$1 --cookie "gouache_session=$SESSION" -v
}

# get all reports
get_all_reports() {
  local last_page_key="$1"

  if [[ key -ne "" ]]; then
    curl localhost:$REPORTING_PORT/api/report?last_page_key=$last_page_key --cookie "gouache_session=$SESSION" -v
  else
    curl localhost:$REPORTING_PORT/api/report --cookie "gouache_session=$SESSION" -v
  fi
}

# register user
register() {
  local username="${1:-'user'}"
  local password="${2:-'password'}"

  curl localhost:$AUTH_PORT/auth/register -d "{\"username\":\"$username\",\"password\":\"$password\"}" -H "Content-Type: application/json"
}

# login user
login() {
  local username="${1:-'user'}"
  local password="${2:-'password'}"

  curl localhost:$AUTH_PORT/auth/login -d "{\"username\":\"$username\",\"password\":\"$password\"}" -H "Content-Type: application/json"
}

# cleanup resources and services
cleanup() {
  local container_id="${1:-$(docker container ls | awk 'NR > 1 {print $1; exit}')}"

  docker container stop $container_id
  sudo service redis-server stop
}

create_table () {
  local table_name="$1"
  local hash_key="$2"

  aws dynamodb --endpoint-url http://localhost:$DYNAMO_PORT --region us-east-1 create-table --table-name $table_name --attribute-definitions AttributeName=$hash_key,AttributeType=S --key-schema AttributeName=$hash_key,KeyType=HASH --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
}

# initialize resources and services needed for local targeting
main () {
  # start redis
  sudo service redis-server start

  # set session
  echo "SET "$SESSION" '{\"username\":\"user\",\"expiry\":\"9999-08-26T15:28:03.683Z\"}'" | redis-cli

	# set redis password
	echo "SET CONFIG requirepass password" | redis-cli

  # start local db
  docker run -p $DYNAMO_PORT:$DYNAMO_PORT amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb &

  # create tables
  create_table resource id
  create_table report id
  create_table user username
}

# stop here if being sourced
return 2>/dev/null

# stop on errors and unset variable refs
set -o errexit
set -o nounset

main $*
