import * as S from "./style";
import PostList from "../PostList";

const Home = () => {
    return (
        <>
            <S.Section>
                <S.SectionContainer>
                    <S.Hgroup>
                        <S.Heading>My Blog</S.Heading>
                        <S.Paragraph>開発用のブログ</S.Paragraph>
                    </S.Hgroup>
                </S.SectionContainer>
            </S.Section>
            <PostList />
        </>
    );
};

export default Home;
