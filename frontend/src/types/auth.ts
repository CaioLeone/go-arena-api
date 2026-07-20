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