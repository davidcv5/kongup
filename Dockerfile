FROM golang:1.12 AS build
WORKDIR /kongup
COPY go.mod ./
COPY go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kongup

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /kongup .
ENTRYPOINT ["./kongup"]
