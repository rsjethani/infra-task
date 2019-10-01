#!/bin/bash


if [ -z "$SERVER_APP_ADDR" ]; then
	echo "SERVER_APP_ADDR environment variable not given"
	exit 1
fi

while true; do
	cmd="curl -s http://${SERVER_APP_ADDR}/prime/$((RANDOM % 100))"
	echo "$cmd => `eval $cmd`"
	sleep 1
done
