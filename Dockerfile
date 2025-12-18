FROM golang:1.25.4 AS build

ENV GOPROXY=https://goproxy.cn
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOARCH=amd64 GOOS=linux go build -o leave-a-message main.go

FROM ubuntu:24.04

WORKDIR /app
COPY --from=build /app/leave-a-message /app/leave-a-message
COPY templates ./templates/
COPY static ./static/
CMD ["/app/leave-a-message"]