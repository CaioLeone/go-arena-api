import { useAuth } from "../../hooks/useAuth";
import { NavLink } from "react-router-dom";

export default function Navbar() {
    const {user, logout} = useAuth();

    return (
        <header
            className="
                sticky
                top-0
                z-50
                border-b
                border-slate-800
                bg-slate-900
            "
        >

            <div
                className="
                    flex
                    h-16
                    items-center
                    justify-between
                    px-4
                    sm:px-6
                "
            >

                <h1
                    className="
                        text-lg
                        font-bold
                        text-amber-400
                        sm:text-xl
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

                    <span
                        className="
                            hidden
                            text-sm
                            text-slate-300
                            sm:block
                        "
                    >
                        {user?.name}
                    </span>

                    <button
                        onClick={logout}
                        className="
                            rounded-lg
                            bg-red-600
                            px-3
                            py-2
                            text-sm
                            font-medium
                            text-white
                            transition
                            hover:bg-red-500
                        "
                    >
                        Sair
                    </button>
                </div>
            </div>

            <nav
                className="
                    flex
                    overflow-x-auto
                    border-t
                    border-slate-800
                    px-2
                    md:hidden
                "
            >

                <NavLink
                    to="/dashboard"
                    className="whitespace-nowrap px-3 py-3 text-sm text-slate-300"
                >
                    Início
                </NavLink>

                <NavLink
                    to="/characters"
                    className="whitespace-nowrap px-3 py-3 text-sm text-slate-300"
                >
                    Personagens
                </NavLink>

                <NavLink
                    to="/battles"
                    className="whitespace-nowrap px-3 py-3 text-sm text-slate-300"
                >
                    Batalhas
                </NavLink>

                <NavLink
                    to="/leaderboard"
                    className="whitespace-nowrap px-3 py-3 text-sm text-slate-300"
                >
                    Ranking
                </NavLink>
            </nav>
        </header>
    );
}