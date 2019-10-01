# Alfaview Infra Task
## Running the setup
*Assumptions: The host system is running docker(>=19.03.0) and docker-compose(>=1.24)*
```
infra-task$ docker-compose up
```
Now you should see the two containers `server` and `client1` interacting with each other over wireguard vpn endpoints.
