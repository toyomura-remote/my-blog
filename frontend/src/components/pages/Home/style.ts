import styled from "styled-components";
import { media } from "../../../styles/media";
import { theme } from "../../../styles/theme";

export const Section = styled.section`
margin-bottom: 50px;
    ${media.desktop} {
        margin-bottom: 70px;
    }
`;
export const SectionContainer = styled.div`
    ${media.tablet} {
        max-width:${theme.Container.default} ;
        margin: 0 auto;
    }
`;
export const Hgroup = styled.hgroup`
    display: flex;
    flex-direction: column;
    gap: 20px;
    ${media.desktop} {

    }
`;
export const Heading = styled.h1`
font-size: 3rem;
    ${media.desktop} {
        font-size: 4rem;
    }
`;
export const Paragraph = styled.p`
font-size: 1.6rem;
    ${media.desktop} {
        font-size: 1.6rem;
    }
`;
