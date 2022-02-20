#!/bin/bash

if [ "$1" ]; then

    ssh "$1" 'cd /srv/docker-alpine && git pull && sudo docker-compose -f docker-compose.yml up -d --build --remove-orphans'

    status=$(ssh "$1" sudo docker ps --format '{{.Status}}' --filter name=docker-alpine-01);
    if [ "$status" == "" ]; then
        status="failed"
    fi

    echo $status

else
    echo "Укажите пользователя и адрес сервера: '. deploy.sh user@example.com'"
fi
