# Arena dos Bárbaros API — Software Design Document (SDD)

# 1. Arquitetura Geral

O sistema seguirá arquitetura em camadas (Layered Architecture), separando responsabilidades para manter:

- organização,
- escalabilidade,
- legibilidade,
- testabilidade,
- manutenção.

---

## Arquitetura da Aplicação

```text
Client
↓
HTTP Router
↓
Middleware Layer
↓
Handlers
↓
Services
↓
Repositories
↓
PostgreSQL / Redis
```

---

## Responsabilidade das Camadas

### Router
Responsável por:
- registrar rotas,
- agrupar endpoints,
- aplicar middlewares.

---

### Middleware
Responsável por:
- autenticação JWT,
- rate limiting,
- logging,
- recuperação de panic,
- validações globais.

---

### Handlers
Responsável por:
- receber requests HTTP,
- validar payloads,
- chamar services,
- retornar responses HTTP.

---

### Services
Responsável por:
- regras de negócio,
- autenticação,
- cálculos de batalha,
- ranking,
- manipulação de tokens.

---

### Repositories
Responsável por:
- queries SQL,
- persistência,
- acesso ao banco de dados.

---

# 2. Estrutura de Pastas

```text
arena-dos-barbaros-api/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── model/
│   ├── dto/
│   ├── auth/
│   ├── battle/
│   ├── ranking/
│   ├── cache/
│   ├── config/
│   └── validator/
│
├── pkg/
│   ├── database/
│   ├── redis/
│   ├── logger/
│   └── utils/
│
├── migrations/
│
├── docker/
│
├── docs/
│
├── .env
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## Explicação das Pastas

### cmd/api
Ponto de entrada da aplicação.

Responsável por:
- iniciar servidor,
- carregar configurações,
- conectar banco de dados,
- registrar rotas.

---

### internal/
Contém toda a lógica interna da aplicação.

---

### handler/
Camada HTTP da aplicação.

Exemplos:
- auth_handler.go
- character_handler.go

---

### service/
Camada de regras de negócio.

Exemplos:
- auth_service.go
- battle_service.go

---

### repository/
Camada de acesso ao banco.

Exemplos:
- user_repository.go
- battle_repository.go

---

### middleware/
Middlewares HTTP.

Exemplos:
- jwt_middleware.go
- rate_limit.go

---

### model/
Modelos principais da aplicação.

Exemplos:
- user.go
- character.go
- battle.go

---

### dto/
Objetos de entrada e saída da API.

Exemplos:
- login_request.go
- register_response.go

---

### auth/
Lógica de autenticação JWT e Refresh Token.

---

### battle/
Lógica de combate e cálculos de batalha.

---

### ranking/
Lógica de ranking e leaderboard.

---

### cache/
Integração com Redis.

---

### config/
Carregamento de variáveis de ambiente e configurações.

---

# 3. Banco de Dados

## Banco Principal
PostgreSQL

## Cache e Leaderboard
Redis

---

## Entidades Principais

```text
User
│
└── Character
       │
       └── Battle
