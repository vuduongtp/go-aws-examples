.PHONY: build clean deploy

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create todos/create/create.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/delete todos/delete/delete.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/get todos/get/get.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list todos/list/list.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/update todos/update/update.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock .serverless

deploy: clean format build
	sls deploy --verbose

format: 
	gofmt -w todos/create/create.go
	gofmt -w todos/delete/delete.go
	gofmt -w todos/get/get.go
	gofmt -w todos/list/list.go
	gofmt -w todos/update/update.go

remove: clean
	sls remove --verbose
