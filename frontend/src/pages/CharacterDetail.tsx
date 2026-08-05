import {useEffect, useState} from "react";
import { useNavigate, useParams } from "react-router-dom";
import characterService from "../services/characterService";
import type { Character } from "../types/character";

export default function CharacterDetail() {
    const { id } = useParams();
    const navigate = useNavigate();
    const [character, setCharacter] = useState<Character | null>(null);
    const [loading, setLoading] = useState(true);
    const [name, setName] = useState("");

    useEffect(() => {
        if(id) {
            loadCharacter();
        }
    }, [id]);

    async function loadCharacter() {
        try{
            const data = await characterService.getById(id!);
            setCharacter(data);
            setName(data.name);
        }finally{
            setLoading(false);
        }
    }

    async function handleUpdate(){
        if(!character) return;
        const updated = await characterService.update(character.id,{name,});
        setCharacter(updated);
    }

    async function handleDelete(){
        if(!character) return;
        if(!confirm("Deseja excluir esse personagem?")) return;

        await characterService.remove(character.id);
        navigate("/characters");
    }

    async function handleAddXp(){
        if(!character) return;
        const updated = await characterService.addExperience(character.id,{ experience: 10 });
        setCharacter(updated);
    }

    async function handleSpendPoints(attribute: string){
        if(!character) return;
        const updated = await characterService.spendAttribute(character.id,{ attribute });
        setCharacter(updated);
    }

    if(loading){
        return <p>Carregando...</p>;
    }

    if(!character){
        return <p>Personagem não encontrado</p>;
    }

    return(
        <div className="max-w-3xl mx-auto space-y-6">
            <h1 className="text-3xl font-bold">
                {character.name}
            </h1>
            <div className="bg-white rounded shadow p-6 space-y-4">
                <div>
                    <label>Nome</label>
                    <input
                        className="border rounded w-full p-2"
                        value={name}
                        onChange={(e)=>setName(e.target.value)}
                    />
                </div>

                <button
                    onClick={handleUpdate}
                    className="bg-blue-600 text-white px-4 py-2 rounded"
                >
                    Salvar
                </button>
            </div>

            <div className="bg-white rounded shadow p-6">
                <h2 className="font-bold mb-4">
                    Informações
                </h2>
                <p>Nível: {character.level}</p>
                <p>XP: {character.experience}</p>
                <p>Pontos: {character.attribute_points}</p>
            </div>

            <div className="bg-white rounded shadow p-6">
                <h2 className="font-bold mb-4">
                    Atributos
                </h2>
                <p>
                    Vida
                    {" "}
                    {character.hp}
                    <button
                        className="ml-2 bg-green-600 text-white px-2 rounded"
                        onClick={()=>handleSpendPoints("hp")}
                    >
                        +
                    </button>
                </p>
                <p>
                    Ataque
                    {" "}
                    {character.attack}
                    <button
                        className="ml-2 bg-green-600 text-white px-2 rounded"
                        onClick={()=>handleSpendPoints("attack")}
                    >
                        +
                    </button>
                </p>

                <p>
                    Defesa
                    {" "}
                    {character.defense}
                    <button
                        className="ml-2 bg-green-600 text-white px-2 rounded"
                        onClick={()=>handleSpendPoints("defense")}
                    >
                        +
                    </button>
                </p>

                <p>
                    Critico
                    {" "}
                    {character.c}
                    <button
                        className="ml-2 bg-green-600 text-white px-2 rounded"
                        onClick={()=>handleSpendPoints("defense")}
                    >
                        +
                    </button>
                </p>
            </div>

            <div className="flex gap-4">
                <button
                    onClick={handleAddXp}
                    className="bg-yellow-500 text-white px-4 py-2 rounded"
                >
                    +10 XP
                </button>

                <button
                    onClick={handleDelete}
                    className="bg-red-600 text-white px-4 py-2 rounded"
                >
                    Excluir
                </button>
            </div>
        </div>
    );
}