FROM golang:alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add --no-cache git
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct


# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/helloworld

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git

# Build Go binary
COPY Makefile go.mod go.sum ./
RUN make init && go mod download 
COPY . .
RUN make proto tidy build

# Deployment container
FROM scratch

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/helloworld/helloworld /helloworld
ENTRYPOINT ["/helloworld"]
CMD []
