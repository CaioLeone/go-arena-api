import axios from "axios";

console.log("API URL:", import.meta.env.VITE_API_URL);

const api = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
    headers: {
        "Content-Type": "application/json",
    },
});

api.interceptors.request.use(
    (config)=> {
        const token = localStorage.getItem("@arena:token");

        console.log("TOKEN ENVIADO:", token);
        
        if(token){
            config.headers.Authorization = `Bearer ${token}`;
        }

        console.log("REQUEST:", config.url);
        
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default api;