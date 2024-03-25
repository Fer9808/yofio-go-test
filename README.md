# YoFio Go Test

Prueba técnica en Go de asignación de créditos

## Instrucciones

**Base de datos**
El proyecto implementa una BD SQLite la cual esta embebida en el proyecto.
Para revisarla es necesario abrirla desde algún Gestor de Bases de Datos como (DBaver) la cual se encuentra en la ruta

yofio-go-test/data/database.db

**Logs**
El proyecto ejecuta una trazabilidad de logs la cual podemos revisar dentro de la siguiente ruta

yofio-go-test/log/api.log

### Ejecución Local
1. **Ejecutar el proyecto con en local**

```shell script
make run
```

2. **Ejecutar la solicitud de asignación crédito**

```shell script
curl --location 'http://localhost:8080/api/credit-assignment' \
--header 'accept: application/json' \
--header 'Content-Type: text/plain' \
--data '{
    "investment": 10000
}'
```

3. **Ejecutar la solicitud para obtener las estadisticas**

```shell script
curl --location --request POST 'http://localhost:8080/api/statistics'
```

### Ejecución con Docker

1. **Ejecutar el comando para generar la imagen**

```shell script
make build
docker build . -t api-rest
```

2. **Ejecutar el comando para correr la imagen**

```shell script
docker run -p 8080:8080 api-rest
```

3. **Ejecutar la solicitud de asignación crédito**

```shell script
curl --location 'http://localhost:8080/api/credit-assignment' \
--header 'accept: application/json' \
--header 'Content-Type: text/plain' \
--data '{
    "investment": 10000
}'
```

4. **Ejecutar la solicitud para obtener las estadisticas**

```shell script
curl --location --request POST 'http://localhost:8080/api/statistics'
```

### Ejecución en ambiente Cloud

1. **Ejecutar la solicitud de asignación crédito**

```shell script
curl --location 'https://img-service-nixm7q4ula-uc.a.run.app/api/credit-assignment' \
--header 'Content-Type: application/json' \
--data '{
    "investment": 700
}'
```

2. **Ejecutar la solicitud para obtener las estadisticas**

```shell script
curl --location --request POST 'https://img-service-nixm7q4ula-uc.a.run.app/api/statistics'
```