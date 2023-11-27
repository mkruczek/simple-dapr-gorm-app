#!/bin/bash

components_path="./components"

dapr run --app-id inMemoryDB --components-path "$components_path" go run ./inMemoryDB/main.go &

dapr run --app-id sub --components-path "$components_path" go run ./subscriber/main.go &

dapr run --app-id pub --components-path "$components_path" go run ./publisher/main.go &

dapr run --app-id auditProduct --components-path "$components_path" go run ./auditProduct/main.go &

wait