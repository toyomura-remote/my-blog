import React, { useEffect, useState } from "react";
import type { Post } from "../types/post";
import { postService } from "../api/postService";

const postList = () => {
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState<boolean>(false);

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

    if (loading) return <div>Loading...</div>;

    return (
        <div>
            <h1>投稿一覧</h1>
            {posts.length > 0 
            ?posts.map((post)=>
            <div key={post.did}>
                <h2>{post.title}</h2>
                <p>{post.content}</p>
                <small>投稿者:{post.user?.name}</small>
            </div>
            ):<p>投稿がありません。</p>
        }   
        </div>
    )
};

export default postList;
