# precize_test

### Run a new postgres container

```bash
docker run --name precize-test -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=sat_resultDB -d postgres
```

### Start the container

```bash
docker start precize-test
```

### Running the Golang application

```bash
go run main.go
```
