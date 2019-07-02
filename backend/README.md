# Desafio S3WF - BACKEND
Severino: Eduardo Figueiredo Gonçalves - @eduardofg87

## Outside Docker

### Export GO111MODULE
`export GO111MODULE=on`

### Cleaning all the trash
`rm go.mod go.sum`

### Initiate go mod
`go mod init`

### Creating vendor folder
`go mod vendor`

### Build the app
`go build`

### Run the app
`./godesafio-s3wf`


# Using Docker

## Build
`sudo docker build -t godesafio-s3wf .`

## RUN
`sudo docker run -p 8080:8080 godesafio-s3wf`