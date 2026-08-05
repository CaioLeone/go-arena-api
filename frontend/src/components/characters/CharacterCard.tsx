import type { Character } from "../../types/character";
import {Link} from "react-router-dom";

interface Props{
    character: Character;
}

export default function CharacterCard({ character }: Props){
    return (
        <Link to={`/characters/${character.id}`}>
            <div
                style={{
                    border: "1px solid #ccc",
                    borderRadius: 8,
                    padding: 20,
                    marginBottom: 15,
                    background: "#fff",
                }}
            >
                <h2 className="text-xl font-bold">{character.name}</h2>
                <p>Classe: {character.class}</p>
                <p>Nivel: {character.level}</p>

                <hr className="my-2"/>
                
                <p>HP: {character.hp}</p>
                <p>Ataque: {character.attack}</p>
                <p>Defesa: {character.defense}</p>
                <p>Chance de Critico: {character.critical_chance}</p>
                
                <hr className="my-2"/>
                
                <p>Experiencia: {character.experience}</p>
                <p>Pontos de Atributo: {character.attribute_points}</p>
            </div>
        </Link>
        
    );
}