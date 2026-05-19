# Arena dos Bárbaros API

## Visão Geral do Projeto

### Nome do Projeto
Arena dos Bárbaros API

### Descrição
Arena dos Bárbaros é uma API RESTful desenvolvida em Go para gerenciamento de batalhas PvP entre jogadores.

O sistema permitirá:
- autenticação de usuários,
- criação de personagens,
- batalhas entre jogadores,
- sistema de ranking,
- leaderboard em tempo real utilizando Redis,
- e gerenciamento seguro de sessões com JWT e Refresh Tokens.

O projeto será construído com foco em:
- arquitetura backend profissional,
- segurança,
- performance,
- escalabilidade,
- e boas práticas modernas de APIs REST.

---

# Objetivo do Projeto

O objetivo principal do projeto é praticar e consolidar conhecimentos avançados de desenvolvimento backend utilizando Go.

O projeto focará principalmente em:

- autenticação moderna com JWT,
- Refresh Tokens,
- middleware,
- cache com Redis,
- rate limiting,
- arquitetura em camadas,
- relacionamento entre entidades,
- validação de dados,
- Dockerização,
- e boas práticas REST.

Além disso, o sistema simulará um ambiente competitivo PvP entre jogadores, tornando o aprendizado mais divertido e próximo de um sistema real.

---

# Tecnologias Utilizadas

## Backend
- Go

## Framework HTTP
- Gin ou Chi

## Banco de Dados
- PostgreSQL

## Cache e Leaderboard
- Redis

## Autenticação
- JWT
- Refresh Tokens

## Containerização
- Docker
- Docker Compose

## Desenvolvimento
- Air (Hot Reload)

## Validação
- go-playground/validator

## ORM / SQL
(Pode ser definido posteriormente)

- GORM
ou
- SQLC
ou
- pgx

---

# Funcionalidades do Sistema

## 1. Autenticação

O sistema deverá permitir:

- Registro de usuário
- Login
- Logout
- Geração de Access Token
- Geração de Refresh Token
- Renovação de sessão
- Proteção de rotas privadas

---

## 2. Sistema de Personagens

Usuários poderão:

- Criar personagens
- Escolher nome do personagem
- Escolher classe
- Visualizar atributos
- Editar personagem

---

## 3. Sistema de Batalha

O sistema deverá permitir:

- Iniciar batalha PvP
- Calcular dano
- Determinar vencedor
- Atualizar ranking
- Registrar histórico de batalhas

---

## 4. Ranking Global

O sistema deverá:

- Exibir ranking global
- Exibir top jogadores
- Atualizar leaderboard em Redis
- Consultar ranking rapidamente via cache

---

# Regras de Negócio

## Usuários
- Cada usuário pode possuir múltiplos personagens.
- O email deve ser único.
- A senha deve ser criptografada.

---

## Personagens
- Cada personagem pertence a um usuário.
- O nome do personagem deve ser único.
- O personagem inicia no nível 1.
- O personagem inicia com 100 de HP.

---

## Batalhas
- Apenas personagens vivos podem lutar.
- O vencedor recebe pontos de ranking.
- O derrotado perde pontos.
- Todas as batalhas devem ser registradas.

---

## Ranking
- O ranking será ordenado por pontos.
- O leaderboard será armazenado em Redis.
- O banco PostgreSQL continuará sendo a fonte oficial de dados.

---

# Entidades do Sistema

## User

| Campo | Tipo |
|---|---|
| id | UUID |
| name | string |
| email | string |
| password | string |
| created_at | timestamp |

---

## Character

| Campo | Tipo |
|---|---|
| id | UUID |
| user_id | UUID |
| name | string |
| class | string |
| level | int |
| hp | int |
| attack | int |
| defense | int |
| ranking_points | int |

---

## Battle

| Campo | Tipo |
|---|---|
| id | UUID |
| attacker_id | UUID |
| defender_id | UUID |
| winner_id | UUID |
| damage_dealt | int |
| created_at | timestamp |

---

# Arquitetura do Projeto

O sistema utilizará arquitetura em camadas:

```text
Handler/API Layer
↓
Service Layer
↓
Repository Layer
↓
Database
```

---

## Responsabilidade das Camadas

### Handler
Responsável por:
- receber requests,
- validar entrada,
- retornar responses HTTP.

---

### Service
Responsável por:
- regras de negócio,
- cálculos,
- autenticação,
- batalhas.

---

### Repository
Responsável por:
- comunicação com PostgreSQL,
- queries,
- persistência.

---

# Segurança

## Autenticação JWT
O sistema utilizará:
- Access Token
- Refresh Token

---

## Rate Limiting
Será implementado:
- limite de requests por IP
- proteção contra spam e brute force

---

## Hash de Senha
As senhas serão criptografadas utilizando:
- bcrypt

---

# Cache e Leaderboard

Redis será utilizado para:

- armazenar leaderboard global,
- acelerar consultas de ranking,
- cachear dados frequentes,
- controle de rate limit,
- possível gerenciamento de sessões.

---

# Fluxo do Usuário

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
Ganha/perde ranking
↓
Leaderboard atualizado
```

---

# Roadmap de Desenvolvimento

## Fase 1 — Setup Inicial
- Estrutura do projeto
- Docker
- PostgreSQL
- Redis
- Air
- Configuração inicial

---

## Fase 2 — Autenticação
- Registro
- Login
- JWT
- Middleware
- Refresh Token

---

## Fase 3 — Personagens
- CRUD de personagens
- Relacionamento User → Character

---

## Fase 4 — Batalhas
- Sistema PvP
- Cálculo de dano
- Histórico

---

## Fase 5 — Ranking
- Redis Leaderboard
- Top jogadores

---

## Fase 6 — Segurança
- Rate Limiting
- Melhorias de middleware
- Hardening

---

# Objetivos Técnicos

Ao finalizar o projeto, os principais conhecimentos desenvolvidos serão:

- APIs REST profissionais em Go
- Arquitetura em camadas
- JWT Authentication
- Refresh Tokens
- Redis
- Cache
- Rate Limiting
- PostgreSQL
- Docker
- Middleware
- Relacionamentos entre entidades
- Segurança backend
- Estruturação de projetos escaláveis

---

# Futuras Melhorias

Possíveis evoluções futuras do projeto:

- Matchmaking automático
- Sistema de guildas
- Sistema de inventário
- Equipamentos e itens
- WebSocket para batalhas em tempo real
- Logs estruturados
- Observabilidade
- Testes automatizados
- CI/CD
- Deploy em nuvem

---

# Objetivo Final

Construir uma API backend moderna e escalável em Go, simulando um sistema PvP entre jogadores, utilizando autenticação segura, Redis para performance e arquitetura profissional para consolidar conhecimentos avançados de backend.