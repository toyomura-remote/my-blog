import { apiClient } from './cliant'

export const authService = {
    Signup: async(email:string,password:string) =>{
        const response = await apiClient.post('/auth/signup',{
            email,
            password
        });
        return response.data;
    }
}   