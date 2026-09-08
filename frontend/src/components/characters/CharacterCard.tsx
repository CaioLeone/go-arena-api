import { Link } from "react-router-dom";

import type { Character } from "../../types/character";

interface Props {
    character: Character;
}

export default function CharacterCard({
    character,
}: Props) {
    return (
        <Link
            to={`/characters/${character.id}`}
            className="
                block
                rounded-xl
                border
                border-slate-700
                bg-slate-900
                p-5
                transition
                hover:-translate-y-1
                hover:border-amber-500
                hover:shadow-lg
            "
        >

            <div
                className="
                    mb-4
                    flex
                    items-start
                    justify-between
                    gap-3
                "
            >

                <div>
                    <h2
                        className="
                            text-xl
                            font-bold
                            text-white
                        "
                    >
                        {character.name}
                    </h2>

                    <p
                        className="
                            text-sm
                            text-amber-400
                        "
                    >
                        {character.class}
                    </p>
                </div>

                <span
                    className="
                        rounded-full
                        bg-slate-800
                        px-3
                        py-1
                        text-xs
                        font-medium
                        text-slate-300
                    "
                >
                    Nv. {character.level}
                </span>

            </div>

            <div
                className="
                    grid
                    grid-cols-2
                    gap-3
                    text-sm
                    sm:grid-cols-4
                "
            >

                <div>
                    <span className="text-slate-400">
                        Vida
                    </span>
                    <p className="font-semibold">
                        {character.hp}
                    </p>
                </div>

                <div>
                    <span className="text-slate-400">
                        Ataque
                    </span>

                    <p className="font-semibold">
                        {character.attack}
                    </p>
                </div>

                <div>
                    <span className="text-slate-400">
                        Defesa
                    </span>

                    <p className="font-semibold">
                        {character.defense}
                    </p>
                </div>

                <div>
                    <span className="text-slate-400">
                        Crítico
                    </span>

                    <p className="font-semibold">
                        {character.critical_chance}%
                    </p>
                </div>
            </div>

            <div
                className="
                    mt-4
                    flex
                    justify-between
                    border-t
                    border-slate-800
                    pt-4
                    text-sm
                    text-slate-400
                "
            >
                <span>
                    XP: {character.experience}
                </span>

                <span>
                    Pontos:{" "}
                    {character.attribute_points}
                </span>
            </div>
        </Link>
    );
}