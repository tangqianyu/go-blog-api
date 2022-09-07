FROM golang:alpine

ENV TZ=Asia/Shanghai

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ADD . .

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
