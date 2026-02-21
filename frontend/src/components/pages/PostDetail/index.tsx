import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { postService } from "../../../api/postService";
import type { Post } from "../../../types/post";
import * as S from "./style";
import { formatToJST } from "../../../utills/date";
import { useSelector } from "react-redux";
import type { RootState } from "../../../app/store";

const PostDetail = () => {
    const { did } = useParams<{ did: string }>();
    const [post, setPost] = useState<Post | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const { user: currentUser} = useSelector((state: RootState) => state.auth);
    const navigate = useNavigate();

    const handleDelete = async (did: string) => {
        try {
            await postService.deletePost(did);
            alert('削除しました');
            navigate("/");
        } catch (err:any) {
            console.error(err);
            alert(err.response?.data?.error);
        }
    };

    console.log(post?.user_id);
    

    useEffect(() => {
        const fetchDetail = async () => {
            if (!did) return;

            try {
                setLoading(true);
                const data = await postService.getPostByDid(did);
                setPost(data);
            } catch (err) {
                console.error(err);
            } finally {
                setLoading(false);
            }
        };
        fetchDetail();
    }, [did]);

    if (!post) return <div>Loading...</div>;

    return (
        <section>
            <S.SectionContainer>
                <S.Heading>{post.title}</S.Heading>
                <S.CreatedAt>
                    <dt><time dateTime={post.created_at}>{formatToJST(post.created_at)}</time></dt>
                    <dd>
                        {post.user.name}
                    </dd>
                </S.CreatedAt>
                <S.ContentText>{post.content}</S.ContentText>
                {currentUser && currentUser.id === post.user_id && (
                    <S.ButtonArea>
                        <S.EditButton>編集する</S.EditButton>
                        <S.DeleteButton onClick={() => did && handleDelete(did)}>削除する</S.DeleteButton>
                    </S.ButtonArea>
                )}
            </S.SectionContainer>
        </section>
    );
};

export default PostDetail;
