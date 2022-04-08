cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm bot_app

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start vinyl-bot

docker-stop:
	docker container stop vinyl-bot

update-gomod:
	go get -u go get github.com/spf13/viper
	make gomod-tidy

download-gomod:
	go get go get github.com/spf13/viper

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
