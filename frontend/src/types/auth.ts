export interface User {
    id: number;
    name: string;
    email: string;
}

export interface InitialAuthState {
    isLogIn:boolean
    user: User | null;
}