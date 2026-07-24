import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import api from "../services/api";

export default function Register() {
    const navigate = useNavigate();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();
        setError("");

        if(password !== confirmPassword){
            setError("As senhas nao coincidem");
            return;
        }

        try {
            setLoading(true);
            
            await api.post("/auth/register", {email, password,});
            
            alert("Usuario cadastrador com sucesso");
            
            navigate("/login");
        } catch (err: any) {
            const message = err.response?.data?.error ?? "Erro ao encontrar usuario";
            
            setError(message);
        } finally {
            setLoading(false);
        }
    }
    return (
        <div>
            <h1>Cadastro</h1>
            <form onSubmit={handleSubmit}>

                <div>
                    <label>Email</label>
                    <br />
                    <input
                        type="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>
                <br />
                <div>
                    <label>Senha</label>
                    <br />
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <br />
                <div>
                    <label>Confirmar senha</label>
                    <br />
                    <input
                        type="password"
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                        required
                    />
                </div>
                <br />
                {
                    error && (<p style={{ color: "red" }}> {error}</p>)
                }

                <button type="submit" disabled={loading}>
                    {
                        loading ? "Cadastrando..." : "Cadastrar"
                    }
                </button>
            </form>
            <br />
            <p> Já possui uma conta? {" "}
                <Link to="/login"> Fazer Login </Link>
            </p>
        </div>
    );
}