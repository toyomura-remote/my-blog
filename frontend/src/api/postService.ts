import type { Post, PostListResponse, PostSingleResponse } from "../types/post";
import { apiClient } from "./cliant";

export const postService = {
    createPost: async (title: string, content: string) => {
        await apiClient.post(
            "/posts",
            {
                title,
                content,
            },
        );
    },
    deletePost: async (did: string) => {
        await apiClient.delete(
            `posts/${did}`
        );
    },
    getPosts: async (): Promise<Post[]> => {
        const response = await apiClient.get<PostListResponse>("/posts");
        return response.data.data;
    },
    getPostByDid: async (did: string): Promise<Post> => {
        const response = await apiClient.get<PostSingleResponse>(`/posts/${did}`);
        return response.data.data;
    },
};
