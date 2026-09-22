import { useEffect, useState } from "react";
import {
    useNavigate,
    useParams,
} from "react-router-dom";

import DashboardLayout
    from "../components/layout/DashboardLayout";

import characterService
    from "../services/characterService";

import type { Character }
    from "../types/character";

type Attribute =
    | "hp"
    | "attack"
    | "defense"
    | "critical_chance";

export default function CharacterDetail() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [character, setCharacter] =
        useState<Character | null>(null);

    const [loading, setLoading] =
        useState(true);

    const [actionLoading, setActionLoading] =
        useState(false);

    const [name, setName] =
        useState("");

    const [error, setError] =
        useState("");

    const [success, setSuccess] =
        useState("");

    useEffect(() => {
        if (id) {
            loadCharacter();
        }
    }, [id]);

    async function loadCharacter() {
        if (!id) {
            return;
        }

        try {
            setLoading(true);
            setError("");

            const data =
                await characterService.getById(id);

            setCharacter(data);
            setName(data.name);

        } catch (err) {
            console.error(
                "Erro ao carregar personagem:",
                err
            );

            setError(
                "Não foi possível carregar o personagem."
            );

        } finally {
            setLoading(false);
        }
    }

    async function refreshCharacter() {
        if (!id) {
            return;
        }

        const data =
            await characterService.getById(id);

        setCharacter(data);
        setName(data.name);
    }

    async function handleUpdate() {
        if (!character) {
            return;
        }

        if (!name.trim()) {
            setError(
                "O nome do personagem não pode ficar vazio."
            );
            return;
        }

        try {
            setActionLoading(true);
            setError("");
            setSuccess("");

            const updated =
                await characterService.update(
                    character.id,
                    {
                        name: name.trim(),
                    }
                );

            setCharacter(updated);

            setSuccess(
                "Personagem atualizado com sucesso."
            );

        } catch (err) {
            console.error(
                "Erro ao atualizar personagem:",
                err
            );

            setError(
                "Não foi possível atualizar o personagem."
            );

        } finally {
            setActionLoading(false);
        }
    }

    async function handleDelete() {
        if (!character) {
            return;
        }

        const confirmed = confirm(
            `Deseja realmente excluir ${character.name}?`
        );

        if (!confirmed) {
            return;
        }

        try {
            setActionLoading(true);
            setError("");

            await characterService.remove(
                character.id
            );

            navigate("/characters");

        } catch (err) {
            console.error(
                "Erro ao excluir personagem:",
                err
            );

            setError(
                "Não foi possível excluir o personagem."
            );

            setActionLoading(false);
        }
    }

    async function handleAddXp() {
        if (!character) {
            return;
        }

        try {
            setActionLoading(true);
            setError("");
            setSuccess("");

            await characterService.addExperience(
                character.id,
                {
                    experience: 10,
                }
            );

            await refreshCharacter();

            setSuccess(
                "10 XP adicionados com sucesso."
            );

        } catch (err) {
            console.error(
                "Erro ao adicionar XP:",
                err
            );

            setError(
                "Não foi possível adicionar experiência."
            );

        } finally {
            setActionLoading(false);
        }
    }

    async function handleSpendPoints(
        attribute: Attribute
    ) {
        if (!character) {
            return;
        }

        if (character.attribute_points <= 0) {
            setError(
                "Este personagem não possui pontos de atributo disponíveis."
            );
            return;
        }

        const attributes = {
            hp: 0,
            attack: 0,
            defense: 0,
            critical_chance: 0,
        };

        attributes[attribute] = 1;

        try {
            setActionLoading(true);
            setError("");
            setSuccess("");

            await characterService.spendAttribute(
                character.id,
                attributes
            );

            await refreshCharacter();

            setSuccess(
                "Ponto de atributo distribuído com sucesso."
            );

        } catch (err) {
            console.error(
                "Erro ao distribuir atributo:",
                err
            );

            setError(
                "Não foi possível distribuir o ponto de atributo."
            );

        } finally {
            setActionLoading(false);
        }
    }

    if (loading) {
        return (
            <DashboardLayout>
                <div className="arena-content">
                    <p className="text-slate-400">
                        Carregando personagem...
                    </p>
                </div>
            </DashboardLayout>
        );
    }

    if (!character) {
        return (
            <DashboardLayout>
                <div className="arena-content">
                    <div className="arena-error">
                        {error ||
                            "Personagem não encontrado."}
                    </div>
                </div>
            </DashboardLayout>
        );
    }

    return (
        <DashboardLayout>
            <div className="arena-content">

                {/* Cabeçalho */}
                <div className="mb-8">
                    <span className="arena-badge">
                        Guerreiro
                    </span>

                    <h1 className="arena-title mt-3">
                        {character.name}
                    </h1>

                    <p className="arena-subtitle">
                        {character.class}
                        {" • "}
                        Nível {character.level}
                    </p>
                </div>

                {/* Mensagens */}
                {error && (
                    <div className="arena-error mb-6">
                        {error}
                    </div>
                )}

                {success && (
                    <div
                        className="
                            mb-6
                            rounded-lg
                            border
                            border-emerald-500/30
                            bg-emerald-500/10
                            px-4
                            py-3
                            text-sm
                            text-emerald-400
                        "
                    >
                        {success}
                    </div>
                )}

                {/* Informações gerais */}
                <section className="arena-card mb-6">
                    <div
                        className="
                            flex
                            flex-col
                            gap-6
                            md:flex-row
                            md:items-center
                            md:justify-between
                        "
                    >
                        <div>
                            <h2 className="arena-section-title">
                                Progressão
                            </h2>

                            <p className="arena-subtitle">
                                Evolução atual do guerreiro.
                            </p>
                        </div>

                        <div
                            className="
                                grid
                                grid-cols-3
                                gap-3
                            "
                        >
                            <InfoBox
                                label="Nível"
                                value={character.level}
                            />

                            <InfoBox
                                label="XP"
                                value={character.experience}
                            />

                            <InfoBox
                                label="Pontos"
                                value={
                                    character.attribute_points
                                }
                                highlight
                            />
                        </div>
                    </div>
                </section>

                {/* Atributos */}
                <section className="arena-card mb-6">
                    <div className="mb-6">
                        <h2 className="arena-section-title">
                            Atributos
                        </h2>

                        <p className="arena-subtitle">
                            Distribua seus pontos para
                            fortalecer o guerreiro.
                        </p>
                    </div>

                    <div
                        className="
                            grid
                            grid-cols-1
                            gap-4
                            sm:grid-cols-2
                        "
                    >
                        <AttributeCard
                            label="Vida"
                            value={character.hp}
                            onAdd={() =>
                                handleSpendPoints("hp")
                            }
                            disabled={
                                actionLoading ||
                                character.attribute_points <= 0
                            }
                        />

                        <AttributeCard
                            label="Ataque"
                            value={character.attack}
                            onAdd={() =>
                                handleSpendPoints(
                                    "attack"
                                )
                            }
                            disabled={
                                actionLoading ||
                                character.attribute_points <= 0
                            }
                        />

                        <AttributeCard
                            label="Defesa"
                            value={character.defense}
                            onAdd={() =>
                                handleSpendPoints(
                                    "defense"
                                )
                            }
                            disabled={
                                actionLoading ||
                                character.attribute_points <= 0
                            }
                        />

                        <AttributeCard
                            label="Crítico"
                            value={
                                character.critical_chance
                            }
                            onAdd={() =>
                                handleSpendPoints(
                                    "critical_chance"
                                )
                            }
                            disabled={
                                actionLoading ||
                                character.attribute_points <= 0
                            }
                        />
                    </div>

                    {character.attribute_points <= 0 && (
                        <p
                            className="
                                mt-5
                                text-sm
                                text-slate-500
                            "
                        >
                            Suba de nível para obter
                            novos pontos de atributo.
                        </p>
                    )}
                </section>

                {/* Nome */}
                <section className="arena-card mb-6">
                    <h2 className="arena-section-title">
                        Editar Personagem
                    </h2>

                    <div className="mt-5">
                        <label className="arena-label">
                            Nome
                        </label>

                        <input
                            className="arena-input"
                            value={name}
                            onChange={(e) =>
                                setName(e.target.value)
                            }
                            disabled={actionLoading}
                        />
                    </div>

                    <button
                        onClick={handleUpdate}
                        className="
                            arena-button-primary
                            mt-5
                        "
                        disabled={actionLoading}
                    >
                        Salvar alterações
                    </button>
                </section>

                {/* Ações */}
                <section className="arena-card">
                    <h2 className="arena-section-title">
                        Ações
                    </h2>

                    <div
                        className="
                            mt-5
                            flex
                            flex-wrap
                            gap-3
                        "
                    >
                        <button
                            onClick={handleAddXp}
                            className="arena-button-primary"
                            disabled={actionLoading}
                        >
                            +10 XP
                        </button>

                        <button
                            onClick={handleDelete}
                            disabled={actionLoading}
                            className="
                                rounded-lg
                                bg-red-600
                                px-4
                                py-2
                                font-semibold
                                text-white
                                transition
                                hover:bg-red-500
                                disabled:cursor-not-allowed
                                disabled:opacity-50
                            "
                        >
                            Excluir personagem
                        </button>
                    </div>
                </section>

            </div>
        </DashboardLayout>
    );
}

