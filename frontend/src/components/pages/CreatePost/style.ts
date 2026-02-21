import styled from "styled-components";
import { media } from "../../../styles/media";
import { theme } from "../../../styles/theme";

export const SectionContainer = styled.div`
    ${media.tablet} {
        max-width: ${theme.Container.default};
        margin: 0 auto;
    }
`;
export const Heading = styled.h2`
    ${media.tablet} {
        margin-bottom: 50px;
        font-size: 2rem;
    }
`;
export const Form = styled.form`
    ${media.tablet} {
        margin-bottom: 30px;
    }
`;
export const ContentArea = styled.div`
    display: flex;
    flex-direction: column;
    ${media.tablet} {
        margin-bottom: 30px;
    }
`;
export const Label = styled.label`
    display: flex;
    flex-direction: column;
    ${media.tablet} {
        margin-bottom: 10px;
        font-size: 1.4rem;
    }
`;
export const InputTitle = styled.input`
    ${media.tablet} {
        margin-bottom: 10px;
        height: 30px;
        font-size: 1.4rem;
    }
`;
export const Content = styled.textarea`
    ${media.tablet} {
        font-size: 1.4rem;
        margin-bottom: 10px;
        height: 30px;
        height: 300px;
    }
`;
export const SubmitButton = styled.button`
    background-color:#000;
    color: #fff;
    padding: 10px;
    font-size: 1.1rem;
    border-radius: 5px;
    ${media.desktop} {
        font-size: 1.4rem;
    }
`;

