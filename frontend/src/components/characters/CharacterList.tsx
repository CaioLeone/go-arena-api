import {useEffect, useState} from 'react';
import characterService from '../../services/characterService';
import type { Character } from '../../types/character';
import CharacterCard from './CharacterCard';
import CreateCharacterModal from './CreateCharacterModal';

export default function CharacterList() {
    const [characters, setCharacters] = useState<Character[]>([]);
    const [loading, setLoading] = useState(true);
    const [openModal, setOpenModal] = useState(false);
    
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
            <div className="flex justify-between mb-6">

                <h1 className="text-2xl font-bold">
                    Personagens
                </h1>

                <button
                    onClick={()=>setOpenModal(true)}
                    className="bg-green-600 text-white px-4 py-2 rounded"
                >
                    Novo Personagem
                </button>

            </div>
        <CreateCharacterModal
            open={openModal}
            onClose={() => setOpenModal(false)}
            onCreate={loadCharacters}
        />
            {characters.map((character) => 
                <CharacterCard key={character.id} character={character} />
            )}
        </>
    );
}

