import api from './api';

import type {
    Character,
    CreateCharacterRequest,
    UpdateCharacterRequest,
    AddExperienceRequest,
    SpendAttributeRequest
} from "../types/character";

async function getAll(): Promise<Character[]> {
    const response = await api.get('/characters');
    return response.data.data.characters;
}

async function getById(id: string): Promise<Character> {
    const response = await api.get(`/characters/${id}`);
    return response.data.data;
}

async function create(data: CreateCharacterRequest): Promise<Character> {
    const response = await api.post('/characters', data);
    return response.data.data;
}

async function update(id: string, data: UpdateCharacterRequest): Promise<Character>{
    const response = await api.put(`/characters/${id}`, data);
    return response.data.data;
}

async function remove(id: string): Promise<void> {
    await api.delete(`/characters/${id}`);
}

async function addExperience(id: string, data: AddExperienceRequest): Promise<Character> {
    const response = await api.post(`/characters/${id}/experience`, data);
    return response.data.data;
}

async function spendAttribute(id: string, data: SpendAttributeRequest): Promise<Character> {
    const response = await api.post(`/characters/${id}/attributes`, data);
    return response.data.data;
}

const characterService = {
    getAll,
    getById,
    create,
    update,
    remove,
    addExperience,
    spendAttribute
};

export default characterService;