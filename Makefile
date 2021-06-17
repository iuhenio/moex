go-build:
	GOOS=linux go build

docker-build:
	docker build . -t moex:v0.2.0

go-clear:
	rm -f ./moex

docker-clear: SHELL := /bin/bash

docker-clear:
	docker rmi `docker images | grep moex | awk '{ print $$3 }' | head -n 1` --force

compose-up: SHELL := /bin/bash

compose-up:
	[ ! -d ~/projects_data/moex ] && mkdir -p ~/projects_data/moex || exit 0
	docker compose up -d 

compose-down:
	docker compose down moex

launch: go-build docker-build go-clear compose-up

clear: compose-down docker-clear