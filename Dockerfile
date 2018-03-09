FROM golang:1.9.2 as go-builder

ENV DEP_VERSION v0.4.1
ENV REPOSITORY_PATH /go/src/github.com/RisingStack/almandite-user-service

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/${DEP_VERSION}/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p ${REPOSITORY_PATH}
WORKDIR ${REPOSITORY_PATH}

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

COPY . ${REPOSITORY_PATH}/

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN go build -o server ./cmd/almandite-user-server/main.go
RUN mkdir /app && cp server /app/server

FROM alpine:3.7
COPY --from=go-builder /app /app
WORKDIR /app
CMD ["/app/server"]