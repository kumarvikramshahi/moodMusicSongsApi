#!/bin/bash

echo "Running mongo container"
docker run -d -p 27023:27017 --name mongo_moodMusic mongo

echo "running main.go"
go run main.go


# docker exec -it mongo_moodMusic /bin/bash
# mongosh