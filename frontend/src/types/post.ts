import type { User } from "./user";

export interface Post {
    user_id: number;
    did: string;
    title: string;
    content: string;
    created_at: string;
    user: User;
}

export interface PostListResponse {
    data: Post[];
}

export interface PostSingleResponse {
    data: Post;
}