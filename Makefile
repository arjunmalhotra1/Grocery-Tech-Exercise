export GO111MODULE=on

run:
	go run cmd/Grocery-Tech-Exercise/main.go

build:
	go build cmd/Grocery-Tech-Exercise/main.go

docker-build:
	docker build -t store .

docker-run:
	docker run -it -p 8086:8086 store

docker-compose-build:
	docker-compose build

docker-compose-up:
	docker-compose up

unit-tests:
	go test -v ./...

api-tests:
	pytest