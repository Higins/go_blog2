############################
# builder container
############################
FROM golang:1.16 AS builder
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin"
RUN chmod -R 777 "$GOPATH"

WORKDIR /goblog

COPY . .

RUN go mod vendor
RUN go build -o main main.go

############################
# app container
############################

FROM debian:bullseye

WORKDIR /goblog
COPY --from=builder /goblog/main .

CMD [ "./main" ]

EXPOSE 8089

