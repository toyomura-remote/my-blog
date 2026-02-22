import { BrowserRouter, Route, Routes } from "react-router-dom";
import { ThemeProvider } from "styled-components";
import { theme } from "./styles/theme";
import "./App.css";
import Signup from "./components/pages/Auth/Signup";
import Header from "./components/layout/Header";
import { GlobalStyle } from "./styles/GlobalStyle";
import Login from "./components/pages/Auth/Login";
import PostDetail from "./components/pages/PostDetail";
import CreatePost from "./components/pages/CreatePost";
import Home from "./components/pages/Home";

function App() {
    return (
        <ThemeProvider theme={theme}>
            <GlobalStyle />
            <BrowserRouter>
                <Header />
                <main>
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/posts/:did" element={<PostDetail />} />
                        <Route path="/posts/new" element={<CreatePost />} />
                        <Route path="/signup" element={<Signup />} />
                        <Route path="/login" element={<Login />} />
                    </Routes>
                </main>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
