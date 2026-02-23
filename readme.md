# Real Estate Advertising System

A aplicação deverá permitir cadastrar, consultar e visualizar os anúncios
de imóveis para venda ou aluguel, bem como o gerenciamento de cotações para
conversão de valores entre Real (BRL) e Dólar (USD).


## Como Executar

Rode o comando
```bash
docker compose up --build -d
```
e aguarde enquanto o sistema fica online.

## Pontos importantes
É possível que na primeira execução o container da api não suba pois o container do postgres ainda não estava disponível. Nesse caso basta executar o container da api novamente
```bash
docker compose up -d api
```


## Links
- http://localhost:3000 para acessar o frontend
- http://localhost:8080/docs/index.html para acessar a documentação do backend


## Internacionalização
Para trocar a língua que o frontend está mostrando basta alterar a variável LOCALE no docker-compose.yaml e buildar o container novamente

### Línguas suportadas
- Português: iserir "pt" no valor da variável de ambiente
- Inglês: inserir "en" no valor da variável de ambiente
