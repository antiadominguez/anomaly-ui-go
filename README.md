# EdgeX custom application service template

An EdgeX custom application service is an application service that creates a function pipeline from code. Pipelines are usually a combination of standard EdgeX functions and custom functions. This template provides a starting point for creating a custom application service.

If you want to build an EdgeX application service that only uses standard EdgeX functions, you better use [EdgeX App Service Configurable](https://github.com/edgexfoundry/app-service-configurable), an executable that reads a configuration file and creates a function pipeline from it, without the need for coding.

## Run EdgeX

To run this device service  you first need to launch the rest of EdgeX services. This device has been tested using the latest LTS version of EdgeX: v3.1.0. Do the following steps to run EdgeX using docker-compose tool (docker and docker-compose must be installed):
```shell
curl https://raw.githubusercontent.com/edgexfoundry/edgex-compose/v3.1.0/docker-compose-no-secty.yml -o docker-compose.yml
sudo usermod -aG docker $USER
docker compose -f docker-compose.yml up -d
```

## Run device service

This template was create following the steps of [EdgeX custom app template](https://github.com/edgexfoundry/app-functions-sdk-go/tree/v3.1.0/app-service-template).

This template reads data from core using Trigger function, generates a new message back to the core data and print it. It's possible to publish in xml or json.

### In hybrid mode in Ubuntu 22.04

To use this template, you need to:

 * Install go from [golang.org](https://golang.org/doc/install) or using apt-get:

```shell
apt-get update
apt-get install golang-1.21
```

 * Clone this repository

 * Open the directory:

 ```shell
 cd edgex-app-template
 ```

 * Compile the app service:

```shell
make build
```

 * Deactivate security in this service using EdgeX environment variables:

```shell
export EDGEX_SECURITY_SECRET_STORE=false
```

 * Run the app service with yml config:

```shell
./edgex-app-template -cp -d
``` 

### Using Docker

The repository has a Dockerfile that can be used to build a Docker image that installs the EdgeX SDK and all other needed dependencies and runs the device service. 

Build the Docker image with:

```shell
docker build -t edgex-app-template .
```

Run the Docker image with:

```shell
docker run --network=host edgex-app-template -d -cp configuration.yaml
```

## Stop and restart EdgeX

To stop and cleanup Edge X we have to stop their containers and remove their associated images and volumes:

```shell
docker compose down -v
```

To fully restart EdgeX use the script scripts/restart_edgex.sh in this repository that stops, cleans up and brings up again a fresh EdgeX instance:

```shell
sh scripts/restart_edgex.sh
```