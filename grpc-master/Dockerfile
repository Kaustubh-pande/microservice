# FROM debian:latest
# RUN mkdir /app
# WORKDIR /app
# COPY . /app
# #COPY proto /app
# CMD ["./app"]
FROM debian:latest
RUN mkdir /app
WORKDIR /app
#RUN apk add --no-cache ca-certificates
#COPY --from=builder /build/bin/server /app/
COPY main /app/
WORKDIR /app
CMD ["./main"]