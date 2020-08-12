FROM golang:1.15 as build
ENV GO111MODULE=on

WORKDIR /go/src/github.com/k-yomo/dig-api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

FROM alpine
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/k-yomo/dig-api/bin/server ./server
EXPOSE 1323
ENTRYPOINT ["./server"]
