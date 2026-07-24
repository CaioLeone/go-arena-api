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
        <div>
            <h1>Login</h1>

            <form onSubmit={handleSubmit}>
                <div>
                    <label>Email</label>
                    <br />
                    <input type="email"
                           value={email} 
                           onChange={(e) => setEmail(e.target.value)}
                           required 
                    />
                </div>
                <br />

                <div>
                    <label>Senha</label>
                    <br />
                    <input type="password"
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
                    <button type="submit" disabled={loading}>
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