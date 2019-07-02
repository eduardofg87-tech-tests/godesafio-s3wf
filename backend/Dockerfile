# build stage
FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/godesafio-s3wf /app/
#export go.png
#COPY go.png /app/
#COPY tdc_logo.png /app/
EXPOSE 8080
ENTRYPOINT ["/app/godesafio-s3wf"]