import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { postService } from "../../../api/postService";
import * as S from "./style";

const New = () => {
    const [titleInput, SetTitleInput] = useState<string>("");
    const [contentInput, SetContentInput] = useState<string>("");
    const navigate = useNavigate();

    const token = localStorage.getItem("token");
    if (!token) {
        alert("ログインしてください");
        navigate("/login");
        return;
    }

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            await postService.createPost(titleInput, contentInput);
            SetTitleInput("");
            SetContentInput("");
            alert('投稿しました');
            navigate("/");
        } catch (err) {
            alert('投稿に失敗しました');
            console.error(err);
        }
    };

    return (
        <section>
            <S.SectionContainer>
                <S.Heading>記事作成</S.Heading>
                <S.Form action="" onSubmit={handleSubmit}>
                    <S.ContentArea>
                        <S.Label htmlFor="inputTitle">タイトル</S.Label>
                        <S.InputTitle id="inputTitle" type="text" onChange={(e) => SetTitleInput(e.target.value)} value={titleInput} />
                    </S.ContentArea>
                    <S.ContentArea>
                        <S.Label htmlFor="Content">本文</S.Label>
                        <S.Content id="Content" onChange={(e) => SetContentInput(e.target.value)} value={contentInput} />
                    </S.ContentArea>
                    <S.SubmitButton type="submit">投稿</S.SubmitButton>
                </S.Form>
            </S.SectionContainer>
        </section>
    );
};

export default New;
