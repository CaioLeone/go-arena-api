import { useEffect, useState } from "react";

import DashboardLayout from "../components/layout/DashboardLayout";
import rankingService from "../services/rankingService";

import type {
    TopPlayer,
    UserRanking,
} from "../types/ranking";

export default function LeaderBoard() {
    const [players, setPlayers] =
        useState<TopPlayer[]>([]);

    const [userRanking, setUserRanking] =
        useState<UserRanking | null>(null);

    const [total, setTotal] =
        useState(0);

    const [loading, setLoading] =
        useState(true);

    const [updating, setUpdating] =
        useState(false);

    const [error, setError] =
        useState("");

    useEffect(() => {
        async function load() {
            try {
                setError("");

                await loadRanking();
            } catch (err) {
                console.error(
                    "Erro ao carregar ranking:",
                    err
                );

                setError(
                    "Erro ao carregar o ranking."
                );
            } finally {
                setLoading(false);
            }
        }

        load();
    }, []);

    async function loadRanking() {
        const leaderboard =
            await rankingService.getTopPlayers();

        console.log(
            "Leaderboard:",
            leaderboard
        );

        const validPlayers =
            (leaderboard.players ?? [])
                .filter(
                    (player): player is TopPlayer =>
                        player != null
                );

        setPlayers(validPlayers);

        setTotal(
            leaderboard.total ??
            validPlayers.length
        );

        try {
            const myRanking =
                await rankingService.getUserRanking();

            console.log(
                "Meu ranking:",
                myRanking
            );

            setUserRanking(
                myRanking ?? null
            );

        } catch (err) {
            console.error(
                "Erro ao carregar posição do usuário:",
                err
            );

            setUserRanking(null);
        }
    }

    async function handleRefresh() {
        try {
            setUpdating(true);
            setError("");

            await loadRanking();

        } catch (err) {
            console.error(
                "Erro ao atualizar ranking:",
                err
            );

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
                <div className="arena-content">
                    <p className="text-slate-400">
                        Carregando ranking...
                    </p>
                </div>
            </DashboardLayout>
        );
    }

    return (
        <DashboardLayout>
            <div className="arena-content">

                {/* Cabeçalho */}
                <div
                    className="
                        mb-8
                        flex
                        flex-col
                        gap-4
                        sm:flex-row
                        sm:items-end
                        sm:justify-between
                    "
                >
                    <div>
                        <span className="arena-badge">
                            Salão da Glória
                        </span>

                        <h1 className="arena-title mt-3">
                            Ranking
                        </h1>

                        <p className="arena-subtitle">
                            Os guerreiros mais poderosos
                            da Arena dos Bárbaros.
                        </p>

                        <p
                            className="
                                mt-2
                                text-sm
                                text-slate-500
                            "
                        >
                            {total}{" "}
                            {total === 1
                                ? "guerreiro no ranking"
                                : "guerreiros no ranking"}
                        </p>
                    </div>

                    <button
                        className="
                            rounded-lg
                            border
                            border-slate-700
                            bg-slate-900
                            px-4
                            py-2
                            font-semibold
                            text-slate-300
                            transition
                            hover:border-amber-500/50
                            hover:text-amber-400
                            disabled:cursor-not-allowed
                            disabled:opacity-50
                        "
                        onClick={handleRefresh}
                        disabled={updating}
                    >
                        {updating
                            ? "Atualizando..."
                            : "↻ Atualizar"}
                    </button>
                </div>

                {/* Erro */}
                {error && (
                    <div className="arena-error mb-6">
                        {error}
                    </div>
                )}

                {/* Minha posição */}
                {userRanking && (
                    <section
                        className="
                            arena-card
                            mb-10
                            border-amber-500/30
                        "
                    >
                        <div
                            className="
                                flex
                                flex-col
                                gap-5
                                sm:flex-row
                                sm:items-center
                                sm:justify-between
                            "
                        >
                            <div>
                                <span
                                    className="
                                        text-xs
                                        font-bold
                                        uppercase
                                        tracking-wider
                                        text-amber-400
                                    "
                                >
                                    Minha posição
                                </span>

                                <div
                                    className="
                                        mt-2
                                        flex
                                        items-center
                                        gap-4
                                    "
                                >
                                    <div
                                        className="
                                            flex
                                            h-14
                                            w-14
                                            items-center
                                            justify-center
                                            rounded-xl
                                            border
                                            border-amber-500/30
                                            bg-amber-500/10
                                            text-xl
                                            font-black
                                            text-amber-400
                                        "
                                    >
                                        #{userRanking.rank}
                                    </div>

                                    <div>
                                        <h2
                                            className="
                                                text-xl
                                                font-bold
                                                text-white
                                            "
                                        >
                                            {userRanking.name ||
                                                "Guerreiro"}
                                        </h2>

                                        <p
                                            className="
                                                text-sm
                                                text-slate-400
                                            "
                                        >
                                            {userRanking.class ||
                                                "Classe não informada"}

                                            {" • "}

                                            Nível{" "}
                                            {userRanking.level ?? "-"}
                                        </p>
                                    </div>
                                </div>
                            </div>

                            <div
                                className="
                                    rounded-xl
                                    border
                                    border-amber-500/20
                                    bg-amber-500/10
                                    px-6
                                    py-3
                                    text-center
                                "
                            >
                                <span
                                    className="
                                        block
                                        text-xs
                                        uppercase
                                        tracking-wider
                                        text-slate-400
                                    "
                                >
                                    Pontos
                                </span>

                                <strong
                                    className="
                                        text-2xl
                                        font-black
                                        text-amber-400
                                    "
                                >
                                    {userRanking.score ?? 0}
                                </strong>
                            </div>
                        </div>
                    </section>
                )}

                {/* Top Players */}
                <section>
                    <div className="mb-5">
                        <h2 className="arena-section-title">
                            Top Guerreiros
                        </h2>

                        <p className="arena-subtitle">
                            Classificação geral dos
                            combatentes da arena.
                        </p>
                    </div>

                    {players.length === 0 ? (
                        <div className="arena-card">
                            <p className="text-slate-400">
                                Nenhum guerreiro no ranking.
                            </p>
                        </div>
                    ) : (
                        <div className="space-y-3">
                            {players.map(
                                (player, index) => (
                                    <div
                                        key={
                                            player.character_id ||
                                            `ranking-${player.rank}-${index}`
                                        }
                                        className="
                                            arena-card
                                            arena-card-hover
                                        "
                                    >
                                        <div
                                            className="
                                                flex
                                                items-center
                                                justify-between
                                                gap-4
                                            "
                                        >
                                            {/* Posição + personagem */}
                                            <div
                                                className="
                                                    flex
                                                    min-w-0
                                                    items-center
                                                    gap-4
                                                "
                                            >
                                                <RankingPosition
                                                    rank={
                                                        player.rank
                                                    }
                                                />

                                                <div className="min-w-0">
                                                    <h3
                                                        className="
                                                            truncate
                                                            font-bold
                                                            text-white
                                                        "
                                                    >
                                                        {player.name ||
                                                            "Guerreiro"}
                                                    </h3>

                                                    <p
                                                        className="
                                                            mt-1
                                                            text-sm
                                                            text-slate-400
                                                        "
                                                    >
                                                        {player.class ||
                                                            "Classe não informada"}

                                                        {" • "}

                                                        Nível{" "}
                                                        {player.level ??
                                                            "-"}
                                                    </p>
                                                </div>
                                            </div>

                                            {/* Pontuação */}
                                            <div
                                                className="
                                                    shrink-0
                                                    text-right
                                                "
                                            >
                                                <strong
                                                    className="
                                                        text-xl
                                                        font-black
                                                        text-amber-400
                                                    "
                                                >
                                                    {player.score ??
                                                        0}
                                                </strong>

                                                <span
                                                    className="
                                                        ml-1
                                                        text-sm
                                                        text-slate-500
                                                    "
                                                >
                                                    pts
                                                </span>
                                            </div>
                                        </div>
                                    </div>
                                )
                            )}
                        </div>
                    )}
                </section>

            </div>
        </DashboardLayout>
    );
}

