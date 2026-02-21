import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { InitialAuthState, User } from "../types/auth";

const savedUser = localStorage.getItem("user");
const initialUser = savedUser ? JSON.parse(savedUser) : null;

const initialState: InitialAuthState = {
    isLogIn: localStorage.getItem("token") !== null,
    user: initialUser,
};

export const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        login: (state, action: PayloadAction<User>) => {
            state.isLogIn = true;
            state.user = action.payload;
        },
        logout: (state) => {
            state.isLogIn = false;
            state.user = null;
        },
    },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;
