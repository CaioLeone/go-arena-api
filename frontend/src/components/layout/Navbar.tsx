import { useAuth } from "../../hooks/useAuth";

export default function Navbar() {
    const {user, logout} = useAuth();

    return (
        <header 
            style={{
                height: "60px",
                background: "#20232a",
                color: "#fff",
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
                padding: "0 20px",
            }}
        >
            <h2>Arena dos Bárbaros</h2>
            <div
                style={{
                    display: "flex",
                    alignItems: "center",
                    gap: "15px",
                }}
            >
                <span>{user?.name}</span>
                <button onClick={logout}>Logout</button>
            </div>
        </header>
    );
}