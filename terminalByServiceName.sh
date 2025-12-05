#!/bin/bash

SERVICE_NAME=$1

NODEIP=$(docker node inspect "$(docker service ps $SERVICE_NAME --format '{{.Node}}')" --format '{{ .Status.Addr }}')

CONTAINERID=$(ssh "$NODEIP" "docker ps -aqf name=$SERVICE_NAME")

ssh -t "$NODEIP" "docker exec -it $CONTAINERID bash"
