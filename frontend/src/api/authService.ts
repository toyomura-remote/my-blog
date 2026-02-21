import { apiClient } from "./cliant";

export const authService = {
    Signup: async (email: string, password: string, name?: string) => {
        const response = await apiClient.post("/auth/signup", {
            name,
            email,
            password,
        });
        return response.data;
    },
    Login: async (email: string, password: string) => {
        const response = await apiClient.post("/auth/login", {
            email,
            password,
        });
        return response.data;
    },
};
