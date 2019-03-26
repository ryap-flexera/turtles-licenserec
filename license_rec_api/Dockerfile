FROM golang:1.12.1

LABEL maintainer="XX XX <XX@XX.com>"

RUN mkdir -p beacon/front_api_service

WORKDIR turtles-licenserec/license_rec_api

COPY cmd/license_rec_api/license_rec_api  .

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["license_rec_api"]