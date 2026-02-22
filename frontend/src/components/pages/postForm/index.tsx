import { useEffect, useState } from "react";
import * as S from "./style";

type props = {
    initialTitle?: string;
    initialContent?: string;
    submitLabel: string;
    onSubmit: (title: string, content: string) => Promise<void>;
};

export const PostForm = ({ initialTitle = "", initialContent = "", submitLabel, onSubmit }: props) => {
    const [titleInput, SetTitleInput] = useState<string>(initialTitle);
    const [contentInput, SetContentInput] = useState<string>(initialContent);

    useEffect(() => {
        SetTitleInput(initialTitle);
        SetContentInput(initialContent);
    }, [initialTitle, initialContent]);

    return (
        <S.Form  onSubmit={(e: React.FormEvent) => {e.preventDefault(); onSubmit(titleInput, contentInput);}}>
            <S.ContentArea>
                <S.Label htmlFor="inputTitle">タイトル</S.Label>
                <S.InputTitle id="inputTitle" type="text" value={titleInput} onChange={(e) => SetTitleInput(e.target.value)} />
            </S.ContentArea>
            <S.ContentArea>
                <S.Label htmlFor="Content">本文</S.Label>
                <S.Content id="Content" value={contentInput} onChange={(e) => SetContentInput(e.target.value)} />
            </S.ContentArea>
            <S.SubmitButton type="submit">{submitLabel}</S.SubmitButton>
        </S.Form>
    );
};
