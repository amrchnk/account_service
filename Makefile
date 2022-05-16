build:
	docker-compose build

start:
	docker-compose up -d

create:
	protoc --proto_path=proto proto/*.proto --go_out=./
	protoc --proto_path=proto proto/*.proto --go-grpc_out=./

goose up:
	goose postgres "user=postgres port=5436 password=postgres dbname=account_service sslmode=disable" up

goose down:
	goose postgres "user=postgres port=5436 password=postgres dbname=account_service sslmode=disable" down

migrate up:
	goose postgres "user=postgres port=5432 password=postgres dbname=account_service sslmode=disable" up

migrate down:
	goose postgres "user=postgres port=5432 password=postgres dbname=account_service sslmode=disable" down