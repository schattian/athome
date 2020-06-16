FROM golang:alpine AS builder

ARG svc_name

# Keys to fetch private deps of pb repo
ARG github_user
ENV github_user=$github_user
ARG github_token
ENV github_token=$github_token
RUN apk add git
RUN git config \
    --global \
    url."https://${github_user}:${github_token}@github.com".insteadOf \
    "https://github.com"
RUN go env -w GOPRIVATE=github.com/athomecomar/athome


LABEL maintainer="Sebastián Chamena <sebachamena@gmail.com>"

WORKDIR /${svc_name}

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./migrate ./migrations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./${svc_name} .

FROM scratch

ARG svc_name

WORKDIR /${svc_name}

COPY --from=builder /${svc_name}/${svc_name} ./${svc_name}
COPY --from=builder /${svc_name}/migrate ./migrate

COPY ./migrations/*.sql ./migrations/

CMD ["/messager/messager"]
