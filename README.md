# Desafio

[![Go Report Card](https://goreportcard.com/badge/github.com/willystadnick/dvdscripter-superheroapi)](https://goreportcard.com/report/github.com/willystadnick/dvdscripter-superheroapi)
[![codecov](https://codecov.io/gh/willystadnick/dvdscripter-superheroapi/branch/master/graph/badge.svg)](https://codecov.io/gh/willystadnick/dvdscripter-superheroapi)
[![Build Status](https://travis-ci.org/willystadnick/dvdscripter-superheroapi.svg?branch=master)](https://travis-ci.org/willystadnick/dvdscripter-superheroapi)

Solução para o [desafio](challenge.md).

## Dependências

- [Go](https://golang.org/) 1.13
- [PostgreSQL](https://www.postgresql.org/) 12.3
- (Opcional) [Docker](https://www.docker.com/) 19.03

## Configurações

Copie o arquivo de configuração e edite as variáveis de ambiente.

```
cp configuration.toml.example configuration.toml
vi configuration.toml
```

As variáveis de exemplo coincidem com as setadas no arquivo [docker-compose.yml](docker-compose.yml) para facilitar a utilização do Docker:

```
# Subir os containers
docker-compose up -d

# Configurar o banco
cat postgres-setup.sql | docker exec -i superheroapi_postgres psql -U superheroapi -d superheroapi

# Descer os containers
docker-compose down
```

## Compilação

```
go build -o superheroapi
```

## Execução

```
./superheroapi --help
```

## Testes

```
go test ./...
```

Também é possível testar utilizando [cURL](curl.md).

## Autores

| [<img src="https://avatars1.githubusercontent.com/u/1256235?s=120&v=4" width=120><br><sub>@dvdscripter</sub>](https://github.com/dvdscripter) | [<img src="https://avatars2.githubusercontent.com/u/1824706?s=120&v=4" width=120><br><sub>@willystadnick</sub>](https://github.com/willystadnick) |
| :---: | :---: |
