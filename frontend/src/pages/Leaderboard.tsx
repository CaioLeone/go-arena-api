import { useEffect, useState, } from "react";

import DashboardLayout from "../components/layout/DashboardLayout";
import rankingService from "../services/rankingService";

import type { TopPlayer, UserRanking,} from "../types/ranking";

export default function LeaderBoard() {
    const [players, setPlayers] = useState<TopPlayer[]>([]);

    const [userRanking, setUserRanking] = useState<UserRanking | null>(null);

    const [total, setTotal] = useState(0);

    const [loading, setLoading] = useState(true);

    const [updating, setUpdating] = useState(false);

    const [error, setError] = useState("");

    useEffect(() => {
        loadRanking();
    }, []);

    async function loadRanking() {
        try {
            setError("");

            const leaderboard = await rankingService.getTopPlayers();

            setPlayers(leaderboard.players);
            setTotal(leaderboard.total);

            try {
                const myRanking = await rankingService.getUserRanking();

                setUserRanking(myRanking);
            } catch {
                setUserRanking(null);
            }
        } catch {
            setError(
                "Erro ao carregar o ranking."
            );
        } finally {
            setLoading(false);
        }
    }

    async function handleRefresh() {
        try {
            setUpdating(true);
            setError("");

            const leaderboard = await rankingService.getTopPlayers();

            setPlayers(leaderboard.players);
            setTotal(leaderboard.total);

            try {
                const myRanking = await rankingService.getUserRanking();

                setUserRanking(myRanking);
            } catch {
                setUserRanking(null);
            }
        } catch {
            setError(
                "Erro ao atualizar o ranking."
            );
        } finally {
            setUpdating(false);
        }
    }

    if (loading) {
        return (
            <DashboardLayout>
                <p>
                    Carregando ranking...
                </p>
            </DashboardLayout>
        );
    }

    return (
        <DashboardLayout>

            <div
                style={{
                    display: "flex",
                    justifyContent:
                        "space-between",
                    alignItems: "center",
                    marginBottom: 20,
                }}
            >
                <div>
                    <h1>
                        Leaderboard
                    </h1>

                    <p>
                        Total de jogadores: {total}
                    </p>
                </div>

                <button
                    onClick={handleRefresh}
                    disabled={updating}
                >
                    {updating
                        ? "Atualizando..."
                        : "Atualizar"}
                </button>
            </div>

            {error && (
                <p
                    style={{
                        color: "red",
                    }}
                >
                    {error}
                </p>
            )}

            {userRanking && (
                <section
                    style={{
                        border:
                            "1px solid #ccc",
                        borderRadius: 8,
                        padding: 20,
                        marginBottom: 30,
                    }}
                >
                    <h2>
                        Minha Posição
                    </h2>

                    <p>
                        #{userRanking.rank}
                    </p>

                    <strong>
                        {userRanking.name}
                    </strong>

                    <p>
                        Classe:{" "}
                        {userRanking.class}
                    </p>

                    <p>
                        Nível:{" "}
                        {userRanking.level}
                    </p>

                    <p>
                        Pontos:{" "}
                        {userRanking.score}
                    </p>
                </section>
            )}

            <section>
                <h2>
                    Top Players
                </h2>

                {players.length === 0 && (
                    <p>
                        Nenhum jogador no ranking.
                    </p>
                )}

                {players.map((player) => (
                    <div
                        key={
                            player.character_id
                        }
                        style={{
                            border:
                                "1px solid #ddd",
                            borderRadius: 6,
                            padding: 15,
                            marginBottom: 10,
                            display: "flex",
                            justifyContent:
                                "space-between",
                            alignItems:
                                "center",
                        }}
                    >
                        <div
                            style={{
                                display: "flex",
                                gap: 15,
                                alignItems:
                                    "center",
                            }}
                        >
                            <strong>
                                #{player.rank}
                            </strong>

                            <div>
                                <strong>
                                    {player.name}
                                </strong>

                                <p>
                                    {
                                        player.class
                                    }
                                    {" - "}
                                    Nível{" "}
                                    {
                                        player.level
                                    }
                                </p>
                            </div>
                        </div>

                        <strong>
                            {player.score} pts
                        </strong>
                    </div>
                ))}
            </section>

        </DashboardLayout>
    );
}