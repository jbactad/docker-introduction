###########
# Demo #1 #
###########
docker-build-basic:
	cd 01-basic && \
	GOOS=linux GOARCH=amd64 go build -o dist/hello-world main.go && \
	docker build -f Dockerfile .
# or docker build .

docker-build-tag-basic:
	cd 01-basic && \
	docker build -t demo-basic .

docker-run-basic:
	docker run demo-basic

###########
# Demo #2 #
###########
docker-build-more:
	cd 02-more && \
	go mod download && \
	GOOS=linux GOARCH=amd64 go build -o dist/hello-world main.go && \
	docker build -t demo-more .

docker-run-more:
	docker run -p=3000:3000 demo-more

###########
# Demo #3 #
###########

docker-build-multi-stage:
	cd 03-multi-stage && \
	docker build -t demo-multi-stage .

docker-run-multi-stage:
	docker run -p=3000:3000 demo-multi-stage

###########
# Demo #4 #
###########

docker-run-from-registry:
	docker run -p3000:3000 -d joshjim27/docker-introduction-presentation


###########
# Demo #5 #
###########

docker-compose-build:
	docker-compose build

docker-compose-start:
	docker-compose up

docker-compose-startd:
	docker-compose up -d

docker-compose-ps:
	docker-compose ps

docker-compose-stop:
	docker-compose down