import type { Battle} from "../../types/battle";

interface Props {
    battle: Battle;
}

export default function BattleResult({ battle }: Props) {
    return (
        <div
            style={{
                border: "1px solid #ccc",
                borderRadius: 8,
                padding: 20,
                marginTop: 20,
            }}
        >
            <h2>Resultado da Batalha</h2>

            <p>
                {battle.attacker_name}
                {" vs "}
                {battle.defender_name}
            </p>

            <h3>
                Vencedor: {battle.winner_name}
            </h3>

            <p>
                Dano total: {battle.damage_dealt}
            </p>

            <p>
                HP final de {battle.attacker_name}:{" "}
                {battle.attacker_hp_final}
            </p>

            <p>
                HP final de {battle.defender_name}:{" "}
                {battle.defender_hp_final}
            </p>

            <p>
                Total de rounds: {battle.rounds_count}
            </p>

            <hr />

            <h3>Rounds</h3>

            {battle.rounds.map((round) => (
                <div
                    key={round.round}
                    style={{
                        marginBottom: 15,
                    }}
                >
                    <strong>
                        Round {round.round}
                    </strong>

                    <p>
                        {round.message}
                    </p>

                    <p>
                        {round.attacker_name}
                        {" -> "}
                        {round.defender_name}
                    </p>

                    <p>
                        Dano: {round.damage}
                    </p>

                    {round.is_critical && (
                        <p>
                            Golpe crítico!
                        </p>
                    )}

                    <p>
                        HP restante:{" "}
                        {round.remaining_hp}
                    </p>
                </div>
            ))}
        </div>
    );
}