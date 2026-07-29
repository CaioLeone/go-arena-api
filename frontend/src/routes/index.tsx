import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "../pages/Home";
import Login from "../pages/Login";
import Register from "../pages/Register";
import Dashboard from "../pages/Dashboard";
import Characters from "../pages/Characters";
import Battles from "../pages/Battles";
import Leaderboard from "../pages/Leaderboard";

import { PrivateRoute } from "../components/PrivateRoute";

export default function AppRoutes() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/register" element={<Register />} />
                <Route 
                    path="/dashboard" 
                    element={
                        <PrivateRoute> 
                            <Dashboard/>
                        </PrivateRoute>}
                />

                <Route 
                    path="/characters" 
                    element={
                        <PrivateRoute> 
                            <Characters/>
                        </PrivateRoute>}
                />

                <Route 
                    path="/battles" 
                    element={
                        <PrivateRoute> 
                            <Battles/>
                        </PrivateRoute>}
                />

                <Route 
                    path="/leaderboard" 
                    element={
                        <PrivateRoute> 
                            <Leaderboard/>
                        </PrivateRoute>}
                />
            </Routes>
        </BrowserRouter>
    );
}