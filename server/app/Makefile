export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

binary_name = "www"
image_name = "training:www"

build:
	go build -o $(binary_name)

run:	build
	- ./$(binary_name)

test:
	go test -v

bechmark:
	go test -bench .

clean:
	rm -vf www training

docker-build:	build
	docker build -t $(image_name) .

docker-run:
	- docker run -it --network=host $(image_name)
