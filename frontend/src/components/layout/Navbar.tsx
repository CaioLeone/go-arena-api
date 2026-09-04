import { useAuth } from "../../hooks/useAuth";

export default function Navbar() {
    const {user, logout} = useAuth();

    return (
                <header
            className="
                sticky
                top-0
                z-50
                flex
                h-16
                items-center
                justify-between
                border-b
                border-slate-800
                bg-slate-900
                px-4
                sm:px-6
            "
        >
            <div>
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
            </div>

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
        </header>
    );
}