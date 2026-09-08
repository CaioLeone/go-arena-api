import { Link } from "react-router-dom";

export default function Home() {
    return (
        <div className="min-h-screen bg-slate-950 text-white">
            <header
                className="
                    flex
                    items-center
                    justify-between
                    border-b
                    border-slate-800
                    px-6
                    py-4
                "
            >
                <h1
                    className="
                        text-xl
                        font-bold
                        text-amber-400
                    "
                >
                    Arena dos Bárbaros
                </h1>

                <div
                    className="
                        flex
                        items-center
                        gap-3
                    "
                >
                    <Link
                        to="/login"
                        className="
                            rounded-lg
                            border
                            border-slate-700
                            px-4
                            py-2
                            font-medium
                            text-slate-200
                            transition
                            hover:border-amber-500
                            hover:text-amber-400
                        "
                    >
                        Entrar
                    </Link>

                    <Link
                        to="/register"
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
                        Criar Conta
                    </Link>
                </div>
            </header>

            <main
                className="
                    mx-auto
                    flex
                    min-h-[calc(100vh-73px)]
                    max-w-7xl
                    items-center
                    px-6
                    py-12
                "
            >
                <div
                    className="
                        grid
                        w-full
                        grid-cols-1
                        items-center
                        gap-12
                        lg:grid-cols-2
                    "
                >
                    <section>
                        <span
                            className="
                                inline-block
                                rounded-full
                                bg-amber-500/10
                                px-3
                                py-1
                                text-sm
                                font-medium
                                text-amber-400
                            "
                        >
                            Entre na Arena
                        </span>

                        <h2
                            className="
                                mt-5
                                text-4xl
                                font-bold
                                leading-tight
                                sm:text-5xl
                                lg:text-6xl
                            "
                        >
                            Forje seu guerreiro.
                            <span className="text-amber-400">
                                {" "}
                                Domine a Arena.
                            </span>
                        </h2>

                        <p
                            className="
                                mt-6
                                max-w-xl
                                text-lg
                                leading-relaxed
                                text-slate-400
                            "
                        >
                            Crie personagens, evolua seus atributos,
                            enfrente outros guerreiros e conquiste seu
                            lugar no ranking da Arena dos Bárbaros.
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
                                className="
                                    rounded-lg
                                    bg-amber-500
                                    px-6
                                    py-3
                                    text-center
                                    font-bold
                                    text-slate-950
                                    transition
                                    hover:bg-amber-400
                                "
                            >
                                Criar Meu Guerreiro
                            </Link>

                            <Link
                                to="/login"
                                className="
                                    rounded-lg
                                    border
                                    border-slate-700
                                    px-6
                                    py-3
                                    text-center
                                    font-semibold
                                    text-white
                                    transition
                                    hover:border-amber-500
                                    hover:text-amber-400
                                "
                            >
                                Já tenho uma conta
                            </Link>
                        </div>
                    </section>

                    <section
                        className="
                            rounded-2xl
                            border
                            border-slate-800
                            bg-slate-900
                            p-6
                            shadow-2xl
                            sm:p-8
                        "
                    >
                        <h3
                            className="
                                text-2xl
                                font-bold
                                text-amber-400
                            "
                        >
                            Prepare-se para a batalha
                        </h3>

                        <div className="mt-6 space-y-4">
                            <div
                                className="
                                    rounded-xl
                                    bg-slate-800
                                    p-4
                                "
                            >
                                <h4 className="font-bold">
                                    Crie seu personagem
                                </h4>

                                <p
                                    className="
                                        mt-1
                                        text-sm
                                        text-slate-400
                                    "
                                >
                                    Escolha uma classe e comece sua jornada.
                                </p>
                            </div>

                            <div
                                className="
                                    rounded-xl
                                    bg-slate-800
                                    p-4
                                "
                            >
                                <h4 className="font-bold">
                                    Evolua seus atributos
                                </h4>

                                <p
                                    className="
                                        mt-1
                                        text-sm
                                        text-slate-400
                                    "
                                >
                                    Aumente vida, ataque, defesa e chance crítica.
                                </p>
                            </div>

                            <div
                                className="
                                    rounded-xl
                                    bg-slate-800
                                    p-4
                                "
                            >
                                <h4 className="font-bold">
                                    Conquiste o ranking
                                </h4>

                                <p
                                    className="
                                        mt-1
                                        text-sm
                                        text-slate-400
                                    "
                                >
                                    Vença batalhas e suba entre os melhores guerreiros.
                                </p>
                            </div>
                        </div>
                    </section>
                </div>
            </main>
        </div>
    );
}