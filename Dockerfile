FROM golang

RUN mkdir -p /home/app
COPY . /home/app
WORKDIR /home/app
RUN go build main.go
CMD ["./main"]
