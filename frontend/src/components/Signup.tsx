import React, { use, useState } from "react";
import { useNavigate } from "react-router-dom";
import { authService } from "../api/authService";

const Signup = () => {
    const [email,setEmail] = useState<string>("")
    const [password,setPassword] = useState<string>("")
    const navigate = useNavigate();

    const onSubmit = async(e:React.FormEvent) => {
        e.preventDefault();
        try{
            await authService.Signup(email,password)
            navigate("/");
        }catch(err){
            console.error(err)
            alert("登録に失敗しました")
        }
    }

    return (
        <form onSubmit={onSubmit}>
            <input type="email" onChange={(e) => setEmail(e.target.value)}/>
            <input type="password" onChange={(e) => setPassword(e.target.value)}/>
            <button type="submit">登録</button>
        </form>
    );
};

export default Signup;
