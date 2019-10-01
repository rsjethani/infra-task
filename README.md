# Alfaview Infra Task

## Running the setup
**Host system requirements:**
* docker(>=19.03.0)
* docker-compose(>=1.24)
* wireguard kernel module installed

**PS:** Tested on ubuntu 16.04
```
infra-task$ docker-compose up
```
Now you should see the two containers `server` and `client1` interacting with each other over wireguard vpn endpoints.

## The wireguard:alpine base image
This is the base image for both the client and the server. It simply installs the wireguard package on the alpine base image. **NOTE:** This still requires wireguard kernel module to be present on the host OS.

## The server:alpine image
This image is built as a mutli-stage docker build, where in the first stage we compile the `server/app` on a golang:1.13-alpine base then copy the resulting app binary `www` to the next stage. When started, the container first configures the wireguard interface via the `init-wireguard.sh` script and then simply starts the `www` app which listens **only on** wg0 IP and 3333 port. The app is simply what we did for backend task earlier. Example command to run:
```
infra-task$ docker run -it --name server --cap-add NET_ADMIN --mount type=bind,source=$PWD/server/wg0.conf,destination=/etc/wireguard/wg0.conf -e WG_IFACE_CIDR=10.0.0.5/24 server:alpine
```
The value of **WG_IFACE_CIDR** environment variable is used to configure the `wg0` interface.

## The client1:alpine image
This image, similar to the server container first configures the wg0 interface and then simply makes curl requests to the server container to get the nth prime no. every second. Example command to run:
```
infra-task$ docker run -it --name client1 --cap-add NET_ADMIN --mount type=bind,source=$PWD/client/wg0.conf,destination=/etc/wireguard/wg0.conf -e WG_IFACE_CIDR=10.0.0.11/24 -e SERVER_APP_ADDR=10.0.0.5:3333 client1:alpine
```
The value of **WG_IFACE_CIDR** environment variable is used to configure the `wg0` interface.
The value of **SERVER_APP_ADDR** environment variable is used tell client app about the `<wireguard_ip>:<port>` of the server app.
