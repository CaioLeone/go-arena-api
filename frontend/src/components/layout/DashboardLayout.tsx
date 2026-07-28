import type { ReactNode } from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";

interface Props {
    children: ReactNode;
}

export default function DashboardLayout({children}: Props){
    return (
        <>
            <Navbar/>
            <div 
                style={{
                    display: "flex",
                }}
            >
                <Sidebar/>
                <main 
                    style={{flex: 1, padding: "30px"}}>
                        {children}
                    </main>
            </div>
        </>
    );
}