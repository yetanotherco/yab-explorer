# yab-explorer

The `yab-explorer` project is responsible for serving an API to fetch transactions done by users in [Yet Another Bridge](https://yetanotherbridge.com/).

## Prerequisites

* [Go 1.22.0](https://go.dev/dl/)

## Getting Started

The project is written in [Go](https://go.dev). In order to install all the needed dependencies you must run:

```sh
make deps
```

### Environment Variables

The following environment variables are required to run the code:

```dotenv
API_PORT=<api_port>
LOGGING_LEVEL=<DEBUG|TRACE|INFO>
POSTGRES_HOST=<postgres_host>
POSTGRES_USER=<postgres_user>
POSTGRES_PASSWORD=<postgres_password>
POSTGRES_DATABASE=<postgres_database>
POSTGRES_PORT=<postgres_port>
```

The following table describes each environment variable:

| Variable          | Description                                                                                                                                              |
|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|
| API_PORT          | Specifies the port on which the API server listens for incoming requests. If not provided, the default value is 8080.                                                                               |
| LOGGING_LEVEL     | Determines the level of detail in the logs produced by the application. Allowed values are DEBUG, TRACE or INFO. If not provided, the default value is INFO.                                               |
| POSTGRES_HOST     | The hostname or IP address of the PostgreSQL database.                                                                                                   |
| POSTGRES_USER     | The username used to authenticate with the PostgreSQL database.                                                                                          |
| POSTGRES_PASSWORD | The password associated with the POSTGRES_USER for authenticating with the PostgreSQL database.                                                          |
| POSTGRES_DATABASE | The name of the PostgreSQL database.                                                                                                                     |
| POSTGRES_PORT     | The port on which the PostgreSQL server is listening for connections. If not provided, the default value is 5432. |

Run the following command to create a `.env` file based on `.env.example` with the required environment variables:

```sh
make create_env
```

This will create a `.env` file with the required environment variables. Make sure to update the values of the variables
in the `.env` file to match your environment.

### Building the code

In order to build the code you must run:

```sh
make build
```

This will [clean](#cleaning-the-code) the project (removing any previous executable file) and build the project. An executable file called:
```
yab-explorer
```
will be created.

### Running the code

In order to run the code you must run:

```sh
make run
```

This will [build](#building-the-code) the code and then run the executable file that is created.

### Cleaning the code

In order to clean the project you must run:

```sh
make clean
```

This will remove the executable file created when the project is [built](#building-the-code).