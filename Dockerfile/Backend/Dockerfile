FROM golang:latest

ADD ./back-drop /backend

WORKDIR /backend
RUN go mod init project
RUN go mod tidy
RUN go build -o project .
RUN go test ./...
RUN chmod +x /backend

ENTRYPOINT ["/backend/project"]
