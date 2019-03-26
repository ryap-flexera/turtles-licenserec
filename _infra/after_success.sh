#!/bin/bash

docker --version  # document the version travis is using
pip install --user awscli # install aws cli w/o sudo
export PATH=$PATH:$HOME/.local/bin # put aws in the path
pwd
cd $GOPATH/src/beacon/front_api_service
export BRANCH_LOWERCASE=`echo $TRAVIS_BRANCH | tr '[:upper:]' '[:lower:]'`
docker build -t frontapi_service-$BRANCH_LOWERCASE-$TRAVIS_BUILD_NUMBER .
docker image ls
docker tag frontapi_service-$BRANCH_LOWERCASE-$TRAVIS_BUILD_NUMBER:latest 754533604493.dkr.ecr.ap-southeast-2.amazonaws.com/flexera/beacon-front-api-service:latest $(aws ecr get-login --no-include-email --region ap-southeast-2)
docker push 754533604493.dkr.ecr.ap-southeast-2.amazonaws.com/flexera/beacon-front-api-service:latest
