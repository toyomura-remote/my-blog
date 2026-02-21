import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAppDispatch } from "../app/hooks";
import { login } from "../features/authSlice";

export const useAuthForm = (authFn: (email: string, password: string, name?: string) => Promise<any>, successPath: string) => {
    const dispatch = useAppDispatch();

    const [name, setName] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");
    const [loading, setLoading] = useState<boolean>(false);
    const navigate = useNavigate();

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setLoading(true);

        try {
            const result = await authFn(email, password, name);
            if (result && result.token) {
                localStorage.removeItem("token");
                localStorage.setItem("token", result.token);
                localStorage.setItem("user", JSON.stringify(result.user));
            }
            dispatch(login(result.user));
            alert("成功しました");
            navigate(successPath);
        } catch (err: any) {
            alert(err.response?.data?.error || "エラーが発生しました");
        } finally {
            setLoading(false);
        }
    };

    return { email, setEmail, password, setPassword, name, setName, loading, handleSubmit };
};
