import api from "./api";

import type {
    LeaderboardResponse,
    UserRanking,
} from "../types/ranking";

async function getTopPlayers(): Promise<LeaderboardResponse> {
    const response = await api.get("/ranking/top");

    return response.data.data;
}

async function getUserRanking(): Promise<UserRanking> {
    const response = await api.get("/ranking");

    return response.data.data;
}

const rankingService = {
    getTopPlayers,
    getUserRanking,
};

export default rankingService;