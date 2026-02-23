# Real Estate Advertising System

A aplicação deverá permitir cadastrar, consultar e visualizar os anúncios
de imóveis para venda ou aluguel, bem como o gerenciamento de cotações para
conversão de valores entre Real (BRL) e Dólar (USD).


## Como Executar

Clone o repositório
```bash
git clone https://github.com/gusgewehr/Real-Estate-Advertising-System.git
```

Entre na pata do projeto
```bash
cd Real-Estate-Advertising-System
```


Rode o comando
```bash
docker compose up --build -d
```
e aguarde enquanto o sistema fica online.

### Pontos importantes
É possível que na primeira execução o container da api não suba pois o container do postgres ainda não estava disponível. Nesse caso basta executar o container da api novamente
```bash
docker compose up -d api
```

### Para rodar o frontend no Desktop
Execute o seguinte comando, prestando atenção nas variáveis de ambiente. O valor da "API_URL" deve ser o local em q a api está rodando. Já o valor do "LOCALE" deve ser um dos valores aceitos demonstrados mais adiante neste documento. Ele só foi testado em ambiente windows.
```bash
cd frontend
flutter run -d windows --dart-define=API_URL=http://localhost:8080 --dart-define=LOCALE=pt
```


## Links
- http://localhost:3000 para acessar o frontend na web
- http://localhost:8080/docs/index.html para acessar a documentação do backend 

### Aviso!
Para que as imagens sejam carregadas na web é necessário descomentar a seguinte linha no pubsec.yaml
```yaml
#      - ./uploads/public/
```

## Internacionalização
Para trocar a língua que o frontend está mostrando basta alterar a variável LOCALE no docker-compose.yaml e buildar o container novamente

### Línguas suportadas
- Português: iserir "pt" no valor da variável de ambiente
- Inglês: inserir "en" no valor da variável de ambiente
