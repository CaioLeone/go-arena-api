export interface BattleCreateRequest {
    attacker_character_id: string;
    defender_character_id: string;
}

export interface BattleRound {
    round: number;
    attacker_name: string;
    defender_name: string;
    damage: number;
    remaining_hp: number;
    is_critical: boolean;
    message: string;
}

export interface Battle {
    id: string;

    attacker_id: string;
    attacker_name: string;

    defender_id: string;
    defender_name: string;

    winner_id: string | null;
    winner_name: string;

    damage_dealt: number;

    attacker_hp_final: number;
    defender_hp_final: number;

    rounds_count: number;
    rounds: BattleRound[];

    created_at: string;
}

export interface BattleHistory {
    id: string;

    attacker_id: string;
    attacker_name: string;

    defender_id: string;
    defender_name: string;

    winner_id: string | null;
    winner_name: string;

    damage_dealt: number;

    attacker_hp_final: number;
    defender_hp_final: number;

    rounds_count: number;

    created_at: string;
}