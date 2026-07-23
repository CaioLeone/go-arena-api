import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

interface PrivateRouteProps {
    children: React.ReactNode;
}

export function PrivateRoute({ children,}: PrivateRouteProps) {
    const {
        isAuthenticated,
        loading,
    } = useAuth();

    if (loading) {
        return <p>Carregando...</p>
    }

    if(!isAuthenticated) {
        return <Navigate to="/login" replace />;
    }

    return children;
}