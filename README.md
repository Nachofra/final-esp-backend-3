# Final ESP Backend 3 Group 1

- Franco Rampazzo
- Tomas Bernardin
- Gonzalo Sanchez
- Ignacio Francisco Mosca

## Prerequisites

Make sure you have the following prerequisites installed before running the application:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Important - Read Before Starting Application with Make](#important)

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
DATABASE_HOST=localhost # or host.docker.internal (1st is for local, 2nd is while running with make option)
DATABASE_PORT=3307
DATABASE_USER=root
DATABASE_PASSWORD=root
DATABASE_SCHEMA=clinic
DATABASE_CHARSET=utf8
DATABASE_PARSE_TIME=true

# Gin configuration (execution mode)
GIN_MODE=debug

# Custom host and port for your application
HOST=locahost # or 0.0.0.0 (1st is for local, 2nd is while running with make option)
PORT=8080
```

### Important:
Firstly, if you decide to proceed with the 'make' option to launch the app, it should function correctly 
as long as there is no conflicting service utilizing port 3307 for the database.

I'm not entirely sure why Docker Compose only works as expected when the application's HOST is set to '0.0.0.0' 
and the DATABASE_HOST is set to 'host.docker.internal'

If you have the application in Docker, and you're using Postman, please consider using '0.0.0.0' or '127.0.0.1' instead of 'localhost' in the requests.
Unfortunately, due to time constraints, I'm unable to thoroughly investigate this issue before the deadline. 

While you cannot directly override the hardcoded variables in the Docker Compose file with an environment file 
(unless you're feeling adventurous), I recommend first attempting to start the app using the 'make' command. 
If that doesn't work, you can try running the Golang application locally, where 'localhost' should function as expected.

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


## Things to improve (never)

Implementation of the validator: I feel that using it every time in each handler is a bit awkward. I would consider a way to generalize its functionality. Also, its configuration and initialization (currently, everything is hard-coded).

Deactivate dentists or patients and logically cancel their appointments. Currently, this causes conflicts, and you have to delete the appointments.

More validations (there's always room for more validation).

Improve the check for existence. In the updates, I have to retrieve the row by ID to check it. It's not very performant, but perhaps there are better or more generalized ways to do it. Doing it manually in the handler is not my preference, I think.

Thats all :)))