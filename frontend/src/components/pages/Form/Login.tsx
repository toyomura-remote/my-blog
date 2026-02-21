import { authService } from "../../../api/authService";
import { useAuthForm } from "../../../hooks/useAuthForm";
import * as S from "./style";

const Login = () => {
    const { setEmail, setPassword, handleSubmit, loading } = useAuthForm(authService.Login, "/");

    return (
        <S.Section>
            <S.SectionContainer>
                <S.Hgroup>
                    <S.Heading>ログイン</S.Heading>
                    <S.subHeading>アカウントにログインしてください</S.subHeading>
                </S.Hgroup>
                <S.Form onSubmit={handleSubmit}>
                    <S.InputArea>
                        <S.Label htmlFor="">email</S.Label>
                        <input type="email" onChange={(e) => setEmail(e.target.value)} />
                    </S.InputArea>
                    <S.InputArea>
                        <S.Label htmlFor="">password</S.Label>
                        <S.Input type="password" onChange={(e) => setPassword(e.target.value)} />
                    </S.InputArea>
                    <S.Button type="submit" disabled={loading}>
                        {loading ? <span>読み込み中</span> : "ログイン"}
                    </S.Button>
                    <S.SignupLink to="/signup">アカウントをお持ちでないですか？<span>新規登録</span></S.SignupLink>
                </S.Form>
            </S.SectionContainer>
        </S.Section>
    );
};

export default Login;
