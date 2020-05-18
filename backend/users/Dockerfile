FROM golang:1

LABEL maintainer="Sebastián Chamena <sebachamena@gmail.com>"

ARG env=development

ENV ENV $env
WORKDIR /user

COPY pbgo ./pbgo
COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o user .

EXPOSE 9991
