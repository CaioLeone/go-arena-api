import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export default function Login(){
    const navigate = useNavigate();
    const { login } = useAuth();
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    async function handleSubmit(event: React.FormEvent<HTMLFormElement>){
        event.preventDefault();
        setError("");
        setLoading(true);

        try {
            await login(email, password);
            navigate("/dashboard");
        } catch (err) {
            setError("Email ou Senha Invalidos");
        }finally{
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
                    Bem-vindo de volta
                </h1>

                <p className="mt-2 text-slate-400">
                    Entre para continuar sua jornada.
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
                            placeholder="Sua senha"
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
                            ? "Entrando..."
                            : "Entrar"}
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
                Ainda não possui uma conta?{" "}

                <Link
                    to="/register"
                    className="
                        font-semibold
                        text-amber-400
                        hover:text-amber-300
                    "
                >
                    Criar conta
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