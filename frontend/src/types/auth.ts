export interface RegisterRequest{
    user: string;
    email: string;
    password: string;
}

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    user: string;
    email: string;
    password: string;
}

export interface AuthResponse{
    token: string;
}

export interface User{
    id: string;
    name: string;
    email: string;
}

export interface AuthContextData{
    user: User | null;
    token: string | null;
    loading: boolean;
    isAuthenticated: boolean;
    login: (email: string, password: string) => Promise<void>;
    logout: () => void;
}