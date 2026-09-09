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
        <div
            className="
                arena-page
                flex
                min-h-screen
                items-center
                justify-center
                px-4
                py-12
            "
        >

            <div className="w-full max-w-md">

                <div className="mb-8 text-center">

                    <Link
                        to="/"
                        className="
                            text-2xl
                            font-black
                            text-white
                        "
                    >
                        Arena dos
                        <span className="text-amber-400">
                            {" "}Bárbaros
                        </span>
                    </Link>

                    <h1
                        className="
                            mt-8
                            text-3xl
                            font-black
                        "
                    >
                        Entre para a Arena
                    </h1>

                    <p className="mt-2 text-slate-400">
                        Crie sua conta e prepare seu primeiro guerreiro.
                    </p>

                </div>

                <div className="arena-card p-6 sm:p-8">

                    <form
                        onSubmit={handleSubmit}
                        className="space-y-5"
                    >

                        <div>

                            <label className="arena-label">
                                Email
                            </label>

                            <input
                                type="email"
                                value={email}
                                onChange={(e) =>
                                    setEmail(e.target.value)
                                }
                                className="arena-input"
                                placeholder="guerreiro@email.com"
                                required
                            />

                        </div>

                        <div>

                            <label className="arena-label">
                                Senha
                            </label>

                            <input
                                type="password"
                                value={password}
                                onChange={(e) =>
                                    setPassword(e.target.value)
                                }
                                className="arena-input"
                                placeholder="Mínimo de 6 caracteres"
                                required
                            />

                        </div>

                        <div>

                            <label className="arena-label">
                                Confirmar senha
                            </label>

                            <input
                                type="password"
                                value={confirmPassword}
                                onChange={(e) =>
                                    setConfirmPassword(
                                        e.target.value
                                    )
                                }
                                className="arena-input"
                                placeholder="Repita sua senha"
                                required
                            />

                        </div>

                        {error && (
                            <p className="arena-error">
                                {error}
                            </p>
                        )}

                        <button
                            type="submit"
                            disabled={loading}
                            className="
                                arena-button-primary
                                w-full
                            "
                        >
                            {loading
                                ? "Criando conta..."
                                : "Criar Conta"}
                        </button>

                    </form>

                </div>

                <p
                    className="
                        mt-6
                        text-center
                        text-sm
                        text-slate-400
                    "
                >
                    Já possui uma conta?{" "}

                    <Link
                        to="/login"
                        className="
                            font-semibold
                            text-amber-400
                            hover:text-amber-300
                        "
                    >
                        Fazer login
                    </Link>
                </p>

                <div className="mt-6 text-center">

                    <Link
                        to="/"
                        className="
                            text-sm
                            text-slate-500
                            hover:text-slate-300
                        "
                    >
                        ← Voltar para a Arena
                    </Link>
                </div>
            </div>
        </div>
    );
}