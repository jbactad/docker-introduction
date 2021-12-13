###########
# Demo #1 #
###########
docker-build-basic:
	cd basic && \
	GOOS=linux GOARCH=amd64 go build -o dist/hello-world main.go && \
	docker build -f Dockerfile .
# or docker build .

docker-build-tag-basic:
	cd basic && \
	docker build -t demo-basic .

docker-run-basic:
	docker run demo-basic

###########
# Demo #2 #
###########
docker-build-more:
	cd more && \
	go mod download && \
	GOOS=linux GOARCH=amd64 go build -o dist/hello-world main.go && \
	docker build -t demo-more .

docker-run-more:
	docker run -p=3000:3000 demo-more

###########
# Demo #3 #
###########

docker-build-multi-stage:
	cd multi-stage && \
	docker build -t demo-multi-stage .

docker-run-multi-stage:
	docker run -p=3000:3000 demo-multi-stage

###########
# Demo #4 #
###########

docker-run-from-registry:
	docker run --privileged -d \
	-p 6080:6080 -p 5554:5554 -p 5555:5555 \
	-e DEVICE="Samsung Galaxy S6" \
	--name android-container \
	budtmo/docker-android-x86-8.1

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