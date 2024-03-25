# YoFio Go Test

Prueba técnica en Go de asignación de créditos

## 1. Run with Docker

1. **Build**

```shell script
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker run -p 8080:8080 api-rest
```

3. **Test**

```shell script
go test -v ./test/...
```