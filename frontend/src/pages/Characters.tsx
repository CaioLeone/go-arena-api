import { useState, } from "react";

import DashboardLayout from "../components/layout/DashboardLayout";
import CharacterList from "../components/characters/CharacterList";
import CreateCharacterModal from "../components/characters/CreateCharacterModal";

export default function Characters() {

    const [openModal, setOpenModal] = useState(false);
    const [refreshKey, setRefreshKey] = useState(0);

    function handleCreated() {
        setOpenModal(false);

        setRefreshKey(
            (current) => current + 1
        );
    }

    return (
        <DashboardLayout>
            <div
                className="
                    mb-6
                    flex
                    flex-col
                    gap-4
                    sm:flex-row
                    sm:items-center
                    sm:justify-between
                "
            >

                <div>
                    <h1
                        className="
                            text-3xl
                            font-bold
                            text-white
                        "
                    >
                        Meus Personagens
                    </h1>

                    <p
                        className="
                            mt-1
                            text-slate-400
                        "
                    >
                        Gerencie seus guerreiros da arena.
                    </p>
                </div>

                <button
                    onClick={() =>
                        setOpenModal(true)
                    }
                    className="
                        rounded-lg
                        bg-amber-500
                        px-4
                        py-2
                        font-semibold
                        text-slate-950
                        transition
                        hover:bg-amber-400
                    "
                >
                    Novo Personagem
                </button>
            </div>

            <CharacterList
                refreshKey={refreshKey}
            />

            <CreateCharacterModal
                open={openModal}
                onClose={() => setOpenModal(false)}
                onCreated={handleCreated}
            />
        </DashboardLayout>
    );
}