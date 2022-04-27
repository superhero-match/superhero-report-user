prepare:
	go mod download
	go mod tidy

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

tests:
	go test ./... -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -func profile.cov

dkb:
	docker build -t superhero-report-user .

dkr:
	docker run -p "9000:9000" superhero-report-user

launch: dkb dkr

api-log:
	docker logs superhero-report-user -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it superhero-report-user /bin/bash

PHONY: prepare build tests dkb dkr launch api-log api-ssh rmc rmi clear