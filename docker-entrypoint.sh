#!/bin/bash

config='/srv/golang-debugging-pages/config/main/config.json';
if [ ! -f $config ]; then
    cp /srv/golang-debugging-pages/config/main/config.example.json $config;
fi

sed -i "s/\"port\": 1999/\"port\": $CONTAINER_PORT/g" $config;
sed -i "s/\"debug\": true/\"debug\": $DEBUG/g" $config;
sed -i "s/\"delay\": 5/\"delay\": $DELAY/g" $config;

cat $config >> /srv/golang-debugging-pages/logs/main.log

/srv/golang-debugging-pages/main;
