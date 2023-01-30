FROM golang:alpine

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/myclient

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Build Go binary
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go mod tidy
RUN go build .

ENTRYPOINT ["/go/src/myclient/client"]
CMD []
