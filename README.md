# Final ESP Backend 3 Group 1

- Franco Rampazzo
- Tomas Bernardin
- Gonzalo Sanchez
- Ignacio Francisco Mosca

## Prerequisites

Make sure you have the following prerequisites installed before running the application:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Build the Application Image

To build the application image, execute the following command:

```bash
make build
```

This will build a Docker image for the application from the Dockerfile located in the `docker/` folder.

## Configuration

Application configuration is done through environment variables that you can define in a `.env` file. 
Ensure that this file is in the same folder as your `docker-compose.yml` file and contains the following variables:

```dotenv
# .env file

# MySQL database configuration 
# (Default values are shown below)
DATABASE_HOST=0.0.0.0
DATABASE_PORT=3307
DATABASE_USER=root
DATABASE_PASSWORD=root
DATABASE_SCHEMA=clinic

# Gin configuration (execution mode)
GIN_MODE=debug
```

Make sure to customize the variables according to your requirements.

You can modify the parameters of the MySQL container in the Docker Compose file if needed. 
The Docker Compose file currently has hardcoded values for the root user, database schema, and ports.

## Start the Application

Once you have built the application image and created the `.env` file (it can be empty, just to fulfill 
the docker-compose dependency), you can start the application with the following command:

```bash
make start
```

This command will use Docker Compose to start two containers: one for the MySQL database and another for the application. 
The application will be available at `http://localhost:8080`.

## Stop the Application

To stop the application and Docker containers, simply press `Ctrl + C` in the terminal where the `make start` 
command is running.