CREATE TABLE IF NOT EXISTS battles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    -- CHAVES ESTRANGEIRAS( QUEM ATACOU, QUEM DEFENDEU, QUEM VENCEU)
    attacker_id UUID NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    attacker_name VARCHAR(100) NOT NULL,

    defender_id UUID NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    defender_name VARCHAR(100) NOT NULL,

    winner_id UUID NOT NULL REFERENCES characters(id) ON DELETE CASCADE,
    winner_name VARCHAR(100),

    attacker_hp_final INT NOT NULL,
    defender_hp_final INT NOT NULL,

    -- DETALHES
    rounds_count INT NOT NULL,
    rounds_data JSONB NOT NULL,

    -- RESULTADO DA BATALHA
    damage_dealt INT NOT NULL CHECK (damage_dealt > 0),

    --TIMESTAMP
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_battles_attacker_id ON battles(attacker_id);
CREATE INDEX IF NOT EXISTS idx_battles_defender_id ON battles(defender_id);
CREATE INDEX IF NOT EXISTS idx_battles_winner_id ON battles(winner_id);
CREATE INDEX IF NOT EXISTS idx_battles_created_at ON battles(created_at DESC);

-- indice composto para buscar batalhas de um personagem (atacou ou defendeu)
CREATE INDEX IF NOT EXISTS idx_battles_participant ON battles((attacker_id),(defender_id), created_at DESC);

COMMENT ON TABLE battles IS 'Historico de batalhas entre personagens';
COMMENT ON COLUMN battles.attacker_id IS 'FK para personagem atacante';
COMMENT ON COLUMN battles.defender_id IS 'FK para personagem defensor';
COMMENT ON COLUMN battles.winner_id IS 'FK para personagem vencedor';
COMMENT ON COLUMN battles.damage_dealt IS 'Dano total causado na batalha';
COMMENT ON COLUMN battles.rounds_data IS 'Dados em JSON de cada round da batalha';