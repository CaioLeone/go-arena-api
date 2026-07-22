import {
    createContext, 
    useState, 
    useEffect 
} from "react";

import api from "../services/api";
import type {AuthContextData, AuthResponse,User} from "../types/auth";

const TOKEN_KEY = "@arena:token";

interface AuthProviderProps {
    children: React.ReactNode;
}

export const AuthContex = createContext({} as AuthContextData);
export function AuthProvider({children}: AuthProviderProps){
        const [user, setUser] = useState<User | null>(null);
        const [token, setToken] = useState<string | null>(null);
        const [loading, setLoading] = useState(true);

    const isAuthenticated = !!token;

    useEffect(() => {
        async function loadUser() {
            const storedToken = localStorage.getItem(TOKEN_KEY);
            if(!storedToken){
                setLoading(false);
                return;
            }

            try {
                setToken(storedToken);
                const response = await api.get("/api/me");

                setUser({
                    id: response.data.user_id,
                    user: "",
                    email: "",
                });
            } catch (error) {
                localStorage.removeItem(TOKEN_KEY);
                setToken(null);
                setUser(null);
            } finally {
                setLoading(false);
            }
        }
        loadUser();
    }, []);

    async function login(email: string, password: string) {
        const response = await api.post<AuthResponse>("/auth/login", { email, password });
        const token = response.data.token;
        localStorage.setItem(TOKEN_KEY, token);
        setToken(token);

        const me = await api.get("/api/me");
        setUser({
            id: me.data.user_id,
            user: "",
            email: "",
        });
    }

    function logout(){
        localStorage.removeItem(TOKEN_KEY);
        setToken(null);
        setUser(null);
    }

    return (
        <AuthContex.Provider value={{user, token, loading, isAuthenticated, login, logout}}>
            {children}
        </AuthContex.Provider>
    );
}

