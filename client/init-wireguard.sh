#!/bin/bash

if [ -z "$WG_IFACE_CIDR" ]; then
	echo "WG_IFACE_CIDR environment variable not given"
	exit 1
fi

ip link add wg0 type wireguard

ip addr add "$WG_IFACE_CIDR" dev wg0

wg setconf wg0 /etc/wireguard/wg0.conf

ip link set wg0 up

