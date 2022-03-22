#!/bin/bash

docker volume rm $(docker volume ls -q)
sudo rm -rf dbdata2
sudo rm -rf mongodb_data
