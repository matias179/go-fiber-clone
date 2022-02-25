FROM golang:1.17-alpine as builder

WORKDIR /etc/go-fiber-clone

COPY . .

RUN apk update && apk upgrade

RUN go mod download
RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o main ./cmd/main.go

FROM alpine as release

EXPOSE 8080

COPY --from=builder /etc/go-fiber-clone/main /bin
COPY ./dotenv/ /bin/dotenv

CMD [ "/bin/main" ]
