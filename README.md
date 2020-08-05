# Desafio
[![Go Report Card](https://goreportcard.com/badge/github.com/dvdscripter/careers)](https://goreportcard.com/report/github.com/dvdscripter/careers)
[![codecov](https://codecov.io/gh/dvdscripter/careers/branch/master/graph/badge.svg)](https://codecov.io/gh/dvdscripter/careers)
[![Build Status](https://travis-ci.org/dvdscripter/careers.svg?branch=master)](https://travis-ci.org/dvdscripter/careers)

Solução para o [desafio](challenge.md).

## Dependências

Esse projeto usa módulos então você precisa da linguagem Go >= 1.13.

## Execute os testes (não precisa de postgresql)

```
go test ./...
```

## Compile o projeto

```
go build -o superheroapi
```

## Execute

```
./superheroapi --help
```

## Testando com cURL

Todos os supers
```
curl localhost:8000/supers/ -H "Content-Type: application/json"
```

Procure um super pelo nome
```
curl localhost:8000/supers/?name=superman -H 'Content-Type: application/json'
```

Todos os supers bonzinhos
```
curl localhost:8000/supers/?alignment=good -H "Content-Type: application/json"
```

Todos os supers malvados
```
curl localhost:8000/supers/?alignment=bad -H "Content-Type: application/json"
```

Adicione novo super
```
curl localhost:8000/supers/ -H "Content-Type: application/json" -d '{"name":"superman",
"groups": [{"name": "Justice League"}]}'
```

Consulte um super, uuid abaixo é um exemplo!
```
curl localhost:8000/supers/68ca3ad5-9dc6-4ee0-b11a-ad0d1513c1d5 -H "Content-Type: application/json"
```

Delete um super, uuid abaixo é um exemplo!
```
curl localhost:8000/supers/68ca3ad5-9dc6-4ee0-b11a-ad0d1513c1d5 -H "Content-Type: application/json" -X DELETE
```

## PostgreSQL
O projeto conta com migrações automáticas e para isso precisamos de um banco de
dados e um usuário com permissões para modificar o banco em questão. Alguns exemplos abaixo.

```
createdb --owner=superadminuser super
```

Sempre existe a opção de rodar um docker de postgres.

Você também precisa renomear o arquivo ```configuration.toml.example``` para ```configuration.toml```.

Precisamos editar o arquivo com as configurações do seu postgres e banco criado.

```
[database]
dsn = "postgres://<usuário>:<senha>@<endereço>/<nome do banco>?<opções>"
```

_dsn_ significa data source name, mais informações podem ser encontradas em
https://pkg.go.dev/github.com/lib/pq?tab=doc

## Outras configurações

Edite o ```configuration.toml``` e escolha o endereço e parta de execução.

```
[server]
bind = "127.0.0.1:8000"
```
