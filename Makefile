BINARY_NAME=imdb-searcher
PROGRAM_TYPE=cmd
TEST_COVERAGE=coverage.out

build_and_run: build run

build:
	go build -o bin/${BINARY_NAME} ${PROGRAM_TYPE}/${BINARY_NAME}/main.go

run:
	./bin/${BINARY_NAME}

clean:
	go clean
	rm -rf ./bin
	rm -rf ${TEST_COVERAGE}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=${TEST_COVERAGE}

dep:
	go mod download
