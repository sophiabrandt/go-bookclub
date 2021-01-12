# Development environment
# -----------------------
FROM golang:1.15-alpine
ENV CGO_ENABLED 0

WORKDIR /usr/src/app

# create user
ARG USER_ID=100001
ENV USER_ID $USER_ID
ARG GROUP_ID=100001
ENV GROUP_ID $GROUP_ID

# add non-root user and give permissions to workdir
RUN addgroup --gid $GROUP_ID user && \
          adduser user --ingroup user --gecos '' --disabled-password --uid $USER_ID && \
          chown -R user:user /usr/src/app

RUN set -eux \
  && apk add --no-cache --virtual .build-deps \
    gcc musl-dev entr

USER user

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && chmod +x /usr/src/app/bin/*
