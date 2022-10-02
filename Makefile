run:
	@LOG_LEVEL=debug bash -c 'go run main.go --aliens-count 2'
test:
	@go test -v ./... -coverprofile cover.out
godoc:
	@godoc -http=0.0.0.0:6060 -v -timestamps=true -links=true -play=true
lint:
	@golangci-lint run -v --config ./.golangci.yml