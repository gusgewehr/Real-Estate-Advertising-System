# Real State API

## Requisitos
O backend deverá ser desenvolvido em Go, utilizando banco de dados PostgreSQL e executado via Docker.

A API deverá ser responsável por:
- Criar anúncios de imóveis;
- Criar cotações BRL -> USD;
- Realizar integração com a API do ViaCEP para consulta de endereços;
- Validar dados recebidos do frontend;
- Tratar adequadamente erros internos e externos;
- Documentação com Swagger (diferencial).
- Testes para as funcionalidades


## Arquitetura
Utilizou-se arquitetura hexagonal nesta aplicaćão com o intuito de separar as responsabilidades e manter o domínio de negócio independente das tecnologias utilizadas.


## Estrutura de Pastas
📁 project/ </br>
├── cmd/ </br>
│   └── api/ </br>
│       └── main.go            # Ponto de entrada</br>
└── internal/</br>
│   ├── domain/                # Entidades e regras de negócio dessas entidades</br>
│   ├── application/           # Camada de regra de negócios </br>
│   │   ├── ports/                 # Interfaces </br>
│   │   └── usecase/               # Casos de uso </br>
│   └── adapters/  # Camada de interacão como mundo externo </br>
│   │   ├── handler/           # Handlers HTTP </br>
│   │   ├── repository/        # Conexões com o DB</br>
│   │   └── gateway/           # Conexões com outras APIs</br>
└docs/ # Swagger

## Bibliotecas utilizadas
- [GORM](https://gorm.io/): por ser um projeto pequeno que nao demanda muito controle;
- [Gin](https://gin-gonic.com/): servidor http de rápido desenvolvimento;
- [Viper](https://github.com/spf13/viper): versão 1.18.0 para leitura de variáveis no dockerfile;
- [Zap](https://github.com/uber-go/zap): logs mais semanticos;
- [Resty](https://resty.dev): boas ferramentas para requisicoes http;
- [Testify](https://github.com/stretchr/testify): ferramenta de testes unitários;


## FileStorage
Decidiu-se por segregar a lógica de upload de imagens para uma possível integraćão com algum sistema de Bucket. Por enquanto não há como deletar ou atualizar imagens, então caso um usuário cancele a criaćão de um imóvel a imagem ficará sem nenhum lastro. Posteriormente as funcionalidades de upload e delete serão implementadas. 


## RealEstateProperty
Construído por dois objetos para realizar a transferência entre as camadas. É o coraćão da aplicaćão, onde todos os registros de imóveis são gerenciados. Por enquanto só é possível criar e listar os registros. As implementaćões de Update e Delete serão feitas posteriormente.


## ExchangeRate
Construído por dois objetos para realizar a transferência entre as camadas.
Salva somente a proporcionalidade de USD para BRL. Ou seja, quantos dólares vale R$1,00.
Planos futuros incluem uma outra općão para cadastro de moedas, dessa forma será possível construir uma conversão para cara relaćão entre moedas.

## Pagination
Construído para fornecer informaćões referentes a quantidade de registros e limitar a quantidade de resultados enviados em uma única requisićão.


## Zipcode
Comunica-se com a aplicaćão da ViaCEP para obter informações sobre o CEP informado pelo usuário.
Como o parceiro sempre retorna o status code 200, foi necessário criar uma estrutura para tratar erros.