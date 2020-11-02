FROM golang:alpine
WORKDIR /app
COPY ./ /app
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
EXPOSE 8000
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main