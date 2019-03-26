FROM golang:1.12.1

LABEL maintainer="XX XX <XX@XX.com>"

RUN mkdir -p turtles-licenserec/license_rec_api

WORKDIR $GOPATH/src/turtles-licenserec/license_rec_api

COPY cmd/license_rec_api/licenseapi_server  .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["license_rec_api"]