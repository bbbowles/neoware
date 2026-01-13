#!/bin/bash

SERVICE_NAME=$1

NODEIP=$(docker node inspect "$(docker service ps $SERVICE_NAME --format '{{.Node}}' | head -n 1)" --format '{{ .Status.Addr }}')

CONTAINERID=$(ssh "$NODEIP" "docker ps -q --filter name=$SERVICE_NAME --filter status=running | head -n 1")

ssh -t "$NODEIP" "docker exec -it $CONTAINERID bash"
