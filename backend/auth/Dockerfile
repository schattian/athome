FROM golang:alpine AS builder

LABEL maintainer="Sebastián Chamena <sebachamena@gmail.com>"

WORKDIR /auth

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./auth .

FROM scratch

WORKDIR /auth

COPY --from=builder /auth/auth ./auth

EXPOSE 9990

CMD ["/auth/auth"]
