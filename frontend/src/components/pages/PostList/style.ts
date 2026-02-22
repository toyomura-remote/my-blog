import styled from "styled-components";
import { media } from "../../../styles/media";
import { Link as RouterLink } from "react-router-dom";
import { theme } from "../../../styles/theme";

export const Section = styled.div`
    ${media.desktop} {
        padding-bottom: 70px;
    }
`;
export const SectionContainer = styled.div`
    margin: 0 auto;
    ${media.tablet} {
        max-width: ${theme.Container.default};
    }
    ${media.desktop} {
    }
`;
export const List = styled.ul`
    display: flex;
    flex-direction: column;
    gap: 20px;
`;
export const Item = styled.li`
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 8px;
    ${media.desktop} {
    }
`;
export const Link = styled(RouterLink)`
    display: block;
    padding: 20px;
    transition: 0.4s;
    &:hover {
        opacity: 0.6;
    }
`;
export const LinkColumn = styled.div`
    display: flex;
    justify-content: flex-start;
    gap: 15px;
`;
export const Heading = styled.h2`
    margin-bottom: 20px;
    ${media.desktop} {
        font-size: 1.6rem;
    }
`;
export const CreatedAt = styled.time`
    ${media.desktop} {
        font-size: 1.2rem;
    }
`;
export const Author = styled.p`
    font-size: 1.2rem;
    ${media.desktop} {
    }
`;
export const NoPost = styled.p`
    font-size: 1.2rem;
    ${media.desktop} {
    }
`;
export const Loading = styled.p`
    font-size: 1.2rem;
    ${media.desktop} {
    }
`;
