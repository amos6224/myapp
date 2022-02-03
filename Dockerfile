FROM golang:1.9

COPY ./main.go /

RUN go build -o /main /main.go

EXPOSE 9000

ENTRYPOINT ["/main"]
