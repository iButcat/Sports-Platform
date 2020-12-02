FROM golang:1.14.3-alpine AS build
RUN mkdir /backend
ADD . /backend
WORKDIR /backend
RUN go build -o main .
CMD ["/backend/main"]
