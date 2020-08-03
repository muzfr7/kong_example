# KONG EXAMPLE

#### Create a Docker network
```
$ docker network create kong-net
```

#### Start a PostgreSQL container by executing
```
$ docker run -d --name kong-database \
    --network=kong-net \
    -p 5432:5432 \
    -e "POSTGRES_DB=kong" \
    -e "POSTGRES_USER=kong" \
    -e "POSTGRES_PASSWORD=kong" \
    postgres:9.6
```

#### Prepare your database
> Run the database migrations with an ephemeral Kong container..
```
$ docker run --rm \
    --network=kong-net \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    kong:latest kong migrations bootstrap
```

#### Start Kong
> When the migrations have run and your database is ready, start a Kong container that will connect to your database container, just like the ephemeral migrations container..
```
$ docker run -d --name kong \
    --network=kong-net \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 127.0.0.1:8001:8001 \
    -p 127.0.0.1:8444:8444 \
    kong:latest
```

#### Use Kong
http://localhost:8001/
http://localhost:8001/services

---

#### Build and create container for your app

> Build an image from Dockerfile
```
$ docker build . -t kong_example_app
```

> Create container from `kong_example_app` image created above
```
$ docker run -p 3000:3000 --name kong_example_app -dit kong_example_app
```
