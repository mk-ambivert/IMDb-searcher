BINARY_NAME=imdb-searcher
PROGRAM_TYPE=cmd
OUTPUT_PATH=bin/${BINARY_NAME}
TEST_COVERAGE=coverage.out

export PROJECT_DIR=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

build_and_run: build run
	
build:
	go build -o ${OUTPUT_PATH} ${PROGRAM_TYPE}/${BINARY_NAME}/main.go

run:
	./${OUTPUT_PATH}
	
clean:
	rm -rf ./bin
	rm -rf ${TEST_COVERAGE}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=${TEST_COVERAGE}

dep:
	go mod download