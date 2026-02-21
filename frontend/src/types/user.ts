import type { Post } from "./post";

export interface User{
    name :string;
    did:string;
    post:Post[] | null;
}