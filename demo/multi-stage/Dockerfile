FROM golang:1.17-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o dist/hello-world

FROM scratch

EXPOSE 3000

COPY --from=build /app/dist/hello-world /bin/hello-world

CMD [ "hello-world" ]