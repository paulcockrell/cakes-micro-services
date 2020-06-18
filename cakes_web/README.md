# Cake web

Cake web is a React PWA for interfacing with the Cake service.

## Development

To develop run `yarn start` which will automatically compile changes and reload
the browser rendering the app. 

See the `package.json` file for further script actions.

The web app interfaces with the `Cake service` api, which should be running
along side this application for it to work. Visit the `Cake service` project README for instructions on executing it.

## Production

You have two options, one is to build the source on the host, or build a Docker image for live deployment to (for example) Kubernetes

We use a `Makefile` (found in the root of the project) to define these commands.

*Option 1* - local build
Compiled files built to `build` folder
```
make build
```

*Option 2* - Docker build
Docker image tagged with `cake-web:latest`
```
make docker
```

