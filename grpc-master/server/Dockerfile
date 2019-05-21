FROM debian:latest
RUN mkdir /app
WORKDIR /app
COPY ../grpc-master/server/ /app
COPY ../grpc-master/proto/. /app
CMD ["./app"]