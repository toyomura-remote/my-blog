import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./App.css";
import PostList from "./components/PostList";
import Signup from "./components/Signup";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<PostList />} />
                <Route path="/signup" element={<Signup />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App;
