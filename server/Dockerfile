FROM golang:1.13-alpine AS go-builder
COPY app /opt/app
WORKDIR /opt/app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -o www


FROM wireguard:alpine
COPY run.sh .
COPY init-wireguard.sh .
COPY --from=go-builder /opt/app/www .

ENTRYPOINT ["./run.sh"]
