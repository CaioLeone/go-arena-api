# FASES DE IMPLEMENTAÇÃO

## FASE 1: SETUP INICIAL (1-2h)
* Objetivo: Projeto compilável + Docker rodando

### Tasks:

* Corrigir packages (models e DTOs)
* Criar estrutura de diretórios
* Mover main.go para /cmd/api/
* Setup go.mod com deps essenciais
* Criar docker-compose.yml (PostgreSQL + Redis)
* Criar .env template
* Criar Dockerfile (simples, com Air para dev)
* Criar migrations SQL básicas (users, characters, battles)
* Implementar cmd/api/main.go MINIMALISTA:
* Carregar config
* Conectar PostgreSQL
* Conectar Redis
* Iniciar servidor (porta 8080)
* Rotas básicas (apenas /health)

### Entregáveis:
* [] Projeto compila com go build
* Docker compila com docker-compose up
* Endpoint /health retorna 200 OK
* Database criado com migrations

## FASE 2: AUTENTICAÇÃO MINIMALISTA (2h)
* Objetivo: Login/Register funcional + JWT middleware

### Tasks:

* Implementar internal/repository/user_repository.go
* GetUserByEmail()
* CreateUser()
* GetUserByID()
* Implementar internal/service/auth_service.go
* Register()
* Login()
* ValidateToken()
* Implementar internal/handler/auth_handler.go
* POST /auth/register
* POST /auth/login
* Implementar internal/middleware/jwt_middleware.go
* Implementar internal/auth/token.go

### Entregáveis:
* POST /auth/register funciona
* POST /auth/login retorna JWT
* Middleware protege rotas privadas

## FASE 3: CRUD DE PERSONAGENS (1.5h)
* Objetivo: Criar, ler, editar, deletar personagens

### Tasks:

* Implementar internal/repository/character_repository.go
* Implementar internal/service/character_service.go
* Implementar internal/handler/character_handler.go

### Endpoints:

* POST /characters
* GET /characters
* GET /characters/:id
* PUT /characters/:id
* DELETE /characters/:id

### Entregáveis:

* CRUD completo de personagens
* Apenas usuário autenticado pode acessar

## FASE 4: SISTEMA DE BATALHA (1.5h)
* Objetivo: Iniciar e resolver batalhas PvP

### Tasks:
* Implementar internal/battle/calculator.go
* CalcDamage() (fórmula: Dano = Ataque - Defesa/2)
* DetermineBattle()
* UpdateRanking()
* Implementar internal/repository/battle_repository.go
* Implementar internal/service/battle_service.go
* Implementar internal/handler/battle_handler.go

### Endpoints:
* POST /battles (inicia batalha)
* GET /battles/history

### Entregáveis:
* Batalhas funcionam
* Ranking atualiza
* Histórico registra tudo

## FASE 5: RANKING + LEADERBOARD (REDIS) (1h)
* Objetivo: Leaderboard em tempo real

### Tasks:

* Implementar pkg/redis/client.go
* Implementar internal/ranking/leaderboard.go
* Implementar internal/handler/ranking_handler.go

### Endpoints:

* GET /ranking
* GET /ranking/top

### Entregáveis:

* Redis atualiza após cada batalha
* Leaderboard consultável

## FASE 6: SEGURANÇA + HARDENING (1h)
* Objetivo: Rate limiting, logs, error handling

### Tasks:

* Implementar internal/middleware/rate_limit.go
* Implementar pkg/logger/logger.go
* Melhorar error handling global

### Entregáveis:

* Rate limiting funciona
* Logs básicos
* Erros são tratados graciosamente

## FASE 7: SETUP REACT + AUTENTICAÇÃÕ FRONTEND(2H)
* Objetivo: setup, autenticação, CRUD, batalhas, ranking, Segurança

### Task:
* Create React App (ou Vite para mais rapido)
* Setup .env com API_URL
* Criar src/services/api.ts (axios com interceptor)
* Criar COntexto de autenticação (src/context/AuthContext.tsx)
* Criar pagina de Login (src/pages/Login.tsx)
* Criar pagina de Register (src/pages/Register.tsx)
* Setup routing basico (React Router v6)
* Proteger rotas com PrivateRoute

### Entregaveis 
* React compila com npm start
* Login/Register funcionam
* JWT armazenado em localStorage
* Rotas privadas protegidas

## FASE 8: UI - PERSONAGENS + BATALHAS (2.5H)
* Objetivo: Interface completa para gerenciar personagens e batalhas.

### Tasks
* Criar Componente CharacterList (listar personagens)
* Criar Modal CreateCharacter
* Criar Pagina CharacterDetail (stats, editar, deletar)
* Criar Pagina Battles (iniciar batalhas, ver historico)
* Criar Componente BattleResult (mostrar resultado)
* Criar Pagina LeaderBoard (top players)
* Setup TailWindCSS para estilo rapido (KISS)

### Entregaveis
* Dashboard funcional
* CRUD de personagens via UI
* Sistema de batalhas via UI
* LearderBoard em tempo real

## FASE 9: MELHORIAS UX + Deploy (1.5h)
* Objetivo: Polish e deploy

### Task
* Toast/Alert para feedback do usuario
* Loading states nos componentes
* Error Handling com mensagens
* Responsividade mobile basica
* Build React para produção
* Documentação no Readme

### Entregaveis
* app polido e pronto
* Deploy em localhost testado
* Readme com instruções