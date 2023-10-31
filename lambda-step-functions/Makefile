.PHONY: clean deploy

clean:
	rm -rf ./bin ./vendor Gopkg.lock .serverless

deploy: clean format
	sls deploy --verbose

format: 
	gofmt -w .

remove: clean
	sls remove --verbose
