# Arena dos Bárbaros — Software Design Document (SDD)
## Fullstack: Go Backend + React Frontend

---

# 1. Arquitetura Geral Fullstack

## Visão Geral
O sistema é dividido em **duas camadas principais**:

```
┌─────────────────────────────────────────┐
│         React Frontend (3000)           │
│  (Components, Pages, Context, Hooks)    │
└──────────────┬──────────────────────────┘
               │ HTTP/JSON (Axios)
               │ JWT Tokens
               ▼
┌─────────────────────────────────────────┐
│      Go Backend API (8080)              │
│  (Handlers, Services, Repositories)     │
└──────────────┬──────────────────────────┘
               │
       ┌───────┴───────┐
       ▼               ▼
  PostgreSQL       Redis
  (Dados)          (Leaderboard)
```

---

## Arquitetura Backend (Go)

```text
Client (React)
↓ HTTP Requests
Router (Gin)
↓
Middleware Layer (JWT, CORS, Rate Limit)
↓
Handlers (recebem requests)
↓
Services (regras de negócio)
↓
Repositories (acesso dados)
↓
PostgreSQL / Redis
```

---

## Arquitetura Frontend (React)

```text
Browser
↓
React App
├── Pages (Routes)
│   ├── Login/Register
│   ├── Dashboard
│   ├── Characters
│   ├── Battles
│   └── Leaderboard
├── Components (Reusable)
├── Context (State Management)
├── Hooks (Custom Logic)
├── Services (API calls)
└── Types (TypeScript)
```

---

# 2. Estrutura de Pastas Completa

## Backend (Go)

```
go-arena-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   │   ├── auth_handler.go
│   │   ├── character_handler.go
│   │   └── battle_handler.go
│   ├── service/
│   │   ├── auth_service.go
│   │   ├── character_service.go
│   │   └── battle_service.go
│   ├── repository/
│   │   ├── user_repository.go
│   │   ├── character_repository.go
│   │   └── battle_repository.go
│   ├── middleware/
│   │   ├── jwt_middleware.go
│   │   ├── rate_limit.go
│   │   └── cors.go
│   ├── model/
│   │   ├── user.go
│   │   ├── character.go
│   │   └── battle.go
│   ├── dto/
│   │   ├── user_dto.go
│   │   ├── character_dto.go
│   │   └── battle_dto.go
│   ├── auth/
│   │   └── token.go
│   ├── battle/
│   │   └── calculator.go
│   ├── ranking/
│   │   └── leaderboard.go
│   ├── cache/
│   │   └── redis_cache.go
│   ├── config/
│   │   └── config.go
│   └── validator/
│       └── validator.go
├── pkg/
│   ├── database/
│   │   └── postgres.go
│   ├── redis/
│   │   └── client.go
│   ├── logger/
│   │   └── logger.go
│   └── utils/
│       └── utils.go
├── migrations/
│   ├── 001_create_users_table.up.sql
│   ├── 002_create_characters_table.up.sql
│   └── 003_create_battles_table.up.sql
├── docker/
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
├── .env
└── README.md
```

## Frontend (React)

```
frontend/
├── src/
│   ├── components/
│   │   ├── Navigation.tsx
│   │   ├── CharacterCard.tsx
│   │   ├── BattleCard.tsx
│   │   ├── LeaderboardRow.tsx
│   │   ├── LoadingSpinner.tsx
│   │   └── Toast.tsx
│   ├── pages/
│   │   ├── Login.tsx
│   │   ├── Register.tsx
│   │   ├── Dashboard.tsx
│   │   ├── Characters.tsx
│   │   ├── Battles.tsx
│   │   └── Leaderboard.tsx
│   ├── context/
│   │   ├── AuthContext.tsx
│   │   └── ToastContext.tsx
│   ├── hooks/
│   │   ├── useAuth.ts
│   │   ├── useCharacters.ts
│   │   ├── useBattles.ts
│   │   └── useLeaderboard.ts
│   ├── services/
│   │   ├── api.ts (Axios instance)
│   │   ├── authService.ts
│   │   ├── characterService.ts
│   │   ├── battleService.ts
│   │   └── leaderboardService.ts
│   ├── types/
│   │   └── index.ts
│   ├── App.tsx
│   ├── App.css
│   └── main.tsx
├── public/
├── index.html
├── package.json
├── tailwind.config.js
├── vite.config.ts
├── tsconfig.json
└── .env
```

