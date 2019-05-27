build:
	ls
	protoc -I. --go_out=plugins=micro:. \
	service.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t grpc-master .
run: 
	docker run -p 50051:50051 grpc-master