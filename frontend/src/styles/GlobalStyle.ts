import { createGlobalStyle } from "styled-components";
import normalize from "styled-normalize";
import { media } from "./media";

export const GlobalStyle = createGlobalStyle`
// reset css
    ${normalize}
    html{
        font-size: 62.5%;
    }
    body{
        background-color: #ffffff;
        color: #333333;
        padding: 30px;
    }
    h1{
        margin: 0;
    }
    h2{
        margin: 0;
        font-weight: unset;
    }
    p{
        margin: 0;
    }
    dl,dt,dd{
        margin: 0;
    }
    ul{
        padding: 0;
        margin: 0;
    }
    li{
        list-style:none;

    }
    a{
        text-decoration:none;
        color: #333333;
    }
    input,textarea{
        border: 1px solid #ddd;
        border-radius: 6px;
        padding:1rem;
    }
    textarea{
        resize: none;
    }
    button{
        border: unset;
        background-color: unset;
        &:hover{
            cursor: pointer;
        }
    }
    `;
