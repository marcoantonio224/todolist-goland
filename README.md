# todolist-goland
A simple CRUD todo list application that is powered by Docker and Go. Docker compose will run two containers **todolist** and **mysql**. Please keep in mind that the app will depend on the success of the database container as it *waits* for it to start (see [./sysops/entrypoint.sh](https://github.com/marcoantonio224/todolist-goland/blob/main/sysops/entrypoint.sh)). All of these configurations are found on the [docker-compose.yml](https://github.com/marcoantonio224/todolist-goland/blob/main/docker-compose.yml) file.
## Getting Started
In order to get this application up and running, please install the following technologies:
  ### Docker
  - [Docker](https://docs.docker.com/get-docker/)
  ### Docker Compose
    brew install docker-compose

  ### Go
    brew install golang

  ### gRPC
  Server side services will be powered by gRPC and protobufs.
    [Here](https://grpc.io/docs/protoc-installation/) is a link to install protobuf compiler


## MySQL
  Docker will bootstrap **mysql** (version `8.0.23`) for usage. The database will be called `todolist`, but feel free to change the actual name if desired. Please keep in mind, if changes are applied, to reflect those changes inside the [docker-compose.yml](https://github.com/marcoantonio224/todolist-goland/blob/main/docker-compose.yml) file.

## Golang
  If you are new to Go, please configure your `Go workspace` ( *bin*, *pkg*, & *src*) **first**. Here is a link to help you get started: [local go](https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04). Once you have those set up, clone this repository inside of your `src` folder.

## gRPC
In order to generate the compiled protobuf files necessary for this application (assuming you have downloaded the protocol buffer compiler) run this command on the root level of this directory

     protoc protobuf/todo.proto --go_out=plugins=grpc:.
This will then generate a compiled go package labled as `todo.pb.go` inside of `protobuf` directory.

## Environmental Variables
  In order to successfully run this application, you need to configure your `.env` file first (as docker will use some of its variables). There is a template of this already [here](https://github.com/marcoantonio224/todolist-goland/blob/main/.env.template). Copy this template in your real `.env` file and fill in the password of your choosing. You can opt in to the default mysql password **root**.

## Running the application
  Once you have properly configured everything, execute the following command on the root level of this app via terminal:
  ### Build the containers (app & database)
    docker compose build

  ### Run the app
    docker compose up

  ### It's alive!
  Application is hosted on [localhost:8080](http://localhost:8080/).

  ## Trouble shooting
  If you are coming across errors, please make sure to remove your docker volumes and re-build everything! (I know it is tedious).
