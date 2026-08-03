import {useEffect, useState} from 'react';
import characterService from '../../services/characterService';
import type { Character } from '../../types/character';
import CharacterCard from './CharacterCard';

const [characters, setCharacters] = useState<Character[]>([]);
const [loading, setLoading] = useState(true);

export default function CharacterList() {
    useEffect(() => {
        loadCharacters();
    }, []);

    async function loadCharacters() {
        try {
            const data = await characterService.getAll();
            setCharacters(data);
        }finally {
            setLoading(false);
        }
    }

    if (loading) {
        return <p>Carregando Personagens...</p>;
    }

    if (characters.length === 0) {
        return (<p>Nenhum personagem encontrado.</p>);
    }

    return (
        <>
            {characters.map((character) => 
                <CharacterCard key={character.id} character={character} />
            )}
        </>
    );
}

