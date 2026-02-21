import { Link } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../../app/hooks";
import { logout } from "../../../features/authSlice";
import * as S from "./style";

const Header = () => {
    const isLogin = useAppSelector((state) => state.auth.isLogIn);

    const dispatch = useAppDispatch();

    const handleLogout = () => {
        dispatch(logout());

        alert("ログアウトしました");
        localStorage.clear();
    };

    return (
        <S.Header>
            <S.HeaderContainer>
                <nav>
                    <S.List>
                        <li>
                            <Link to="/">ホーム</Link>
                        </li>
                        {isLogin ? (
                            <>
                                <li>
                                    <Link to="/posts/new">記事作成</Link>
                                </li>
                                <li>
                                    <Link to="/mypage">マイページ</Link>
                                </li>
                                <li>
                                    <button onClick={handleLogout}>ログアウト</button>
                                </li>
                            </>
                        ) : (
                            <>
                                <li>
                                    <Link to="/login">ログイン</Link>
                                </li>
                                <li>
                                    <Link to="/signup">新規登録</Link>
                                </li>
                            </>
                        )}
                    </S.List>
                </nav>
            </S.HeaderContainer>
        </S.Header>
    );
};

export default Header;
