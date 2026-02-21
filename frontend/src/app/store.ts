import { configureStore } from "@reduxjs/toolkit";
import authReducer from "../features/authSlice";

export const store = configureStore({
    reducer:{
        auth: authReducer
    },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;

//getStateは呼んだ結果が欲しいがdispatchは機能そのものが欲しい
// RootState: 「魔法の箱（Store）を開けるボタン（getState）」を**押した後に中から出てくる「宝物」**のリスト。
// AppDispatch: 「魔法をかけるための杖（dispatch）」そのもの。杖の形や振り方が分かればいいので、振った後の結果を気にする必要はありません。