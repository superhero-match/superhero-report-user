prepare:
	go mod download

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

dkb:
	docker build -t superhero-report-user .

dkr:
	docker run -p "9000:9000" -p "8270:8270" superhero-report-user

launch: dkb dkr

api-log:
	docker logs superhero-report-user -f

es-log:
	docker logs es -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it superhero-report-user /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build dkb dkr launch api-log es-log api-ssh es-ssh rmc rmi clear