PORT = 3000
CONTAINERS = $(shell docker ps -a -q)
IMAGES = $(shell docker images -q)

build_target:
	GOOS=linux GOARCH=amd64 go build

build: build_target
	docker build --no-cache -t go-web-seed .

run: build
	docker run -p $(PORT):3000 -t go-web-seed

kill:
ifeq ($(CONTAINERS),)
    $(info No Containers to remove)
else
	docker rm -f $(CONTAINERS)
endif

clean: kill
ifeq ($(IMAGES),)
    $(info No Images to remove)
else
	docker rmi -f $(IMAGES)
endif