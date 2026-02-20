# Real State API

Utilizou-se arquitetura hexagonal nesta aplicaćão com o intuito de separar as responsabilidades e manter o domínio de negócio independente das tecnologias utilizadas.


## Estrutura de Pastas
📁 project/ </br>
├── cmd/ </br>
│   └── api/ </br>
│       └── main.go            # Ponto de entrada</br>
├── internal/</br>
│   ├── domain/                # Entidades e regras de negócio puras</br>
│   ├── application/           # Casos de uso (orquestração)</br>
│   ├── ports/                 # Interfaces (contratos)</br>
│   ├── adapters/</br>
│   │   ├── handler/           # Handlers HTTP (entrada)</br>
│   │   ├── repository/        # Implementações de repositório (saída)</br>
│   │   └── middleware/</br>
└── migrations/</br>


## Bibliotecas utilizadas
- [GORM](https://gorm.io/)
- [Gin](https://gin-gonic.com/)
- [Viper](https://github.com/spf13/viper) versão 1.18.0 para leitura de variáveis no dockerfile
- [Zap](https://github.com/uber-go/zap) 