/* ----------------------------------
   Posição no ranking
----------------------------------- */

interface RankingPositionProps {
    rank: number;
}

function RankingPosition({
    rank,
}: RankingPositionProps) {

    if (rank === 1) {
        return (
            <div
                className="
                    flex
                    h-12
                    w-12
                    shrink-0
                    items-center
                    justify-center
                    rounded-xl
                    bg-amber-500/15
                    text-xl
                "
            >
                🥇
            </div>
        );
    }

    if (rank === 2) {
        return (
            <div
                className="
                    flex
                    h-12
                    w-12
                    shrink-0
                    items-center
                    justify-center
                    rounded-xl
                    bg-slate-500/15
                    text-xl
                "
            >
                🥈
            </div>
        );
    }

    if (rank === 3) {
        return (
            <div
                className="
                    flex
                    h-12
                    w-12
                    shrink-0
                    items-center
                    justify-center
                    rounded-xl
                    bg-orange-500/10
                    text-xl
                "
            >
                🥉
            </div>
        );
    }

    return (
        <div
            className="
                flex
                h-12
                w-12
                shrink-0
                items-center
                justify-center
                rounded-xl
                border
                border-slate-800
                bg-slate-950/50
                font-bold
                text-slate-400
            "
        >
            #{rank}
        </div>
    );
}