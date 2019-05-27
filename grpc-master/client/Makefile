build:
	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t grpc-master/client .

run:
	sudo docker run grpc-master/client