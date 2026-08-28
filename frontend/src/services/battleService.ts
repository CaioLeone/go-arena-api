import api from "./api";
import type { Battle, BattleCreateRequest, BattleHistory } from "../types/battle";

async function start(data: BattleCreateRequest): Promise<Battle> {
    const response = await api.post('/battles', data);
    return response.data.data;
}

async function getHistory(): Promise<BattleHistory[]> {
    const response = await api.get('/battles/history');
    return response.data.data;
}

const battleService = {
    start,
    getHistory,
};

export default battleService;