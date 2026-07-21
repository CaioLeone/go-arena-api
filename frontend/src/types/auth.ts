export interface RegisterRequest{
    username: string;
    email: string;
    password: string;
}

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    username: string;
    email: string;
    password: string;
}

export interface AuthResponse{
    token: string;
}

export interface User{
    id: string;
    username: string;
    email: string;
}

export interface AuthContextData{
    user: User | null;
    token: string | null;
    isAuthenticated: boolean;
    login: (email: string, password: string) => Promise<void>;
    logout: () => void;
}