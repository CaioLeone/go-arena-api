import { useEffect, useState, } from "react";

import DashboardLayout from "../components/layout/DashboardLayout";
import BattleResult from "../components/battles/BattleResult";

import characterService from "../services/characterService";
import battleService from "../services/battleService";

import type { Character } from "../types/character";

import type { Battle, BattleHistory, } from "../types/battle";

export default function Battles() {
    const [characters, setCharacters] = useState<Character[]>([]);
    const [attackerId, setAttackerId] = useState("");
    const [defenderId, setDefenderId] = useState("");
    const [battleResult, setBattleResult] = useState<Battle | null>(null);
    const [history, setHistory] = useState<BattleHistory[]>([]);
    const [loading, setLoading] = useState(true);
    const [battleLoading, setBattleLoading] = useState(false);
    const [error, setError] = useState("");

    useEffect(() => { loadData(); }, []);

    async function loadData() {
        try {
            setLoading(true);
            setError("");
            
            const [charactersData, historyData] = await Promise.all([
                characterService.getAll(),
                battleService.getHistory(),
            ]);
            
            setCharacters(charactersData ?? []);
            setHistory(historyData ?? []);

        } catch (err){
            console.error("Erro ao carregar dados das batalhas:", err);
            
            setCharacters([]);
            setHistory([]);
            
            setError("Erro ao carregar dados das batalhas. Por favor, tente novamente.");
        } finally {
            setLoading(false);
        }
    }

    async function handleBattle() {
        setError("");

        if (!attackerId || !defenderId) {
            setError("Escolha os dois personagens.");
            return;
        }

        if (attackerId === attackerId) {
            setError("Escolha personagens diferentes.");
            return;
        }

        try {
            setBattleLoading(true);

            const result =
                await battleService.start({
                    attacker_character_id: attackerId,
                    defender_character_id: defenderId,
                });

            setBattleResult(result);

            const updatedHistory = await battleService.getHistory();

            setHistory(updatedHistory);
        } catch (err){
            console.error("Erro ao iniciar batalha:", err);
            
            setError("Erro ao iniciar batalha. Por favor, tente novamente.");
        } finally {
            setBattleLoading(false);
        }
    }

    if (loading) {
        return (
            <DashboardLayout>
                <p> Carregando batalhas... </p>
            </DashboardLayout>
        );
    }

    return (
    <DashboardLayout>
        <div className="arena-content">

            {/* Cabeçalho */}
            <div className="mb-8">
                <span className="arena-badge">
                    Arena de Combate
                </span>

                <h1 className="arena-title mt-3">
                    Batalhas
                </h1>

                <p className="arena-subtitle">
                    Escolha seus guerreiros e descubra
                    quem dominará a arena.
                </p>
            </div>

            {error && (
                <p className="arena-error mb-6">
                    {error}
                </p>
            )}

            {/* Iniciar batalha */}
            <section className="arena-card mb-10">
                <h2 className="arena-section-title mb-6">
                    Iniciar Batalha
                </h2>

                <div
                    className="
                        grid
                        grid-cols-1
                        items-end
                        gap-5
                        md:grid-cols-[1fr_auto_1fr]
                    "
                >
                    {/* Atacante */}
                    <div>
                        <label className="arena-label">
                            Atacante
                        </label>

                        <select
                            className="arena-select"
                            value={attackerId}
                            onChange={(e) =>
                                setAttackerId(
                                    e.target.value
                                )
                            }
                        >
                            <option value="">
                                Escolha um personagem
                            </option>

                            {characters
                                .filter(
                                    (character) =>
                                        character.id !==
                                        defenderId
                                )
                                .map(
                                    (character) => (
                                        <option
                                            key={character.id}
                                            value={character.id}
                                        >
                                            {character.name}
                                            {" - "}
                                            {character.class}
                                            {" - Nv. "}
                                            {character.level}
                                        </option>
                                    )
                                )}
                        </select>
                    </div>

                    {/* VS */}
                    <div
                        className="
                            hidden
                            pb-3
                            text-xl
                            font-black
                            text-amber-400
                            md:block
                        "
                    >
                        VS
                    </div>

                    {/* Defensor */}
                    <div>
                        <label className="arena-label">
                            Defensor
                        </label>

                        <select
                            className="arena-select"
                            value={defenderId}
                            onChange={(e) =>
                                setDefenderId(
                                    e.target.value
                                )
                            }
                        >
                            <option value="">
                                Escolha um personagem
                            </option>

                            {characters
                                .filter(
                                    (character) =>
                                        character.id !==
                                        attackerId
                                )
                                .map(
                                    (character) => (
                                        <option
                                            key={character.id}
                                            value={character.id}
                                        >
                                            {character.name}
                                            {" - "}
                                            {character.class}
                                            {" - Nv. "}
                                            {character.level}
                                        </option>
                                    )
                                )}
                        </select>
                    </div>
                </div>

                <button
                    className="
                        arena-button-primary
                        mt-6
                        w-full
                        sm:w-auto
                    "
                    onClick={handleBattle}
                    disabled={battleLoading}
                >
                    {battleLoading
                        ? "Batalhando..."
                        : "⚔ Iniciar Batalha"}
                </button>
            </section>

            {/* Resultado */}
            {battleResult && (
                <div className="mb-10">
                    <BattleResult
                        battle={battleResult}
                    />
                </div>
            )}

            {/* Histórico */}
            <section>
                <div className="mb-5">
                    <h2 className="arena-section-title">
                        Histórico
                    </h2>

                    <p className="arena-subtitle">
                        Últimos combates realizados na arena.
                    </p>
                </div>

                {history.length === 0 ? (
                    <div className="arena-card">
                        <p className="text-slate-400">
                            Nenhuma batalha realizada.
                        </p>
                    </div>
                ) : (
                    <div className="space-y-4">
                        {history.map((battle) => (
                            <div
                                key={battle.id}
                                className="
                                    arena-card
                                    arena-card-hover
                                "
                            >
                                <div
                                    className="
                                        flex
                                        flex-col
                                        gap-4
                                        sm:flex-row
                                        sm:items-center
                                        sm:justify-between
                                    "
                                >
                                    <div>
                                        <div
                                            className="
                                                flex
                                                flex-wrap
                                                items-center
                                                gap-2
                                                text-lg
                                                font-bold
                                            "
                                        >
                                            <span className="text-white">
                                                {battle.attacker_name}
                                            </span>

                                            <span className="text-amber-400">
                                                VS
                                            </span>

                                            <span className="text-white">
                                                {battle.defender_name}
                                            </span>
                                        </div>

                                        <p className="mt-2 text-sm text-slate-400">
                                            Vencedor:{" "}
                                            <strong className="text-amber-400">
                                                {battle.winner_name}
                                            </strong>
                                        </p>
                                    </div>

                                    <div
                                        className="
                                            rounded-xl
                                            border
                                            border-amber-500/20
                                            bg-amber-500/10
                                            px-4
                                            py-2
                                            text-center
                                        "
                                    >
                                        <span
                                            className="
                                                block
                                                text-xs
                                                uppercase
                                                tracking-wide
                                                text-slate-400
                                            "
                                        >
                                            Dano
                                        </span>

                                        <strong className="text-xl text-amber-400">
                                            {battle.damage_dealt}
                                        </strong>
                                    </div>
                                </div>

                                <div
                                    className="
                                        mt-5
                                        grid
                                        grid-cols-2
                                        gap-3
                                        border-t
                                        border-slate-800
                                        pt-4
                                        sm:grid-cols-4
                                    "
                                >
                                    <div>
                                        <span className="text-xs text-slate-500">
                                            HP {battle.attacker_name}
                                        </span>

                                        <p className="font-semibold">
                                            {battle.attacker_hp_final}
                                        </p>
                                    </div>

                                    <div>
                                        <span className="text-xs text-slate-500">
                                            HP {battle.defender_name}
                                        </span>

                                        <p className="font-semibold">
                                            {battle.defender_hp_final}
                                        </p>
                                    </div>

                                    <div>
                                        <span className="text-xs text-slate-500">
                                            Rounds
                                        </span>

                                        <p className="font-semibold">
                                            {battle.rounds_count}
                                        </p>
                                    </div>

                                    <div>
                                        <span className="text-xs text-slate-500">
                                            Data
                                        </span>

                                        <p className="text-sm font-semibold">
                                            {new Date(
                                                battle.created_at
                                            ).toLocaleString(
                                                "pt-BR"
                                            )}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </section>

        </div>
    </DashboardLayout>
);
}