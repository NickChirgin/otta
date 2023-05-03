FROM golang:1.20.1-alpine

WORKDIR /app

RUN apk add git curl wget upx protoc libc6-compat && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 

COPY . .

RUN go mod tidy

RUN protoc --go_out=. proto/tinyurl.proto --go-grpc_out=. proto/tinyurl.proto

CMD go run cmd/main.go