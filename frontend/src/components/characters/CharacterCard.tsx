import type { Character } from "../../types/character";

interface Props{
    character: Character;
}

export default function CharacterCard({ character }: Props){
    return (
        <div
            style={{
                border: "1px solid #ccc",
                borderRadius: 8,
                padding: 20,
                marginBottom: 15,
                background: "#fff",
            }}
        >
            <h2>{character.name}</h2>
            <p>Class: {character.class}</p>
            <p>Level: {character.level}</p>
            <p>HP: {character.hp}</p>
            <p>Attack: {character.attack}</p>
            <p>Defense: {character.defense}</p>
            <p>Experience: {character.experience}</p>
            <p>Attribute Points: {character.attribute_points}</p>
        </div>
    );
}