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

Application configuration is managed via environment variables that can be specified in a .env file. 
You can leave this file empty if you want to use the default settings. 
It's only necessary for satisfying the dependency of the docker-compose.yml file. 
If you want to modify the default env vars, make sure that this file is located in the same folder as your 
docker-compose.yml file:

```dotenv
# .env file

# These are all the configurable variables; you have the option to configure all of them or none at all.

# MySQL database configuration 
DATABASE_HOST=0.0.0.0
DATABASE_PORT=3307
DATABASE_USER=root
DATABASE_PASSWORD=root
DATABASE_SCHEMA=clinic
DATABASE_CHARSET=utf8
DATABASE_PARSE_TIME=true

# Gin configuration (execution mode)
GIN_MODE=debug
```

### Important:
If you intend to create an .env file and modify the DATABASE_PORT variable, please remember to update the
database port in the docker-compose file as well. This step is crucial because if you only modify the .env file,
the database container will still use the hardcoded port 3307.
(This applies also for DATABASE_PASSWORD and DATABASE_SCHEMA)

Make sure to customize the variables according to your requirements.

## Start the Application

Once you have built the application image and created the `.env` file, you can start the application with the 
following command:

```bash
make start
```

This command will use Docker Compose to start two containers: one for the MySQL database and another for the application. 
The application will be available at `http://localhost:8080`.

## Stop the Application

To stop the application and Docker containers, simply press `Ctrl + C` in the terminal where the `make start` 
command is running.