FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod ./
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o build/task ./cmd/main.go

FROM alpine:latest

COPY --from=build /app/build/* /opt/
COPY test_file.txt /opt/

ENTRYPOINT [ "/opt/task", "/opt/test_file.txt"]
CMD [ "run" ]