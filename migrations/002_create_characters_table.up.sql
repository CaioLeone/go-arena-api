CREATE TABLE IF NOT EXISTS characters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    -- CHAVE ESTRANGEIRA
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    -- DADOS DO PERSONAGEM
    name VARCHAR(50) UNIQUE NOT NULL,
    class VARCHAR(50) NOT NULL,

    -- STATS DO PERSONAGEM(VALORES INICIAIS PADRÃO)
    level INT DEFAULT 1 CHECK (level >= 1 AND level <= 100),
    hp INT DEFAULT 100 CHECK (hp > 0),
    attack INT DEFAULT 10 CHECK (attack > 0),
    defence INT DEFAULT 5 CHECK (defence > 0), 

    -- RANKING 
    ranking_points INT DEFAULT 0 CHECK (ranking_points >= 0),

    -- TIMESTAMPS
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- INDICES PARA CONSULTAS FREQUENTES
CREATE INDEX IF NOT EXISTS idx_characters_user_id ON characters(user_id);
CREATE INDEX IF NOT EXISTS idx_characters_user_name ON characters(name);
CREATE INDEX IF NOT EXISTS idx_characters_user_ranking_points ON characters(ranking_points DESC);

COMMENT ON TABLE characters IS 'Personsagens dos usuarios';
COMMENT ON TABLE characters.user_id IS 'FK para usuario proprietario';
COMMENT ON TABLE characters.class IS 'Classe: Barbaro, Mago, Arqueiro, Assassino';
COMMENT ON TABLE characters.hp IS 'Health Points (vida do personagem)'; 
COMMENT ON TABLE characters.ranking_points IS 'Pontos acumulados em batalha';