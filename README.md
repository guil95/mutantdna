# mutant-dna

# Como subir a aplicação

Execute o seguinte comand `make start-migrate` além de subir a aplicação o mesmo ja executa as migrations.

# Recursos

## DNA mutant

### Rota
`http://localhost/mutant`

### Payload
```json
{
	"dna": ["ATGCGA", "CAGTGC", "TTATGT", "AGAGAG", "CCCTTA", "TCACTG"]
}
```

### Reponse
Response esperado para esse DNA é 403 por se tratar de DNA humano

## Estatísticas

### Rota
`http://localhost/stats`

### Response
```json
{
	"count_mutant_dna": 2,
	"count_human_dna": 5,
	"ratio": 0.4
}
```

# Outros comandos

Para executar os testes da aplicação execute `make test`
Para rodar as migrations e subir a aplicação separadamente execute `make start` e após `make migrate`

Demais comandos podem ser encontrados [aqui](./Makefile)
