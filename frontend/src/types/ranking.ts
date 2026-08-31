export interface UserRanking {
    character_id: string;
    name: string;
    class: string;
    level: number;
    rank: number;
    score: number;
}

export interface TopPlayer {
    rank: number;
    character_id: string;
    name: string;
    class: string;
    level: number;
    score: number;
}

export interface LeaderboardResponse {
    players: TopPlayer[];
    total: number;
}