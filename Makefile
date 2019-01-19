build:
	GO111MODULE=on go mod download
	go build -o bin/pong

run: build
	bin/pong

test:
	GO111MODULE=on go mod download && go test ./... --cover
