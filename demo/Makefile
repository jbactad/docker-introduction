# Demo #1 Basics
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

# Demo #2
docker-build-ports:
	cd ports && \
	go mod download && \
	GOOS=linux GOARCH=amd64 go build -o dist/hello-world main.go && \
	docker build -t demo-ports .

docker-run-ports:
	docker run -p=3000:3000 demo-ports

# Demo #3
docker-build-multi-stage:
	cd multi-stage && \
	docker build -t demo-multi-stage .

docker-run-multi-stage:
	docker run -p=3000:3000 demo-multi-stage
