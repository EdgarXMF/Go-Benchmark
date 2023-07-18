FROM golang:1.20.5

WORKDIR /wd

COPY . .

RUN go build /wd/goab.go

CMD ["./goab","-c","10","-n","100","-k","https://example.com"]
