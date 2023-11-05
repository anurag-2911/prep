import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Home from "./Home";
import About from "./About";
import Posts from "./Posts";

function App() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/posts">Posts</Link>
            </li>
            {/* Adding some sample links with parameters and queries */}
            <li>
              <Link to="/posts/1?author=John">Post 1 by John</Link>
            </li>
            <li>
              <Link to="/posts/2?author=Jane">Post 2 by Jane</Link>
            </li>
          </ul>
        </nav>

        <Routes>
          <Route path="/about" element={<About />} />
          <Route path="/posts/:postId" element={<Posts />} />
          <Route path="/" element={<Home />} />
        </Routes>
      </div>
    </Router>
  );
}


export default App;
