FROM golang:1.12.1

WORKDIR $GOPATH/src/github.com/flexera-ryap/turtles-licenserec/license_rec_api

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["licenseapi_server"]