# Waracle Cakes

Mono repo for two services, the first being `cake_service` which is a Cake CRUD
API service, and `cakes_web` which is a web user interface for managing cakes,
which interfaces with the API

## Design overview

The applications have been split, as its generally good to split concerns, so
the API deals with data processing, and the Web app deals with user interfaces.

This structure is also known as microservices, where each application is independent with a focused single purpose.

This allows different technologies and languages to be used, so teams can focus on what they know.

1. Cake Service API

This is a Go application using Gin.

2. Cake Web App

This is a Javascript application using React.

3. Database

This is MongoDB.

4. Web Server (docker-compose only)

This is nginx, only exposed service which routes requests to services.

## Usage

This project is intended to be run using docker and docker-compose, but you may
spin up the services locally but you will have to install and setup
dependencies.

### Quick start

Make sure you have `docker` and `docker-compose` installed.

```
docker-compose up
```

Once booted, you can access the web interface at `http://localhost`

### Futher reading

The two services `cakes_web` and `cake_service` have their own READMEs
expanding on what they are built with and the tooling the make use of.
