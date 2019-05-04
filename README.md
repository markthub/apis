# Form APIs

The APIs are very simple and they are being written in Go using the following:

- [Gin](https://github.com/gin-gonic/gin) as HTTP Framework
- MongoDB for database

## How to run it locally

If you use `docker-compose` then it is super easy.

```bash
$: docker-compose up -d
```

The previous command will build the Docker Images and then start the containers.
Check if both services are up and running with `docker ps -a`. You should see something like the following

```bash
CONTAINER ID        IMAGE               COMMAND                    CREATED              STATUS              PORTS                      NAMES
8040785fa4c9        martkhub_api        "./bin/markthub server"     2 seconds ago       Up 1 second         0.0.0.0:8000->8000/tcp     server
c62001cddecf        mysql:5.6           "docker-entrypoint.s…"      2 days ago          Up 2 days           0.0.0.0:3306->3306/tcp     mysql
```

If you want to stop, run the following command `docker-compose down`. You can also build the image from scratch but in that case you need to make sure that the dependencies for the API (the Go stuff) are already present in your machine (see the next section).

## Manually

**I will assume that you have installed Go (with `dep` installed as well) and that you are able to run any go program within you PATH.**

Open a new terminal and go to the folder `apis/api` and do the following:

1. `dep ensure`. This should install all the dependencies
2. `go run main.go server`

This should be sufficient to make the APIs working. Open a browser and go to `http://localhost:8000/api/version`. You should see something like the following

```json
{
    "version": "MarktHub APIs v0.0.1"
}
```

## Test

You can use the `Makefile` in the `/api` folder and run `GOCACHE=off make test`.

## Docs

In the current situation there are only few endpoints available considering the scope of the project. In fact, you can see those by running in the browser the following `http://{ip-docker-machine}:8000/api/docs` (if running via docker-compose) or `http://localhost:8000/api/docs` (if running by using `go run main.go server`). The API code is designed to be extendend with eventually other endpoints.