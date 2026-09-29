#! /bin/bash

status=$(sudo systemctl status docker | grep running)
if  [ "$status" == "" ]; then
  sudo systemctl start docker
fi
sudo docker-compose -f ./db/docker-compose.yml up -d

