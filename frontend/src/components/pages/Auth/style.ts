import styled from "styled-components";
import { theme } from "../../../styles/theme";
import { media } from "../../../styles/media";
import { Link } from "react-router-dom";

export const Section = styled.section`
    min-height: calc(100vh - 250px);
    display: flex;
    flex-direction: column;
    justify-content: center;
`;

export const SectionContainer = styled.div`
    border: 1px solid #ddd;
    border-radius: 6px;
    padding: 20px;
    max-width: ${theme.Container.form};
    width: auto;
        background-color: #fff;
    ${media.tablet} {
        margin: 0 auto;
        width: 100%;
    }
`;

export const Hgroup = styled.hgroup`
    display: flex;
    flex-direction: column;
    gap: 15px;
    margin-bottom: 30px;
    ${media.tablet} {
    }
`;
export const Heading = styled.h2`
    font-size: 2rem;
    ${media.tablet} {
    }
`;
export const subHeading = styled.p`
    font-size: 1.3rem;
    ${media.tablet} {
    }
`;

export const Form = styled.form`
    display: flex;
    flex-direction: column;
    gap: 15px;
    ${media.tablet} {
        gap: 30px;
    }
`;

export const InputArea = styled.div`
    display: flex;
    flex-direction: column;
    gap: 10px;
    ${media.tablet} {
    }
`;
export const Label = styled.label`
    font-size: 1.6rem;
    ${media.tablet} {
        font-size: 1.6rem;
    }
`;
export const Input = styled.input`
    font-size: 1.6rem;
    margin-bottom: 10px;
    ${media.tablet} {
        font-size: 1.6rem;
    }
`;

export const Button = styled.button`
    background-color: #000;
    color: #fff;
    padding: 10px;
    font-size: 1.3rem;
    border-radius: 5px;
    margin-bottom: 10px;
    ${media.desktop} {
        font-size: 1.4rem;
    }
`;
export const SignupLink = styled(Link)`
    text-align: center;
    display: block;
    margin:  0 auto;
    margin-bottom: 10px;
    font-size: 1.4rem;
    span{
        font-weight: bold;
        display: block;
        line-height:2;
            ${media.tablet} {
        display: inline-block;
    }
    }

    ${media.desktop} {
    }
`;
