FROM golang:1.18-alpine
ENV CGO_ENABLED=0

WORKDIR /usr/src/app
COPY main.go .
COPY go.* .
RUN go build -o bouncer *.go

FROM alpine
COPY --from=0 /usr/src/app/bouncer /

CMD [ "./bouncer" ]
