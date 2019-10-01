FROM wireguard:alpine

RUN apk add curl

COPY run.sh .
COPY client.sh .
COPY init-wireguard.sh .

ENTRYPOINT ["./run.sh"]
