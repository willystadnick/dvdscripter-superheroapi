# Desafio

[![Go Report Card](https://goreportcard.com/badge/github.com/willystadnick/dvdscripter-superheroapi)](https://goreportcard.com/report/github.com/willystadnick/dvdscripter-superheroapi)
[![codecov](https://codecov.io/gh/willystadnick/dvdscripter-superheroapi/branch/master/graph/badge.svg)](https://codecov.io/gh/willystadnick/dvdscripter-superheroapi)
[![Build Status](https://travis-ci.org/willystadnick/dvdscripter-superheroapi.svg?branch=master)](https://travis-ci.org/willystadnick/dvdscripter-superheroapi)

Solução para o [desafio](challenge.md).

## Dependências

Esse projeto usa módulos então você precisa da linguagem Go >= 1.13.

## Configurações

Copie o arquivo de configuração e edite as variáveis de ambiente.

```
cp configuration.toml.example configuration.toml
vi configuration.toml
```

As variáveis de exemplo coincidem com as setadas no arquivo [docker-compose.yml](docker-compose.yml) para facilitar a utilização do Docker.

## Execute os testes (não precisa de postgresql)

```
go test ./...
```

Também é possível testar utilizando [cURL](curl.md).

## Compile o projeto

```
go build -o superheroapi
```

## Execute

```
./superheroapi --help
```

## PostgreSQL

O projeto conta com migrações automáticas e para isso precisamos de um banco de dados e um usuário com permissões para modificar o banco em questão.

Existe a opção de utilizar o [Docker](https://www.docker.com/):

```
# Subir os containers
docker-compose up -d

# Configurar o banco
cat postgres-setup.sql | docker exec -i superheroapi_postgres psql -U superheroapi -d superheroapi

# Descer os containers
docker-compose down
```
