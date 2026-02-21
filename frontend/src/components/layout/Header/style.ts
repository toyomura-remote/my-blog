import styled from "styled-components";
import { media } from "../../../styles/media";
import { theme } from "../../../styles/theme";

export const Header = styled.header`
    margin-bottom: 50px;

    ${media.desktop} {
    }
`;

export const HeaderContainer = styled.div`
    ${media.tablet} {
        max-width: ${theme.Container.default};
        margin: 0 auto;
    }
`;

export const List = styled.ul`
    display: flex;
    justify-content: right;
    align-items: center;
    gap: 20px;
    font-size: 1.4rem;
    ${media.desktop} {
        font-size: 1.6rem;
    }
`;
