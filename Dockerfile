FROM golang:1.19
RUN apt-get install -y ca-certificates
WORKDIR /opt/hello-gotuna
COPY . .
CMD ["go", "run", "examples/fullapp/cmd/main.go"]
