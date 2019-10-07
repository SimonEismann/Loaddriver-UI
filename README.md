# LoadDriver-UI

This project aims to provide a more comfortable and user-friendly way to run load tests using the [HTTP Load Generator](https://github.com/joakimkistowski/HTTP-Load-Generator).
The HTTP Load Generator can be used to test web applications.

# Running the application **locally**
The exmple setup uses Docker Compose and consists of two parts.
The first part starts up a full TeaStore which is used to run load against.
The second part runs the LoadDriver-UI, which can then be used to conduct the load experiments.

Prerequisites are:
- Docker
- Docker-Compose

The example setup can now be started as follows:
  1. docker network create load-driver-network
  2. docker-compose -f docker-compose-teastore.yaml up -d
  3. docker-compose up [-d] (don't worry about the warning about orphan containers)

The LoadDriver-UI can now be accessed under http://localhost.
