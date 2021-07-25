FROM golang:alpine as builder
RUN apk add --no-cache ca-certificates make tzdata
COPY . /app
RUN cd /app && \
    go get -d -v && \
    go mod download && \
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o stock_service .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
VOLUME /data
COPY --from=builder /app/stock_service /app/
RUN apk add --no-cache ca-certificates tzdata tini
USER nobody
EXPOSE 80
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/stock_service"]
