build:
	#protoc -I. --go_out=plugins=micro:. proto/service.proto
	# # --micro_out=proto/
	protoc --proto_path=proto \
	  -I. --micro_out=proto \
	 --go_out=plugins=micro:proto service.proto

	GOOS=linux GOARCH=amd64 go build
	sudo docker build -t grpc-master .
run: 
	#sudo docker run -p 50051:50051 grpc-master
	sudo docker run -p 50051:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 \
         grpc-master	