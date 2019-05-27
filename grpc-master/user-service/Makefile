build:
	protoc --proto_path=proto \
	  -I. --micro_out=proto \
	 --go_out=plugins=micro:proto user.proto

	GOOS=linux GOARCH=amd64 go build