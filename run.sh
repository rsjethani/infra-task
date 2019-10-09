#!/bin/bash

set -e

export ALFA_USER="$1"
export ALFA_KEY="$2"

terraform init

terraform apply -auto-approve

sudo pip install ansinv
./generate_inventory.py

ansible-playbook -i inventory site.yml -v

ssh -i $ALFA_KEY ${ALFA_USER}@`terraform output ip` "cd infra-task && docker-compose up"
