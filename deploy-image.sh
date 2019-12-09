#!/usr/bin/env bash
set -e
ORG=swarmpit
REPO=agent
NAMESPACE=$ORG/$REPO
TAG=${TRAVIS_BRANCH/\//-}

if [ $TRAVIS_PULL_REQUEST == "false" ]
then
    if [ $TAG == "master" ]
    then
        TAG="latest"
    fi
else
    TAG="pr-$TRAVIS_PULL_REQUEST-$TAG"
fi

IMAGE=$NAMESPACE:$TAG
TAGS=($IMAGE-amd64 $IMAGE-arm64 $IMAGE-armv7 $IMAGE-armv6)

docker build --build-arg GOARCH=amd64 -t ${TAGS[0]} .
docker build --build-arg GOARCH=arm64 -t ${TAGS[1]} .
docker build --build-arg GOARCH=arm --build-arg GOARM=7 -t ${TAGS[2]} .
docker build --build-arg GOARCH=arm --build-arg GOARM=6 -t ${TAGS[3]} .

if [ $TRAVIS_SECURE_ENV_VARS == "true" ]
then
    docker login -u "${DOCKER_USERNAME}" -p "${DOCKER_PASSWORD}"
    docker push $NAMESPACE
    docker manifest create $IMAGE ${TAGS[*]}
    docker manifest annotate --os linux --arch amd64 $IMAGE ${TAGS[0]}
    docker manifest annotate --os linux --arch arm64 $IMAGE ${TAGS[1]}
    docker manifest annotate --os linux --arch arm --variant v7 $IMAGE ${TAGS[2]}
    docker manifest annotate --os linux --arch arm --variant v6 $IMAGE ${TAGS[3]}
    docker manifest push $IMAGE
    docker run --rm lumir/remove-dockerhub-tag --user "${DOCKER_USERNAME}" --password "${DOCKER_PASSWORD}" ${TAGS[*]}
else
	echo "images can be pushed only from base repo"
fi
