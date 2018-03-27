FROM golang:alpine AS b0
RUN apk add --no-cache git ca-certificates
RUN go get github.com/gorilla/mux \
    github.com/llater/datatube/signals
COPY main.go .
RUN go build -v -o signals-server

FROM alpine:3.6
COPY --from=b0 /go/signals-server .
ENTRYPOINT ./signals-server
