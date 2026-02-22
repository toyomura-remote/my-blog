import { authService } from "../../../api/authService";
import { useAuthForm } from "../../../hooks/useAuthForm";
import * as S from "./style";

const Signup = () => {
    const { setName,setEmail, setPassword, handleSubmit } = useAuthForm(authService.Signup, "/");

    return (
        <S.Section>
            <S.SectionContainer>
                <S.Hgroup>
                    <S.Heading>新規登録</S.Heading>
                    <S.subHeading>アカウントを作成してください</S.subHeading>
                </S.Hgroup>
                <S.Form onSubmit={handleSubmit}>
                    <S.InputArea>
                        <S.Label htmlFor="">name</S.Label>
                        <S.Input type="text" onChange={(e) => setName(e.target.value)} />
                    </S.InputArea>
                    <S.InputArea>
                        <S.Label htmlFor="">email</S.Label>
                        <S.Input type="email" onChange={(e) => setEmail(e.target.value)} />
                    </S.InputArea>
                    <S.InputArea>
                        <S.Label htmlFor="">password</S.Label>
                        <S.Input type="password" onChange={(e) => setPassword(e.target.value)} />
                    </S.InputArea>
                    <S.Button type="submit">登録</S.Button>
                    <S.SignupLink to="/login">
                        アカウントをお持ちの場合<span>ログイン</span>
                    </S.SignupLink>
                </S.Form>
            </S.SectionContainer>
        </S.Section>
    );
};

export default Signup;
