APP_NAME=gobackend
BUILD_DIR="./build/$(APP)"

run:
	nodemon --exec "go run" .

build:
	mkdir -p ./build && CGO_ENABLED=0 GOOS=linux go build -o ${BUILD_DIR}

test:
	go test -cover -v ./...