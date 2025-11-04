# Desafio Clean Architecture Full Cycle

Este projeto é parte de um desafio do curso Pós Go Expert Full Cycle, cujo objetivo era desenvolver uma listagem de orders, com suporte a três diferentes abordagens para acessar o serviço: REST, gRPC e GraphQL.

## 🚧 Como instalar e rodar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/dabraga/Desafio03_Clean-Architecture.git
cd Desafio03_Clean-Architecture
```

### 2. Rodar o docker compose

Subir uma instância do mysql e rabbitmq com o docker compose.

```bash
docker compose up -d
```

### 3. Rodar as migrations

```bash
make migrate
```

### 4. Rodar a aplicação

```bash
make run
```

A aplicação estará acessível nas seguintes portas:

- REST API: http://localhost:8000/orders
- gRPC: http://localhost:50051
- GraphQL: http://localhost:8080
