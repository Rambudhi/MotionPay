run:
	go run cmd/main.go

build:
	go build -o bin/MotionPay cmd/main.go

test:
	go test ./...