/* -----------------------------
   Componentes auxiliares
------------------------------ */

interface InfoBoxProps {
    label: string;
    value: string | number;
    highlight?: boolean;
}

function InfoBox({
    label,
    value,
    highlight = false,
}: InfoBoxProps) {
    return (
        <div
            className="
                min-w-20
                rounded-xl
                border
                border-slate-800
                bg-slate-950/40
                p-3
                text-center
            "
        >
            <span
                className="
                    block
                    text-xs
                    uppercase
                    text-slate-500
                "
            >
                {label}
            </span>

            <strong
                className={
                    highlight
                        ? "text-xl text-amber-400"
                        : "text-xl text-white"
                }
            >
                {value}
            </strong>
        </div>
    );
}

interface AttributeCardProps {
    label: string;
    value: number;
    onAdd: () => void;
    disabled: boolean;
}

function AttributeCard({
    label,
    value,
    onAdd,
    disabled,
}: AttributeCardProps) {
    return (
        <div
            className="
                flex
                items-center
                justify-between
                rounded-xl
                border
                border-slate-800
                bg-slate-950/40
                p-4
            "
        >
            <div>
                <span className="text-sm text-slate-400">
                    {label}
                </span>

                <p
                    className="
                        mt-1
                        text-2xl
                        font-black
                        text-white
                    "
                >
                    {value}
                </p>
            </div>

            <button
                onClick={onAdd}
                disabled={disabled}
                className="
                    flex
                    h-9
                    w-9
                    items-center
                    justify-center
                    rounded-lg
                    bg-emerald-600
                    text-xl
                    font-bold
                    text-white
                    transition
                    hover:bg-emerald-500
                    disabled:cursor-not-allowed
                    disabled:bg-slate-700
                    disabled:text-slate-500
                "
            >
                +
            </button>
        </div>
    );
}