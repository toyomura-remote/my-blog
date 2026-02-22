import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { postService } from "../../../api/postService";
import * as S from "./style";
import { PostForm } from "../postForm";

const CreatePost = () => {

    const navigate = useNavigate();

    const token = localStorage.getItem("token");
    if (!token) {
        alert("ログインしてください");
        navigate("/login");
        return;
    }

    const handleCreate = async (title: string, content: string) => {
        try {
            await postService.createPost(title, content);
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
                <PostForm submitLabel="投稿" onSubmit={handleCreate} />
            </S.SectionContainer>
        </section>
    );
};

export default CreatePost;
