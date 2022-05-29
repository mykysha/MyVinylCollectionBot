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
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u github.com/spf13/viper
	go get -u github.com/volatiletech/sqlboiler/v4
	go get -u github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
	go get -u github.com/xuri/excelize/v2

gomod-tidy:
	go mod tidy -go=1.18 -compat=1.18
