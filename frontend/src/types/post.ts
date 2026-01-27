import type { User } from "./user";

export interface Post {
    did: string;
    title: string;
    content: string;
    user: User | null;
}

    export interface PostResponse {
        data: Post[];
    } 