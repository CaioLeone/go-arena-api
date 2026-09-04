import type { ReactNode } from "react";
import Navbar from "./Navbar";
import Sidebar from "./Sidebar";

interface Props {
    children: ReactNode;
}

export default function DashboardLayout({children}: Props){
    return (
 <div className="min-h-screen bg-slate-950 text-slate-100">
            <Navbar />
            <div className="flex">
                <Sidebar />
                <main
                    className="
                        min-w-0
                        flex-1
                        p-4
                        sm:p-6
                        lg:p-8
                    "
                >
                    <div className="mx-auto max-w-7xl">
                        {children}
                    </div>
                </main>
            </div>
        </div>
    );
}