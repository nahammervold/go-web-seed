# go-web-seed
Boiler plate project for go(golang) web api's

## Prerequisites
   
   [Docker](https://www.docker.com/get-docker) used to encapsulate application in a predictable environment
   
   [Go](https://golang.org/dl/) language used to build target executable
   
   
   (Optional) Make contains all the commands for development/release that could be ran with make
   

## Usage

### Build
Builds the executable for alpine linux then packages it in alpine linux docker container
```
make build
```

### Run
Will first build then run the docker container at port specified in Makefile (3000 default)
```
make run
```

### Clean
Careful this will remove ALL docker containers and images
```
make clean
```

## Docker Info
Uses latest alpine image