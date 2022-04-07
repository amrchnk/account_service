goose up:
	goose postgres "user=postgres port=5436 password=postgres dbname=account_service sslmode=disable" up

goose down:
	goose postgres "user=postgres port=5436 password=postgres dbname=account_service sslmode=disable" down

docker:
	docker-compose up