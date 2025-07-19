# GoFitnessApp - API de Exercícios

Uma API REST simples para gerenciar exercícios físicos, construída em Go.

## Estrutura do Projeto

```
GoFitnessApp/
├── main.go                          # Arquivo principal da aplicação
├── go.mod                           # Módulo Go com dependências
├── internal/                        # Código interno da aplicação
│   ├── model/                       # Modelos de dados
│   │   └── Exercicio.go
│   ├── repository/                  # Camada de acesso a dados
│   │   └── ExercicioRepository.go
│   ├── service/                     # Lógica de negócio
│   │   └── ExercicioService.go
│   └── api/                         # Handlers/Controllers da API
│       └── ExercicicioHandler.go
├── db/                              # Scripts de banco de dados
├── docs/                            # Documentação
└── terraform/                       # Infraestrutura como código
```

## Como Executar

1. **Instalar dependências:**
   ```bash
   go mod tidy
   ```

2. **Executar a aplicação:**
   ```bash
   go run main.go
   ```

3. **A API estará disponível em:** `http://localhost:8080`

## Endpoints da API

### Health Check
- **GET** `/health` - Verificar se a API está funcionando

### Exercícios
- **POST** `/exercicios` - Criar um novo exercício
- **GET** `/exercicios` - Listar todos os exercícios
- **GET** `/exercicios/:id` - Obter um exercício específico
- **PUT** `/exercicios/:id` - Atualizar um exercício
- **DELETE** `/exercicios/:id` - Deletar um exercício

## Exemplos de Uso

### Criar Exercício
```bash
curl -X POST http://localhost:8080/exercicios \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Flexão de Braço",
    "series": 3,
    "repeticoes": 15,
    "descanso": 60
  }'
```

### Listar Exercícios
```bash
curl http://localhost:8080/exercicios
```

### Obter Exercício Específico
```bash
curl http://localhost:8080/exercicios/{id}
```

### Atualizar Exercício
```bash
curl -X PUT http://localhost:8080/exercicios/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Flexão de Braço Modificada",
    "series": 4,
    "repeticoes": 12,
    "descanso": 45
  }'
```

### Deletar Exercício
```bash
curl -X DELETE http://localhost:8080/exercicios/{id}
```

## Modelo de Dados

### Exercício
```json
{
  "id": "uuid",
  "nome": "string",
  "series": "int",
  "repeticoes": "int", 
  "descanso": "int (em segundos)",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

## Arquitetura

A aplicação segue uma arquitetura em camadas:

1. **API Layer (Handlers)**: Recebe requests HTTP e retorna responses
2. **Service Layer**: Contém a lógica de negócio
3. **Repository Layer**: Abstrai o acesso aos dados
4. **Model Layer**: Define as estruturas de dados

Atualmente usa armazenamento em memória, mas pode ser facilmente expandido para usar banco de dados.

## Próximos Passos

- [ ] Adicionar persistência em banco de dados (PostgreSQL/MySQL)
- [ ] Implementar autenticação e autorização
- [ ] Adicionar testes unitários e de integração
- [ ] Documentação com Swagger
- [ ] Docker e docker-compose
- [ ] Logging estruturado
- [ ] Validações mais robustas
- [ ] Middleware de CORS