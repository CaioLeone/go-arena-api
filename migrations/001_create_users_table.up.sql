CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50),
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indice para consultas por email (muito usado para login)
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Comentario documentando a tabela
COMMENT ON TABLE users IS 'Table de usuarios registrados no sistema';
COMMENT ON COLUMN users.id IS 'Identificador unico (UUID)';
COMMENT ON COLUMN users.email IS 'Email unci do usuario, usado para login';
COMMENT ON COLUMN users.password IS 'Hash bcrypt da senha (nunca texto puro)';