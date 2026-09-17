import {useEffect, useState} from 'react';
import characterService from '../../services/characterService';
import type { Character } from '../../types/character';
import CharacterCard from './CharacterCard';
import CreateCharacterModal from './CreateCharacterModal';

interface Props {
    refreshKey: number;
}

export default function CharacterList({refreshKey,}: Props) {
    const [characters, setCharacters] = useState<Character[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState("");

    useEffect(() => { loadCharacters(); }, [refreshKey]);

    async function loadCharacters() {
        try {
            setLoading(true);
            setError("");

            const data = await characterService.getAll();

            setCharacters(data);

        } catch (error) {
            console.error("Erro ao carregar personagens:", error);

            setCharacters([]);

            setError("Erro ao carregar personagens.");

        } finally {
            setLoading(false);
        }
    }

    if (loading) {
        return (
            <p className="text-slate-400">
                Carregando personagens...
            </p>
        );
    }

    if (error) {
        return (
            <p className="arena-error">
                {error}
            </p>
        );
    }

    if (characters.length === 0) {
        return (
            <p className="text-slate-400">
                Nenhum personagem encontrado.
            </p>
        );
    }

    return (
        <div
            className="
                grid
                grid-cols-1
                gap-4
                md:grid-cols-2
                xl:grid-cols-3
            "
        >
            {characters.map(
                (character) => (
                    <CharacterCard
                        key={character.id}
                        character={character}
                    />
                )
            )}
        </div>
    );
}

