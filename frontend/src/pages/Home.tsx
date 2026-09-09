import { Link } from "react-router-dom";

export default function Home() {
    return (
        <div className="arena-page">

            <header
                className="
                    border-b
                    border-slate-800/80
                    bg-slate-950/80
                    backdrop-blur
                "
            >
                <div
                    className="
                        arena-container
                        flex
                        h-20
                        items-center
                        justify-between
                    "
                >

                    <Link
                        to="/"
                        className="
                            text-xl
                            font-black
                            tracking-tight
                            text-white
                            sm:text-2xl
                        "
                    >
                        Arena dos
                        <span className="text-amber-400">
                            {" "}Bárbaros
                        </span>
                    </Link>

                    <div className="flex items-center gap-3">

                        <Link
                            to="/login"
                            className="arena-button-secondary py-2"
                        >
                            Entrar
                        </Link>

                        <Link
                            to="/register"
                            className="
                                arena-button-primary
                                hidden
                                py-2
                                sm:inline-flex
                            "
                        >
                            Criar Conta
                        </Link>

                    </div>

                </div>
            </header>

            <main>

                <section
                    className="
                        arena-container
                        grid
                        min-h-[calc(100vh-80px)]
                        items-center
                        gap-12
                        py-16
                        lg:grid-cols-[1.15fr_0.85fr]
                        lg:py-20
                    "
                >

                    <div>

                        <span className="arena-badge">
                            ⚔ Entre na Arena
                        </span>

                        <h1
                            className="
                                mt-6
                                max-w-3xl
                                text-5xl
                                font-black
                                leading-[1.05]
                                tracking-tight
                                text-white
                                sm:text-6xl
                                lg:text-7xl
                            "
                        >
                            Forje seu guerreiro.

                            <span
                                className="
                                    block
                                    text-amber-400
                                "
                            >
                                Domine a Arena.
                            </span>
                        </h1>

                        <p
                            className="
                                mt-6
                                max-w-2xl
                                text-lg
                                leading-8
                                text-slate-400
                                sm:text-xl
                            "
                        >
                            Crie personagens, desenvolva seus
                            atributos, enfrente adversários e lute
                            pelo topo do ranking da Arena dos
                            Bárbaros.
                        </p>

                        <div
                            className="
                                mt-8
                                flex
                                flex-col
                                gap-3
                                sm:flex-row
                            "
                        >

                            <Link
                                to="/register"
                                className="arena-button-primary"
                            >
                                Criar meu guerreiro
                            </Link>

                            <Link
                                to="/login"
                                className="arena-button-secondary"
                            >
                                Já tenho uma conta
                            </Link>

                        </div>

                        <div
                            className="
                                mt-12
                                grid
                                max-w-2xl
                                grid-cols-3
                                gap-4
                                border-t
                                border-slate-800
                                pt-6
                            "
                        >

                            <div>
                                <strong className="text-lg text-white">
                                    4
                                </strong>

                                <p className="text-sm text-slate-500">
                                    Classes
                                </p>
                            </div>

                            <div>
                                <strong className="text-lg text-white">
                                    PvP
                                </strong>

                                <p className="text-sm text-slate-500">
                                    Batalhas
                                </p>
                            </div>

                            <div>
                                <strong className="text-lg text-white">
                                    Ranking
                                </strong>

                                <p className="text-sm text-slate-500">
                                    Competitivo
                                </p>
                            </div>

                        </div>

                    </div>

                    <div
                        className="
                            relative
                            mx-auto
                            w-full
                            max-w-lg
                        "
                    >

                        <div
                            className="
                                absolute
                                -inset-10
                                -z-10
                                rounded-full
                                bg-amber-500/10
                                blur-3xl
                            "
                        />

                        <div className="arena-card p-6 sm:p-8">

                            <div
                                className="
                                    mb-8
                                    flex
                                    items-center
                                    justify-between
                                "
                            >

                                <div>
                                    <p
                                        className="
                                            text-xs
                                            font-bold
                                            uppercase
                                            tracking-[0.2em]
                                            text-amber-400
                                        "
                                    >
                                        Arena
                                    </p>

                                    <h2
                                        className="
                                            mt-1
                                            text-2xl
                                            font-black
                                        "
                                    >
                                        Prepare-se para a batalha
                                    </h2>
                                </div>

                                <div
                                    className="
                                        flex
                                        h-12
                                        w-12
                                        items-center
                                        justify-center
                                        rounded-xl
                                        bg-amber-500
                                        text-2xl
                                    "
                                >
                                    ⚔
                                </div>

                            </div>

                            <div className="space-y-3">

                                <Feature
                                    number="01"
                                    title="Crie seu personagem"
                                    description="Escolha entre Bárbaro, Mago, Arqueiro ou Assassino."
                                />

                                <Feature
                                    number="02"
                                    title="Evolua seus atributos"
                                    description="Melhore vida, ataque, defesa e chance crítica."
                                />

                                <Feature
                                    number="03"
                                    title="Entre em batalha"
                                    description="Enfrente adversários e descubra quem domina a Arena."
                                />

                                <Feature
                                    number="04"
                                    title="Suba no ranking"
                                    description="Acumule pontos e conquiste sua posição entre os melhores."
                                />

                            </div>

                        </div>

                    </div>

                </section>

            </main>

        </div>
    );
}

interface FeatureProps {
    number: string;
    title: string;
    description: string;
}

function Feature({
    number,
    title,
    description,
}: FeatureProps) {
    return (
        <div
            className="
                flex
                gap-4
                rounded-xl
                border
                border-slate-800
                bg-slate-950/50
                p-4
            "
        >

            <span
                className="
                    font-black
                    text-amber-500
                "
            >
                {number}
            </span>

            <div>

                <h3 className="font-bold text-white">
                    {title}
                </h3>

                <p
                    className="
                        mt-1
                        text-sm
                        leading-6
                        text-slate-400
                    "
                >
                    {description}
                </p>

            </div>

        </div>
    );
}