```

---

## Relacionamentos

### User → Character
- Um usuário pode possuir vários personagens.

---

### Character → Battle
- Um personagem pode participar de várias batalhas.

---

# 4. Fluxo de Autenticação

## Registro

```text
Usuário envia email e senha
↓
Senha é criptografada com bcrypt
↓
Usuário salvo no PostgreSQL
```

---

## Login

```text
Usuário envia credenciais
↓
Sistema valida senha
↓
Gera Access Token
↓
Gera Refresh Token
↓
Retorna tokens
```

---

## Tokens

### Access Token
- curta duração,
- utilizado para autenticação das rotas privadas.

Tempo sugerido:
- 15 minutos.

---

### Refresh Token
- longa duração,
- utilizado para renovação de sessão.

Tempo sugerido:
- 7 dias.

---

## Fluxo Refresh Token

```text
Access Token expira
↓
Cliente envia Refresh Token
↓
Sistema valida token
↓
Novo Access Token é gerado
```

---

# 5. Middlewares

## JWT Middleware
Responsável por:
- validar token,
- proteger rotas privadas,
- identificar usuário autenticado.

---

## Rate Limiting Middleware
Responsável por:
- limitar requests por IP,
- evitar spam,
- evitar brute force.

Redis será utilizado para armazenar contadores temporários.

---

## Recovery Middleware
Responsável por:
- evitar crash da aplicação.

---

## Logging Middleware
Responsável por:
- registrar requests HTTP.

---

# 6. Sistema de Batalha

## Fluxo da batalha

```text
Jogador desafia outro jogador
↓
Sistema carrega atributos
↓
Calcula dano
↓
Define vencedor
↓
Atualiza ranking
↓
Salva histórico
↓
Atualiza leaderboard Redis
```

---

## Fórmula inicial de dano

```text
Dano = Ataque - (Defesa / 2)
```

---

## Critério de vitória

Vence:
- quem reduzir o HP do adversário para 0 primeiro.

---

# 7. Estratégia Redis

Redis será utilizado para:

- leaderboard,
- cache de ranking,
- rate limiting,
- sessões temporárias,
- cache de personagens populares.

---

## Leaderboard

Estrutura utilizada:

```text
ZSET leaderboard
```

Onde:
- chave = character_id
- score = ranking_points

---

# 8. Endpoints

## Auth

| Método | Endpoint |
|---|---|
| POST | /auth/register |
| POST | /auth/login |
| POST | /auth/refresh |

---

## Characters

| Método | Endpoint |
|---|---|
| POST | /characters |
| GET | /characters |
| GET | /characters/:id |
| PUT | /characters/:id |
| DELETE | /characters/:id |

---

## Battles

| Método | Endpoint |
|---|---|
| POST | /battles |
| GET | /battles/history |

---

## Ranking

| Método | Endpoint |
|---|---|
| GET | /ranking |
| GET | /ranking/top |

---

# 9. Response Pattern

## Sucesso

```json
{
  "success": true,
  "data": {}
}
```

---

## Erro

```json
{
  "success": false,
  "error": "invalid credentials"
}
```

---

# 10. Dockerização

O ambiente será composto por:

- API Go,
- PostgreSQL,
- Redis.

---

## Docker Compose

```text
services:
  - api
  - postgres
  - redis
```

---

# 11. Segurança

## Medidas de Segurança

- bcrypt para hash de senhas,
- JWT seguro,
- expiração curta para access token,
- refresh token,
- rate limiting,
- validação de payload,
- variáveis sensíveis em `.env`.

---

# 12. Estratégia de Escalabilidade

O sistema será preparado para:

- separação futura em microserviços,
- horizontal scaling,
- cache distribuído,
- filas assíncronas futuramente.

---

# 13. Fluxo Geral do Sistema

```text
Registro
↓
Login
↓
Recebe JWT
↓
Cria personagem
↓
Entra em batalha
↓
Sistema calcula combate
↓
Ranking atualizado
↓
Redis atualiza leaderboard
```

---

# 14. Objetivos Técnicos do Projeto

Ao finalizar o projeto, os conhecimentos desenvolvidos serão:

- APIs REST profissionais em Go,
- arquitetura em camadas,
- JWT Authentication,
- Refresh Tokens,
- Redis,
- cache,
- rate limiting,
- PostgreSQL,
- Docker,
- middleware,
- relacionamentos entre entidades,
- segurança backend,
- estruturação de projetos escaláveis.

---

# 15. Futuras Melhorias

Possíveis evoluções futuras:

- matchmaking automático,
- sistema de guildas,
- inventário,
- equipamentos,
- WebSocket para batalhas em tempo real,
- logs estruturados,
- observabilidade,
- testes automatizados,
- CI/CD,
- deploy em nuvem.

---

# 16. Objetivo Final

Construir uma API backend moderna e escalável em Go, simulando um sistema PvP entre jogadores, utilizando autenticação segura, Redis para performance e arquitetura profissional para consolidar conhecimentos avançados de backend.