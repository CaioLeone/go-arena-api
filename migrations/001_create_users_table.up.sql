CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

COMMENT ON TABLE users IS 'Tabela de usuarios registrados no sistema';
COMMENT ON COLUMN users.id IS 'Identificador unico (UUID)';
COMMENT ON COLUMN users.email IS 'Email unico do usuario, usado para login';
COMMENT ON COLUMN users.password IS 'Hash bcrypt da senha (nunca texto puro)';