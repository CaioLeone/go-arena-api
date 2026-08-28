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
            const charactersData =
                await characterService.getAll();

            const historyData =
                await battleService.getHistory();

            setCharacters(charactersData);
            setHistory(historyData);
        } catch {
            setError(
                "Erro ao carregar dados das batalhas."
            );
        } finally {
            setLoading(false);
        }
    }

    async function handleBattle() {
        setError("");

        if (!attackerId || !defenderId) {
            setError(
                "Escolha os dois personagens."
            );
            return;
        }

        if (attackerId === defenderId) {
            setError(
                "Escolha personagens diferentes."
            );
            return;
        }

        try {
            setBattleLoading(true);

            const result =
                await battleService.start({
                    attacker_character_id:
                        attackerId,

                    defender_character_id:
                        defenderId,
                });

            setBattleResult(result);

            const updatedHistory =
                await battleService.getHistory();

            setHistory(updatedHistory);
        } catch {
            setError(
                "Não foi possível iniciar a batalha."
            );
        } finally {
            setBattleLoading(false);
        }
    }

    if (loading) {
        return (
            <DashboardLayout>
                <p>
                    Carregando batalhas...
                </p>
            </DashboardLayout>
        );
    }

    return (
        <DashboardLayout>
            <h1>Batalhas</h1>

            {error && (
                <p
                    style={{
                        color: "red",
                    }}
                >
                    {error}
                </p>
            )}

            <section>
                <h2>Iniciar Batalha</h2>

                <div>
                    <label>
                        Atacante
                    </label>

                    <br />

                    <select
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

                        {characters.map(
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

                <br />

                <div>
                    <label>
                        Defensor
                    </label>

                    <br />

                    <select
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

                        {characters.map(
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

                <br />

                <button
                    onClick={handleBattle}
                    disabled={battleLoading}
                >
                    {battleLoading
                        ? "Batalhando..."
                        : "Iniciar Batalha"}
                </button>
            </section>

            {battleResult && (
                <BattleResult
                    battle={battleResult}
                />
            )}

            <hr />

            <section>
                <h2>Histórico</h2>

                {history.length === 0 && (
                    <p>
                        Nenhuma batalha realizada.
                    </p>
                )}

                {history.map((battle) => (
                    <div
                        key={battle.id}
                        style={{
                            border:
                                "1px solid #ddd",
                            padding: 15,
                            marginBottom: 10,
                            borderRadius: 6,
                        }}
                    >
                        <strong>
                            {battle.attacker_name}
                            {" vs "}
                            {battle.defender_name}
                        </strong>

                        <p>
                            Vencedor:{" "}
                            {battle.winner_name}
                        </p>

                        <p>
                            Dano total:{" "}
                            {battle.damage_dealt}
                        </p>

                        <p>
                            HP final de{" "}
                            {battle.attacker_name}:{" "}
                            {
                                battle.attacker_hp_final
                            }
                        </p>

                        <p>
                            HP final de{" "}
                            {battle.defender_name}:{" "}
                            {
                                battle.defender_hp_final
                            }
                        </p>

                        <p>
                            Rounds:{" "}
                            {battle.rounds_count}
                        </p>

                        <p>
                            {new Date(
                                battle.created_at
                            ).toLocaleString(
                                "pt-BR"
                            )}
                        </p>
                    </div>
                ))}
            </section>
        </DashboardLayout>
    );
}