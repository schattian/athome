FROM golang:alpine AS builder

LABEL maintainer="Sebastián Chamena <sebachamena@gmail.com>"

WORKDIR /users

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./users .

FROM scratch

COPY --from=builder /users/users /users

EXPOSE 9991

CMD ["/users"]
