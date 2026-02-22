import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { postService } from "../../../api/postService";
import * as S from "./style";
import { PostForm } from "../postForm";

const UpdatePost = () => {
    const navigate = useNavigate();
    const { did } = useParams<{ did: string }>();
    const [post, setPost] = useState<{ title: string; content: string } | null>(null);
    const token = localStorage.getItem("token");
    if (!token) {
        alert("ログインしてください");
        navigate("/login");
        return;
    }

    useEffect(() => {
        if (!did) return;
        const fetchPost = async () => {
            try {
                const data = await postService.getPostByDid(did);
                setPost({ title: data.title, content: data.content });
            } catch (err) {
                alert("投稿の取得に失敗しました");
                console.error(err);
            }
        };
        fetchPost();
    }, [did, navigate]);

    const handleUpdate = async (did: string, title: string, content: string) => {
        if (!did) {
            alert("投稿IDが見つかりません");
            return;
        }
        try {
            await postService.updatePost(did, title, content);
            alert("投稿を更新しました");
            navigate("/");
        } catch (err) {
            alert("投稿の更新に失敗しました");
            console.error(err);
        }
    };

    return (
        <section>
            <S.SectionContainer>
                <S.Heading>記事編集</S.Heading>
                {post && <PostForm initialTitle={post.title} initialContent={post.content} submitLabel="更新" onSubmit={(title, content) => handleUpdate(did!, title, content)} />}
            </S.SectionContainer>
        </section>
    );
};

export default UpdatePost;
