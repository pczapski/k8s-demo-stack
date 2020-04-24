FROM golang:alpine as deps
RUN apk add --no-cache git ca-certificates
ADD go.* /test-app/
WORKDIR /test-app
RUN go mod download

FROM deps as builder
ARG version=dev
ENV version=${version}
ADD . /test-app
WORKDIR /test-app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-extldflags -w -s -X 'test-app/cmd.AppVersion=${version}'" -o app main.go


FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /test-app/app /app
ENTRYPOINT ["/app"]