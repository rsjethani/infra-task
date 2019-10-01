FROM wireguard:alpine


COPY www .
COPY run.sh .
COPY init-wireguard.sh .

ENTRYPOINT ["./run.sh"]
