ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine

ARG TZ

RUN apk update \
    && apk upgrade \
    && apk add --no-cache bash \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone

RUN apk add build-base
RUN apk add git

WORKDIR /srv/golang-debugging-pages

COPY ./config ./config
COPY ./public ./public
COPY ./templates ./templates
COPY ./main.go .
COPY ./go.mod .
COPY ./go.sum .
COPY ./.env .

COPY ./docker-entrypoint.sh .
RUN chmod 755 /srv/golang-debugging-pages/docker-entrypoint.sh

RUN go build main.go

ENTRYPOINT ["/srv/golang-debugging-pages/docker-entrypoint.sh"]
