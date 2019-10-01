#!/bin/bash

set -e

./init-wireguard.sh


./server.sh -listen ${WG_IFACE_CIDR%/*}:3333
