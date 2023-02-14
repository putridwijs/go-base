# go-base

###

## Run Application:

```
go run cmd/main.go http
```


## Create Migration
```
docker compose run -v ../db/migrations:/migrations migrate create -ext sql -dir migrations <migration_name>
```
