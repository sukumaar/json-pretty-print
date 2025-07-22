FROM golang:1.24.5-alpine3.21 AS BUILDER

RUN apk add upx
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
RUN CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o json-pretty-print ./cmd/app
RUN upx json-pretty-print

FROM gcr.io/distroless/static-debian12
COPY --from=BUILDER /src/json-pretty-print /json-pretty-print
CMD ["/json-pretty-print"]