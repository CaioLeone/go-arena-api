export interface Character{
    id: string;
    name: string;
    class: string;
    level: number;
    hp: number;
    attack: number;
    defense: number;
    experience: number;
    attribute_points: number;
    created_at: string;
}

export interface CreateCharacterRequest {
    name: string;
    class: string;
}

export interface UpdateCharacterRequest {
    name?: string;
    class?: string;
}

export interface AddExperienceRequest {
    experience: number;
}

export interface SpendAttributeRequest {
    attack?: number;
    defense?: number;
    hp?: number;
}