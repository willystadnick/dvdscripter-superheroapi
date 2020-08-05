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
curl localhost:8000/supers/ -H "Content-Type: application/json" -d '{"name":"superman","groups":[{"name":"Justice League"}]}'
```

Consulte um super, uuid abaixo é um exemplo!
```
curl localhost:8000/supers/68ca3ad5-9dc6-4ee0-b11a-ad0d1513c1d5 -H "Content-Type: application/json"
```

Delete um super, uuid abaixo é um exemplo!
```
curl localhost:8000/supers/68ca3ad5-9dc6-4ee0-b11a-ad0d1513c1d5 -H "Content-Type: application/json" -X DELETE
```

### Dicas

Você pode utilizar a ferramenta [jq](https://stedolan.github.io/jq/) para formatar a resposta dos comandos acima. Por exemplo:
```
curl ... | jq .
```
