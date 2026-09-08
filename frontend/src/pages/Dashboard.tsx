import DashboardLayout from "../components/layout/DashboardLayout";
import { useAuth } from "../hooks/useAuth";

export default function Dashboard() {
    const { user } = useAuth();

    return (
        <DashboardLayout>

            <div className="space-y-8">

                <section>

                    <h1
                        className="
                            text-3xl
                            font-bold
                            text-white
                        "
                    >
                        Bem-vindo à Arena
                    </h1>

                    <p
                        className="
                            mt-2
                            text-slate-400
                        "
                    >
                        Guerreiro {user?.name}, prepare seus personagens para a batalha.
                    </p>

                </section>

                <div
                    className="
                        grid
                        grid-cols-1
                        gap-4
                        md:grid-cols-3
                    "
                >

                    <div
                        className="
                            rounded-xl
                            border
                            border-slate-700
                            bg-slate-900
                            p-6
                        "
                    >
                        <h2 className="font-bold text-amber-400">
                            Personagens
                        </h2>

                        <p className="mt-2 text-sm text-slate-400">
                            Crie e evolua seus guerreiros.
                        </p>
                    </div>

                    <div
                        className="
                            rounded-xl
                            border
                            border-slate-700
                            bg-slate-900
                            p-6
                        "
                    >
                        <h2 className="font-bold text-red-400">
                            Batalhas
                        </h2>

                        <p className="mt-2 text-sm text-slate-400">
                            Coloque seus personagens à prova.
                        </p>
                    </div>

                    <div
                        className="
                            rounded-xl
                            border
                            border-slate-700
                            bg-slate-900
                            p-6
                        "
                    >
                        <h2 className="font-bold text-emerald-400">
                            Ranking
                        </h2>

                        <p className="mt-2 text-sm text-slate-400">
                            Dispute posições na Arena.
                        </p>
                    </div>
                </div>
            </div>
        </DashboardLayout>
    );
}