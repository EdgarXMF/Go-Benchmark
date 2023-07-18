FROM golang:1.20.5

WORKDIR /wd

COPY . .

RUN go build /wd/goserver.go

EXPOSE 443 80

CMD ["./goserver"]
