FROM golang:1.14.3-alpine AS build
WORKDIR ./
COPY . .
RUN go build -o /main.go .
FROM scratch AS bin
COPY --from=build /main
