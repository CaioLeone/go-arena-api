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

    return(
        <div
            className="
            flex
            min-h-screen
            items-center
            justify-center
            bg-slate-950
            p-4
            "
        >
            <h1>Login</h1>

            <form onSubmit={handleSubmit}>
                <div
                    className="
                    w-full
                    max-w-md
                    rounded-xl
                    border
                    border-slate-700
                    bg-slate-900
                    p-6
                    shadow-xl
                    sm:p-8
                "
                >
                    <label>Email</label>
                    <br />
                    <input 
                        className="
                        w-full
                        rounded-lg
                        border
                        border-slate-700
                        bg-slate-800
                        px-3
                        py-2
                        text-white
                        outline-none
                        focus:border-amber-500
                    "
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
                        className="
                        w-full
                        rounded-lg
                        border
                        border-slate-700
                        bg-slate-800
                        px-3
                        py-2
                        text-white
                        outline-none
                        focus:border-amber-500
                    "
                        type="password"
                        value={password} 
                        onChange={(e) => setPassword(e.target.value)}
                            required 
                    />
                </div>
                <br />
                    {
                        error && 
                        (
                            <p style={{color: "red"}}>{error}</p>
                        )
                    }
                    <button 
                        className="
                        w-full
                        rounded-lg
                        bg-amber-500
                        px-4
                        py-2
                        font-bold
                        text-slate-950
                        transition
                        hover:bg-amber-400
                        disabled:opacity-50
                    "
                        type="submit" disabled={loading}>
                        {loading ? "Entrando" : "Entrar"}
                    </button>
            </form>
            <br />
            <p>
                Ainda nao possui conta?
                {" "}
                <Link to="/register">Registrar</Link>
            </p>
        </div>
    );
}