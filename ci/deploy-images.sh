#!/bin/bash
docker build -t luskidotme/catalog-service:v1 ../src/catalog-service/
docker build -t luskidotme/presentation-service:v1 ../src/presentation-service/

docker push luskidotme/catalog-service:v1
docker push luskidotme/presentation-service:v1
