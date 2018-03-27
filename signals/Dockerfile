FROM golang:alpine
RUN apk add --no-cache ca-certificates git protobuf
RUN go get github.com/twitchtv/twirp/protoc-gen-twirp \
    github.com/golang/protobuf/protoc-gen-go
ADD signals.proto .
RUN protoc --go_out=. --twirp_out=. signals.proto
