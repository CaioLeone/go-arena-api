import {useState} from "react";
import  characterService from "../../services/characterService";

interface Props {
    open: boolean;
    onClose: () => void;
    onCreated: () => void;
}

export default function CreateCharacterModal({ open, onClose, onCreated }: Props) {
    const [name, setName] = useState("");
    const [characterClass, setCharacterClass] = useState("Barbaro");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    if(!open) return null;

    async function handleSubmit(e: React.FormEvent) {
        e.preventDefault();

        try{
            setLoading(true);
            setError("");

            await characterService.create({name, class: characterClass});
            setName("");
            setCharacterClass("Barbaro");
            onCreated();
            onClose();
        }catch{
            setError("Erro ao criar personagem");
        }finally{
            setLoading(false);
        }
    }
    return (
        <div className="fixed inset-0 bg-black/50 flex items-center justify-center">
            <div className="bg-white rounded-lg p-6 w-96">
                <h2 className="text-xl font-bold mb-4">
                    Criar Personagem
                </h2>

                <form 
                    onSubmit={handleSubmit}
                    className="space-y-4"
                >
                    <div>
                        <label className="block mb-1">
                            Nome
                        </label>
                        <input
                            className="border w-full p-2 rounded"
                            value={name}
                            onChange={(e)=>setName(e.target.value)}
                            required
                        />
                    </div>

                    <div>
                        <label className="block mb-1">
                            Classe
                        </label>
                        <select
                            className="border w-full p-2 rounded"
                            value={characterClass}
                            onChange={(e)=>setCharacterClass(e.target.value)}
                        >
                            <option value="Barbaro">Bárbaro</option>
                            <option value="Mago">Mago</option>
                            <option value="Arqueiro">Arqueiro</option>
                            <option value="Assassino">Assassino</option>
                        </select>
                    </div>

                    {error && (
                        <p className="text-red-500">{error}</p>
                    )}

                    <div className="flex justify-end gap-2">
                        <button
                            type="button"
                            onClick={onClose}
                            className="px-4 py-2 border rounded"
                        >
                            Cancelar
                        </button>

                        <button
                            type="submit"
                            disabled={loading}
                            className="bg-blue-600 text-white px-4 py-2 rounded"
                        >
                            {loading ? "Criando..." : "Criar"}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
}