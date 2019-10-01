#!/bin/bash

set -e

./init-wireguard.sh


./www -listen ${WG_IFACE_CIDR%/*}:3333
