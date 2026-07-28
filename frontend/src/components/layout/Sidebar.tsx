import { NavLink } from "react-router-dom";

export default function Sidebar(){
    return(
        <aside
            style={{
                width: "220px",
                background: "#f4f4f4",
                padding: "20px",
                minHeight: "calc(100vh - 60px)",
            }}
        >
            <nav
                style={{
                    display: "flex",
                    flexDirection: "column",
                    gap: "15px",
                }}
            >
                <NavLink to={"/"}>Dashboard</NavLink>
                <NavLink to={"/characters"}>Guerreiros</NavLink>
                <NavLink to={"/battles"}>Batalhas</NavLink>
                <NavLink to={"/leaderboard"}>Ranking</NavLink>
            </nav>
        </aside>
    );
}