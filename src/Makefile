OUT_DIR=out
LINT_VERSION=v1.51.0

# https://pkg.go.dev/github.com/golangci/golangci-lint
LINT = github.com/golangci/golangci-lint/cmd/golangci-lint@${LINT_VERSION}

# must be first
default: clean dep build lint

# clean - removes object files from package src directories
clean:
	go clean
	rm -rf ./${OUT_DIR}

# install dependencies
dep:
	go mod download
	go mod tidy

# build project
build: clean
	GOARCH=amd64 GOOS=darwin go build -tags="" -o ${OUT_DIR}/darwin/ ./...
	GOARCH=amd64 GOOS=windows go build -tags="" -o ${OUT_DIR}/windows/ ./...
	GOARCH=amd64 GOOS=linux go build -tags="" -o ${OUT_DIR}/linux/ ./...

# run project
run:
	go run ./...

# run tests
test:
	go test ./...

# run lint check
lint: 
	go run ${LINT} run

lint-tool:
	golangci-lint run

# generate swag docs - https://github.com/swaggo/swag & https://pkg.go.dev/github.com/swaggo/swag@v1.8.10
# TODO: Implement Swaggo, also need to add to "build" command (ex: `build: clean swag`)
# swag:	
# 	go run github.com/swaggo/swag@v1.8.10 fmt
# 	go run github.com/swaggo/swag@v1.8.10 init -g ./assets/docs/swagger --parseDependency --parseInternal -parseDepth 1 --dir ./api,./model -q
# 	Could also add a generate.go file -> then add line: go generate ./assets/docs/...
