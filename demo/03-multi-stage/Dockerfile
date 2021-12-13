###############
# Build stage #
###############
FROM golang:1.17-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN GOOS=linux GOARCH=amd64 go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o dist/hello-world


#################
# Runtime stage #
#################
FROM alpine

ENV PERSON_NAME="Grocery Discovery"

EXPOSE 3000

COPY --from=build /app/dist/hello-world /bin/hello-world

CMD [ "hello-world" ]