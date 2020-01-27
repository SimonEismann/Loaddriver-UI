# LoadDriver-UI

This project aims to provide a more comfortable and user-friendly way to run load tests on web applications using the [HTTP Load Generator](https://github.com/joakimkistowski/HTTP-Load-Generator).

# Run the application using Docker Compose
Prerequisites:
- Docker
- Docker-Compose

The application can be started as follows:
  1. Navigate to the root of this project
  2. Run `docker-compose up -d`

The LoadDriver-UI can now be accessed under http://localhost/ or - if available - through your configured DNS.

# Run the application using Kubernetes
Prerequisites:
- Running Kubernetes cluster
- Configured `kubectl` installation

The application can be deployed to your k8s cluster as follows:
  1. Adjust the variable values in the `manifests/loaddriver-all.yaml` file:
     1. `<API_LOCATION>` should point to the location where the `loaddriver-master` instance can be accessed by the frontend
     2. (Optional) If you use Ingresses in your cluster, change the `<HOST_URL>` variable to the configured DNS
  2. Run `kubectl apply -f manifests/`

# Configure user profiles
The HTTP Load Generator uses Lua scripts to define the user profiles for the load tests.
The docker images are built with two example scripts which can also be found in this project under `loaddriver-master/scripts/`.
The `example_script` script is a very simple example user profile, that just visits the different pages of the frontend of this application itself.
The `teastore_buy` script is a more elaborate script, that shows how to perform more advanced actions like `POST` requests.
If you plan to use the `teastore_buy` script, you will need a running [TeaStore](https://github.com/descartesresearch/teastore) installation.