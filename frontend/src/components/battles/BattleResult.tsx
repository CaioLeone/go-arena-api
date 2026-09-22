import type { Battle } from "../../types/battle";

interface Props {
    battle: Battle;
}

export default function BattleResult({ battle }: Props) {
    return (
        <section className="arena-card">
            {/* Cabeçalho */}
            <div
                className="
                    border-b
                    border-slate-800
                    pb-5
                "
            >
                <span className="arena-badge">
                    Resultado
                </span>

                <h2 className="arena-section-title mt-3">
                    Resultado da Batalha
                </h2>

                <div
                    className="
                        mt-5
                        flex
                        flex-col
                        items-center
                        justify-center
                        gap-3
                        sm:flex-row
                    "
                >
                    <span className="text-xl font-bold text-white">
                        {battle.attacker_name}
                    </span>

                    <span
                        className="
                            text-lg
                            font-black
                            text-amber-400
                        "
                    >
                        VS
                    </span>

                    <span className="text-xl font-bold text-white">
                        {battle.defender_name}
                    </span>
                </div>
            </div>

            {/* Vencedor */}
            <div
                className="
                    mt-6
                    rounded-xl
                    border
                    border-amber-500/20
                    bg-amber-500/10
                    p-5
                    text-center
                "
            >
                <p
                    className="
                        text-xs
                        font-semibold
                        uppercase
                        tracking-widest
                        text-slate-400
                    "
                >
                    Vencedor
                </p>

                <h3
                    className="
                        mt-1
                        text-2xl
                        font-black
                        text-amber-400
                    "
                >
                    ⚔ {battle.winner_name}
                </h3>
            </div>

            {/* Estatísticas */}
            <div
                className="
                    mt-6
                    grid
                    grid-cols-2
                    gap-3
                    md:grid-cols-4
                "
            >
                <BattleStat
                    label="Dano Total"
                    value={battle.damage_dealt}
                />

                <BattleStat
                    label={`HP ${battle.attacker_name}`}
                    value={battle.attacker_hp_final}
                />

                <BattleStat
                    label={`HP ${battle.defender_name}`}
                    value={battle.defender_hp_final}
                />

                <BattleStat
                    label="Rounds"
                    value={battle.rounds_count}
                />
            </div>

            {/* Rounds */}
            <div className="mt-8">
                <div
                    className="
                        mb-5
                        border-b
                        border-slate-800
                        pb-3
                    "
                >
                    <h3 className="arena-section-title">
                        Combate
                    </h3>

                    <p className="arena-subtitle">
                        Acompanhe cada golpe da batalha.
                    </p>
                </div>

                <div className="space-y-3">
                    {battle.rounds.map(
                        (round, index) => (
                            <div
                                key={`${round.round}-${index}`}
                                className="
                                    rounded-xl
                                    border
                                    border-slate-800
                                    bg-slate-950/40
                                    p-4
                                "
                            >
                                <div
                                    className="
                                        flex
                                        flex-col
                                        gap-3
                                        sm:flex-row
                                        sm:items-center
                                        sm:justify-between
                                    "
                                >
                                    <div>
                                        <div
                                            className="
                                                mb-2
                                                flex
                                                items-center
                                                gap-2
                                            "
                                        >
                                            <span
                                                className="
                                                    rounded-md
                                                    bg-slate-800
                                                    px-2
                                                    py-1
                                                    text-xs
                                                    font-bold
                                                    text-slate-300
                                                "
                                            >
                                                Round {round.round}
                                            </span>

                                            {round.is_critical && (
                                                <span
                                                    className="
                                                        rounded-md
                                                        bg-red-500/10
                                                        px-2
                                                        py-1
                                                        text-xs
                                                        font-bold
                                                        text-red-400
                                                    "
                                                >
                                                    CRÍTICO!
                                                </span>
                                            )}
                                        </div>

                                        <p className="font-semibold text-white">
                                            {round.attacker_name}

                                            <span className="mx-2 text-amber-400">
                                                →
                                            </span>

                                            {round.defender_name}
                                        </p>

                                        <p className="mt-1 text-sm text-slate-400">
                                            {round.message}
                                        </p>
                                    </div>

                                    <div
                                        className="
                                            flex
                                            gap-3
                                            sm:text-right
                                        "
                                    >
                                        <div>
                                            <span className="text-xs text-slate-500">
                                                Dano
                                            </span>

                                            <p className="font-bold text-red-400">
                                                -{round.damage}
                                            </p>
                                        </div>

                                        <div>
                                            <span className="text-xs text-slate-500">
                                                HP restante
                                            </span>

                                            <p className="font-bold text-emerald-400">
                                                {round.remaining_hp}
                                            </p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        )
                    )}
                </div>
            </div>
        </section>
    );
}

interface BattleStatProps {
    label: string;
    value: string | number;
}

function BattleStat({
    label,
    value,
}: BattleStatProps) {
    return (
        <div
            className="
                rounded-xl
                border
                border-slate-800
                bg-slate-950/40
                p-4
            "
        >
            <span
                className="
                    text-xs
                    uppercase
                    tracking-wide
                    text-slate-500
                "
            >
                {label}
            </span>

            <p
                className="
                    mt-1
                    text-xl
                    font-bold
                    text-white
                "
            >
                {value}
            </p>
        </div>
    );
}