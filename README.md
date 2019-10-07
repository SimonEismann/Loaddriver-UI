# LoadDriver-UI

This project aims to provide a more comfortable and user-friendly way to run load tests using the [HTTP Load Generator](https://github.com/joakimkistowski/HTTP-Load-Generator).
The HTTP Load Generator can be used to test web applications.

# Running the application **locally**
The exmple setup uses Docker Compose and besides the LoadDriver UI spins up a full TeaStore which is used to run load against.
Before starting load experiments, make sure the TeaStore is up and running first.

Prerequisites are:
- Docker
- Docker-Compose

The example setup can now be started as follows:
  1. docker-compose up -d

The LoadDriver-UI can now be accessed under http://localhost.