---

# 3. Stack Tecnológico

## Backend
- **Go 1.22** - Linguagem
- **Gin** - Framework HTTP
- **PostgreSQL** - Banco de dados
- **Redis** - Cache e Leaderboard
- **JWT** - Autenticação
- **Docker** - Containerização

## Frontend
- **React 18** - UI Library
- **TypeScript** - Type safety
- **React Router v6** - Routing
- **Axios** - HTTP Client
- **TailwindCSS** - Styling
- **Vite** - Build tool

---

# 4. Fluxos Principais

## 4.1 Autenticação

### Registro
```
1. Usuário preenche form (email, password)
2. React valida client-side
3. Envia POST /auth/register
4. Backend:
   - Valida dados
   - Hash password com bcrypt
   - Cria user no PostgreSQL
   - Retorna UserResponse
5. React armazena user em Context
6. Redireciona para Dashboard
```

### Login
```
1. Usuário preenche form
2. Envia POST /auth/login
3. Backend:
   - Valida credenciais
   - Gera Access Token (15 min)
   - Gera Refresh Token (7 dias)
   - Retorna LoginResponse
4. React:
   - Armazena tokens em localStorage
   - Configura AuthContext
   - Redireciona Dashboard
```

### Token Flow
```
1. Todas as requisições incluem: Authorization: Bearer {token}
2. Backend valida JWT no middleware
3. Se válido: prossegue
4. Se expirado: retorna 401
5. React intercepta 401: redireciona login
6. Se Refresh token válido: gera novo Access token
```

---

## 4.2 CRUD de Personagens

```
Criar:
  POST /characters + {name, class}
  → Cria character com stats iniciais
  → Retorna CharacterResponse
  
Listar:
  GET /characters
  → Retorna lista de personagens do usuário
  
Detalhes:
  GET /characters/:id
  → Retorna stats completos
  
Editar:
  PUT /characters/:id + {name, class}
  → Atualiza dados
  
Deletar:
  DELETE /characters/:id
  → Remove personagem
```

---

## 4.3 Sistema de Batalha

```
Desafiar:
  1. Frontend: POST /battles + {defender_id}
  2. Backend:
     - Carrega stats do atacante e defensor
     - Calcula dano: Dano = Ataque - (Defesa/2)
     - Determina vencedor (quem reduz HP a 0)
     - Atualiza ranking (vencedor +10, perdedor -5)
     - Salva em PostgreSQL
     - Atualiza Redis Leaderboard
     - Retorna BattleResponse
  3. Frontend mostra resultado com animação
```

---

## 4.4 Leaderboard (Redis)

```
Estrutura: ZSET "leaderboard"
  - Member: character_id (string)
  - Score: ranking_points (int)

Após cada batalha:
  ZADD leaderboard {winner_points} {winner_id}
  ZADD leaderboard {loser_points} {loser_id}

Consulta Top 10:
  ZRANGE leaderboard 0 9 WITHSCORES (REVERSE)

Frontend: Poll a cada 5 segundos (GET /ranking/top)
```

---

# 5. DTOs (Backend)

## Auth
```go
UserCreateRequest { email, password }
UserLoginRequest { email, password }
LoginResponse { access_token, refresh_token, user }
UserResponse { id, name, email }
```

## Characters
```go
CharacterCreateRequest { name, class }
CharacterUpdateRequest { name, class }
CharacterResponse { id, user_id, name, class, level, hp, attack, defense, ranking_points }
CharacterListResponse { characters[], total }
```

