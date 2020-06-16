# Waracle Cake API

Extreemly basic in-memory cake CRUD api.

*Endpoints:*
* GET /cakes
* GET /cakes/{id}
* POST /cakes
* PUT /cakes/{id}
* DELETE /cakes/{id}

## Build
### Local build
Build a binary, requires Go being installed on your host machine
```
make build
```
_or_ 
### Docker build
If you have Docker installed, build the binary in an image
```
make docker
```

## Run
Either run the binary created from the `Local build` instruction:
```
./waracle-cake-service
```
_or_ if you built the Docker image, run that and map ports
```
docker run -p 8080:8080 waracle-cake-service:latest
```

## Usage examples

When either executing the binary, or running the docker image as described in
the previous steps, you can access the cake API using curl as shown in the
examples below.

Note: Two cakes are created by default when the application starts. Cakes are
stored in memory and not persisted between service restarts.

### Get all cakes
```
curl localhost:8080/cakes
```

### Get a cake
```
curl localhost:8080/cakes/1
```

### Create a cake
```
curl -XPOST -H 'Content-Type: application/json' \
     -d '{ "name": "Slimey Cake","comment": "Slippy","yum_factor": 5, "image_url": "/cake.pic.jpg" }' \
     http://localhost:8080/cakes
```

### Update a cake
```
curl -XPUT -H 'Content-Type: application/json' \
     -d '{ "name": "Triangle Donut","comment": "Amazing concept","yum_factor": 2, "image_url": "/triangle.pic.jpg"}' \
     http://localhost:8080/cakes/1
```

### Delete a cake
```
curl -XDELETE -H 'Content-Type: application/json' \
     http://localhost:8080/cakes/1
```
