import { useState } from "react";
import characterService from "../../services/characterService";

interface Props {
    open: boolean;
    onClose: () => void;
    onCreated: () => void;
}

export default function CreateCharacterModal({
    open,
    onClose,
    onCreated,
}: Props) {
    const [name, setName] = useState("");
    const [characterClass, setCharacterClass] = useState("Barbaro");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    if (!open) {
        return null;
    }

    async function handleSubmit(
        e: React.FormEvent
    ) {
        e.preventDefault();

        try {
            setLoading(true);
            setError("");

            await characterService.create({
                name,
                class: characterClass,
            });

            setName("");
            setCharacterClass("Barbaro");

            onCreated();
        } catch {
            setError(
                "Erro ao criar personagem."
            );
        } finally {
            setLoading(false);
        }
    }

    return (
        <div
            className="
                fixed
                inset-0
                z-50
                flex
                items-center
                justify-center
                bg-black/70
                p-4
            "
        >

            <div
                className="
                    w-full
                    max-w-md
                    rounded-xl
                    border
                    border-slate-700
                    bg-slate-900
                    p-6
                    shadow-2xl
                "
            >

                <h2
                    className="
                        mb-6
                        text-2xl
                        font-bold
                        text-white
                    "
                >
                    Criar Personagem
                </h2>
                <form
                    onSubmit={handleSubmit}
                    className="space-y-5"
                >

                    <div>
                        <label
                            className="
                                mb-2
                                block
                                text-sm
                                font-medium
                                text-slate-300
                            "
                        >
                            Nome
                        </label>

                        <input
                            value={name}
                            onChange={(e) =>
                                setName(
                                    e.target.value
                                )
                            }
                            required
                            className="
                                w-full
                                rounded-lg
                                border
                                border-slate-700
                                bg-slate-800
                                px-3
                                py-2
                                text-white
                                outline-none
                                transition
                                focus:border-amber-500
                            "
                        />

                    </div>

                    <div>
                        <label
                            className="
                                mb-2
                                block
                                text-sm
                                font-medium
                                text-slate-300
                            "
                        >
                            Classe
                        </label>

                        <select
                            value={
                                characterClass
                            }
                            onChange={(e) =>
                                setCharacterClass(
                                    e.target.value
                                )
                            }
                            className="
                                w-full
                                rounded-lg
                                border
                                border-slate-700
                                bg-slate-800
                                px-3
                                py-2
                                text-white
                                outline-none
                                focus:border-amber-500
                            "
                        >
                            <option value="Barbaro">
                                Bárbaro
                            </option>

                            <option value="Mago">
                                Mago
                            </option>

                            <option value="Arqueiro">
                                Arqueiro
                            </option>

                            <option value="Assassino">
                                Assassino
                            </option>
                        </select>

                    </div>

                    {error && (
                        <p
                            className="
                                rounded-lg
                                bg-red-950
                                p-3
                                text-sm
                                text-red-300
                            "
                        >
                            {error}
                        </p>
                    )}

                    <div
                        className="
                            flex
                            justify-end
                            gap-3
                        "
                    >

                        <button
                            type="button"
                            onClick={onClose}
                            className="
                                rounded-lg
                                border
                                border-slate-600
                                px-4
                                py-2
                                text-slate-300
                                transition
                                hover:bg-slate-800
                            "
                        >
                            Cancelar
                        </button>

                        <button
                            type="submit"
                            disabled={loading}
                            className="
                                rounded-lg
                                bg-amber-500
                                px-4
                                py-2
                                font-semibold
                                text-slate-950
                                transition
                                hover:bg-amber-400
                                disabled:opacity-50
                            "
                        >
                            {loading
                                ? "Criando..."
                                : "Criar"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}