## Battles
```go
BattleCreateRequest { defender_id }
BattleResponse { id, attacker_id, defender_id, winner_id, damage_dealt, created_at }
BattleHistoryResponse { battles[], total }
```

---

# 6. Endpoints REST

## Auth
```
POST   /auth/register          → Register
POST   /auth/login             → Login
POST   /auth/refresh           → Refresh Token (futuro)
POST   /auth/logout            → Logout (futuro)
```

## Characters (protegidas por JWT)
```
POST   /characters             → Criar
GET    /characters             → Listar
GET    /characters/:id         → Detalhes
PUT    /characters/:id         → Editar
DELETE /characters/:id         → Deletar
```

## Battles (protegidas por JWT)
```
POST   /battles                → Iniciar batalha
GET    /battles/history        → Histórico
```

## Ranking
```
GET    /ranking                → Ranking do usuário
GET    /ranking/top            → Top 10 players
```

---

# 7. Integração Frontend ↔ Backend

## CORS Configuration (Backend)
```go
config := cors.DefaultConfig()
config.AllowOrigins = []string{"http://localhost:3000"}
router.Use(cors.New(config))
```

## Environment Variables

**Backend (.env):**
```
DB_HOST=postgres
DB_PORT=5432
REDIS_HOST=redis
REDIS_PORT=6379
JWT_SECRET=seu-secret-key
SERVER_PORT=8080
```

**Frontend (.env):**
```
VITE_API_URL=http://localhost:8080/api
```

---

# 8. Response Pattern

### Success
```json
{
  "success": true,
  "data": { }
}
```

### Error
```json
{
  "success": false,
  "error": "detailed error message",
  "code": "ERROR_CODE"
}
```

---

# 9. Security

### Backend
- **JWT**: Access token 15 min, Refresh token 7 dias
- **Password**: Bcrypt com salt
- **CORS**: Configurado para frontend local
- **Rate Limiting**: 100 requests/min por IP
- **Validação**: Client-side (React) + Server-side (Go)

### Frontend
- **localStorage**: Armazena JWT seguramente
- **HTTPS Ready**: Em produção, usar HTTPS
- **XSS Protection**: React escapa HTML automaticamente
- **CSRF**: Backend valida origin

---

# 10. Dockerização

## docker-compose.yml
```yaml
services:
  api:
    build: .
    ports: 8080:8080
    environment:
      - DB_HOST=postgres
      - REDIS_HOST=redis
    depends_on:
      - postgres
      - redis
  
  postgres:
    image: postgres:16
    ports: 5432:5432
  
  redis:
    image: redis:7
    ports: 6379:6379
```

## Frontend (npm)
```bash
npm run dev      # Desenvolvimento (Vite)
npm run build    # Build produção
npm run preview  # Preview build
```

---

# 11. Fluxo de Desenvolvimento

### Fase 1-6: Backend
- Setup, autenticação, CRUD, batalhas, ranking, segurança

### Fase 7: Frontend Setup
- React + Vite + Auth Context + Router

### Fase 8: UI Completa
- Components, pages, integração API

### Fase 9: Polish
- Toasts, loading states, error handling

---

# 12. Escalabilidade

O projeto está preparado para:
- Separar frontend e backend em repos distintos
- Deploy independente (Frontend: Vercel/Netlify, Backend: Heroku/AWS)
- Horizontal scaling com load balancer
- Cache distribuído com Redis
- Filas assíncronas (futuro)

---

# 13. Futuras Melhorias

- Testes automatizados (Jest, Go testing)
- CI/CD (GitHub Actions)
- WebSocket para batalhas em tempo real
- Sistema de guildas
- Inventário e equipamentos
- Logs estruturados (ELK Stack)
- Observabilidade (Prometheus, Grafana)
- Deploy em Kubernetes

---

# 14. Objective Final

Construir uma **aplicação fullstack profissional** que demonstra:
- Backend robusto em Go
- Frontend moderno em React
- Integração perfeita entre camadas
- Pronto para ambiente de produção
