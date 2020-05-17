FROM golang:1

LABEL maintainer="Sebastián Chamena <sebachamena@gmail.com>"

ARG env=development

ENV ENV $env
WORKDIR /user

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o user .

EXPOSE 9991

# CMD if [ ${ENV} = development ]; \
# 	then \
#     go get github.com/pilu/fresh && \
# 	fresh; \
# 	else \
# 	./user; \
# 	fi


# CMD ./user
