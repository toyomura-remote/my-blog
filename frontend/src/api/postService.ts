import type { Post, PostResponse } from "../types/post";
import { apiClient } from "./cliant";

export const postService = {
    getPosts: async (): Promise<Post[]> => {
        const response = await apiClient.get<PostResponse>("/posts");
        return response.data.data;
    },
    getPostByDid: async (did: string): Promise<Post> => {
        const response = await apiClient.get<Post>(`/posts/${did}`);
        return response.data;
    },
};
