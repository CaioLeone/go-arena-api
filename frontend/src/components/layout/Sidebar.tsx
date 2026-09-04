import { NavLink } from "react-router-dom";

export default function Sidebar(){
    function linkClass({
        isActive,
    }: {
        isActive: boolean;
    }) {
        return `
            rounded-lg
            px-3
            py-2
            text-sm
            font-medium
            transition

            ${
                isActive
                    ? "bg-amber-500 text-slate-950"
                    : "text-slate-300 hover:bg-slate-800 hover:text-white"
            }
        `;
    }

    return (
        <aside
            className="
                hidden
                min-h-[calc(100vh-4rem)]
                w-56
                shrink-0
                border-r
                border-slate-800
                bg-slate-900
                p-4
                md:block
            "
        >
            <nav
                className="
                    flex
                    flex-col
                    gap-2
                "
            >

                <NavLink
                    to="/dashboard"
                    className={linkClass}
                >
                    Dashboard
                </NavLink>

                <NavLink
                    to="/characters"
                    className={linkClass}
                >
                    Personagens
                </NavLink>

                <NavLink
                    to="/battles"
                    className={linkClass}
                >
                    Batalhas
                </NavLink>

                <NavLink
                    to="/leaderboard"
                    className={linkClass}
                >
                    Ranking
                </NavLink>
            </nav>
        </aside>
    );
}