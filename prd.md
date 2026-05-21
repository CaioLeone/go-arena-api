# Arena dos Bárbaros — Fullstack (Go + React)

## Visão Geral do Projeto

### Nome do Projeto
Arena dos Bárbaros

### Descrição
**Arena dos Bárbaros** é uma aplicação **fullstack** para gerenciamento de batalhas PvP entre jogadores. Combina:

- **Backend**: API RESTful em Go com autenticação segura
- **Frontend**: Interface moderna em React + TypeScript

O sistema permitirá:
- autenticação de usuários,
- criação e gerenciamento de personagens,
- batalhas entre jogadores em tempo real,
- sistema de ranking e leaderboard,
- interface responsiva e intuitiva.

---

## Objetivo do Projeto

Praticar e consolidar conhecimentos **avançados de desenvolvimento fullstack**:

**Backend (Go):**
- Arquitetura profissional em camadas
- Autenticação JWT segura
- PostgreSQL + Redis
- Middleware e rate limiting
- Docker e containerização

**Frontend (React):**
- Componentes modernos em React 18
- TypeScript para type safety
- Context API para gerenciamento de estado
- Integração com API REST
- Interface responsiva com TailwindCSS

**Meta:** Criar um portfólio fullstack competitivo no mercado.

---

## Tecnologias Utilizadas

### Backend
- **Go** 1.22+
- **Gin** (framework HTTP)
- **PostgreSQL** (banco de dados)
- **Redis** (cache e leaderboard)
- **JWT** (autenticação)
- **Docker** (containerização)

### Frontend
- **React** 18
- **TypeScript**
- **React Router** v6 (routing)
- **Axios** (HTTP client)
- **TailwindCSS** (styling)
- **Vite** (build tool)

### DevTools
- **Air** (hot reload - backend)
- **npm** (package manager - frontend)
- **Docker Compose** (orquestração)

---

## Funcionalidades do Sistema

### 1. Autenticação
- Registro de novo usuário
- Login com email/password
- Tokens JWT (access + refresh)
- Logout
- Proteção de rotas privadas

### 2. Sistema de Personagens
- Criar personagem com classe (Barbáro, Mago, Arqueiro, Assassino)
- Visualizar lista de personagens
- Ver detalhes do personagem (stats)
- Editar personagem
- Deletar personagem

### 3. Sistema de Batalha PvP
- Desafiar outro jogador
- Cálculo automático de dano
- Determinação automática de vencedor
- Atualização automática de ranking
- Histórico completo de batalhas

### 4. Ranking Global
- Leaderboard em tempo real
- Top 10 players
- Posição do jogador atual
- Atualização com cada batalha

### 5. Interface Gráfica (UI/UX)
- Dashboard intuitivo
- Formulários com validação
- Feedback visual (toasts, loading states)
- Responsividade em mobile
- Tema "barbárico" (visual coerente)

---

## Regras de Negócio

### Usuários
- Cada usuário pode possuir múltiplos personagens.
- O email deve ser único.
- A senha deve ser criptografada com bcrypt.

### Personagens
- Cada personagem pertence a um usuário.
- O nome do personagem deve ser único.
- O personagem inicia no nível 1.
- O personagem inicia com 100 de HP.

### Batalhas
- Apenas personagens vivos podem lutar.
- O vencedor recebe pontos de ranking.
- O derrotado perde pontos.
- Todas as batalhas devem ser registradas.

### Ranking
- O ranking será ordenado por pontos.
- O leaderboard será armazenado em Redis para performance.
- O banco PostgreSQL continua sendo a fonte oficial.

---

## Entidades do Sistema

### User
| Campo | Tipo |
|---|---|
| id | UUID |
| name | string |
| email | string |
| password | string (hash) |
| created_at | timestamp |

### Character
| Campo | Tipo |
|---|---|
| id | UUID |
| user_id | UUID (FK) |
| name | string |
| class | string |
| level | int |
| hp | int |
| attack | int |
| defense | int |
| ranking_points | int |

### Battle
| Campo | Tipo |
|---|---|
| id | UUID |
| attacker_id | UUID (FK) |
| defender_id | UUID (FK) |
| winner_id | UUID (FK) |
| damage_dealt | int |
| created_at | timestamp |

---

## Arquitetura Geral

### Backend (Go)
```
Handler/API Layer (Gin)
    ↓
Service Layer (Regras de negócio)
    ↓
Repository Layer (Acesso a dados)
    ↓
PostgreSQL / Redis
```

### Frontend (React)
```
Pages (Login, Characters, Battles, Leaderboard)
    ↓
Components (Cards, Modals, Forms)
    ↓
Hooks (useAuth, useCharacters, useBattles)
    ↓
Services (API calls com Axios)
    ↓
API Backend (Go)
```

---

## Fluxo do Usuário

```
1. Acessa aplicação (localhost:3000)
2. Não autenticado → Redireciona para Login
3. Faz Register ou Login
4. Recebe JWT, salvo em localStorage
5. Acessa Dashboard
6. Cria personagem
7. Seleciona personagem
8. Entra em batalha
9. Vê resultado
10. Ranking atualiza em tempo real
11. Consulta Leaderboard
```

---

## Roadmap de Desenvolvimento

### Fases 1-6: Backend
- Setup inicial, Autenticação, CRUD, Batalhas, Ranking, Segurança
- **Duração:** ~8-9h

### Fases 7-9: Frontend
- Setup React, UI Principal, Polish & Deploy
- **Duração:** ~6h

### Total: ~14-15h

---

## Tecnologias por Fase

### Fase 1-6 (Backend)
- Go, Gin, PostgreSQL, Redis, JWT, Docker

### Fase 7 (Frontend)
- React, TypeScript, React Router, Axios, Context API

### Fase 8 (UI)
- TailwindCSS, Componentes React, Custom Hooks

### Fase 9 (Polish)
- Toast notifications, Loading states, Error handling

---

## Segurança

- **Autenticação:** JWT com expiration curta
- **Hash de Senha:** bcrypt
- **CORS:** Configurado no Backend
- **Rate Limiting:** Proteção contra brute force
- **Validação:** Client-side (React) + Server-side (Go)

---

## Futuras Melhorias

- Testes automatizados (Jest, Go testing)
- CI/CD com GitHub Actions
- WebSocket para batalhas em tempo real
- Sistema de guildas
- Inventário e equipamentos
- Logs estruturados e observabilidade
- Deploy em nuvem (AWS, Vercel)

---

## Objetivo Final

Construir uma **aplicação fullstack moderna e escalável** que demonstre:
- Conhecimento profundo de Go backend
- Habilidades em React frontend
- Arquitetura escalável
- Boas práticas de segurança
- Experiência pronta para vagas de Fullstack/Backend em React+Go
