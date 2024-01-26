#!/bin/bash

components_path="./components"

dapr run --app-id inMemoryDB --components-path "$components_path" go run ./inMemoryDB/main.go &
dapr run --app-id auditProduct --dapr-http-port 35001 --components-path "$components_path" go run ./auditProduct/main.go &
sleep 5
dapr run --app-id sub --log-level debug --components-path "$components_path" go run ./subscriber/main.go &
sleep 5
dapr run --app-id pub --app-port 6005 --components-path "$components_path" go run ./publisher/main.go &

wait