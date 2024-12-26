# golang_api

# Aula 1
 
```bash
go mod init golang-api
```

# Aula 2

## Instalando gin

```bash
go get github.com/gin-gonic/gin
```

## Chamando a api

```bash
curl http://localhost:8080/v1/horas
```

## Chamada de api com educação

```bash
curl -X POST http://localhost:8080/v1/horas -H "Content-Type: application/json" -d '{"mensagem":"por_favor"}'
```
