#! /bin/bash

sudo docker-compose -f ./db/docker-compose.yml down
sudo systemctl stop docker.socket
sudo systemctl stop docker
