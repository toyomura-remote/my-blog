import styled from "styled-components";
import { media } from "../../../styles/media";
import { theme } from "../../../styles/theme";

export const SectionContainer = styled.div`
    padding: 30px;
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 8px;
    ${media.tablet} {
        max-width: ${theme.Container.default};
        margin: 0 auto;
    }
`;
export const Heading = styled.h2`
    font-size: 2rem;
    margin-bottom: 20px;
    ${media.desktop} {
        margin-bottom: 20px;
    }
`;
export const CreatedAt = styled.dl`
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 10px;
    font-size: 1.4rem;
    margin-bottom: 15px;
    ${media.desktop} {
        font-size: 1.4rem;
        margin-bottom: 20px;
    }
`;
export const ContentText = styled.p`
    font-size: 1.4rem;
    margin-bottom: 20px;
    line-height: 35px;
    border-top: 1px solid #ccc;
    padding-top: 30px;
    margin-bottom: 70px;
    white-space: pre-wrap;
    word-break: break-all;
    ${media.desktop} {
        font-size: 1.6rem;
    }
`;
export const ButtonArea = styled.div`
    display: flex;
    gap: 15px;
`;
export const DeleteButton = styled.button`
    background-color: #333;
    color: #fff;
    padding: 10px;
    font-size: 1.1rem;
    border-radius: 5px;
    ${media.desktop} {
        font-size: 1.4rem;
    }
`;
export const EditButton = styled.button`
    background-color: #fff;
    padding: 10px;
    font-size: 1.1rem;
    border-radius: 5px;
    border: 1px solid #ddd;
    ${media.desktop} {
        font-size: 1.4rem;
    }
`;
