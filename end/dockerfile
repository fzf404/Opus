FROM golang

MAINTAINER "nmdfzf404@163.com"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
	
WORKDIR /opt/opus-go

COPY . /opt/opus-go 

RUN go build .

EXPOSE 8080

CMD ["./Opus"]
