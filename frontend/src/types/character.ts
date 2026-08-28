export interface Character{
    id: string;
    user_id: string;
    name: string;
    class: string;
    level: number;
    experience: number;
    hp: number;
    attack: number;
    defense: number;
    attribute_points: number;
    ranking_points: number;
    critical_chance: number;
    created_at: string;
    updated_at: string;
}

export interface CharacterListResponse {
    characters: Character[];
    total: number;
}

export interface CreateCharacterRequest {
    name: string;
    class: string;
}

export interface UpdateCharacterRequest {
    name: string;
}

export interface AddExperienceRequest {
    experience: number;
}

export interface SpendAttributeRequest {
    hp?: number;
    attack?: number;
    defense?: number;
    critical_chance?: number;
}