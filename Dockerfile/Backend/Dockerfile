FROM golang:latest

ADD . /backend

WORKDIR /backend
RUN chmod +x /backend
EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/backend/project"]
