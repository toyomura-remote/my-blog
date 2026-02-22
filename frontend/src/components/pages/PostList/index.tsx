import { useEffect, useState } from "react";
import type { Post } from "../../../types/post";
import { postService } from "../../../api/postService";
import { Link, useNavigate } from "react-router-dom";
import * as S from "./style";
import { formatToJST } from "../../../utills/date";

const PostList = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchPosts = async () => {
            try {
                const data = await postService.getPosts();
                setPosts(data);
            } catch (error) {
                console.error("Error fetching posts:", error);
                alert("Error fetching posts");
            } finally {
                setLoading(false);
            }
        };

        fetchPosts();
    }, []);

    return (
        <S.Section>
            <S.SectionContainer>
                {loading && <S.Loading>Loading...</S.Loading>}
                <S.List>
                    {posts.length > 0 ? (
                        posts.map((post) => (
                            <S.Item key={post.did}>
                                <S.Link to={`/posts/${post.did}`}>
                                    <S.Heading>{post.title}</S.Heading>
                                    <S.LinkColumn>
                                        <S.CreatedAt dateTime={post.created_at}>{formatToJST(post.created_at)}</S.CreatedAt>
                                        <S.Author>{post.user.name}</S.Author>
                                    </S.LinkColumn>
                                </S.Link>
                            </S.Item>
                        ))
                    ) : (
                        <S.NoPost>投稿がありません</S.NoPost>
                    )}
                </S.List>
            </S.SectionContainer>
        </S.Section>
    );
};

export default PostList;
