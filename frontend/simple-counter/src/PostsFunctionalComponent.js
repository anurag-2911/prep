import React from "react";
import { useEffect, useState } from "react";

function PostFunctionalComponent(props) {
  const [posts, setPosts] = useState([]);
  useEffect(() => {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then((response) => response.json())
      .then((data) => setPosts(data));
  }, []); // The empty dependency array means this useEffect will run once after the initial render.

  return (
    <>
      <div>
        <h1>Posts (Functional Component)</h1>
        <ul>
          {posts.map((post) => (
            <li key={post.id}>{post.title}</li>
          ))}
        </ul>
      </div>
    </>
  );
}
export default PostFunctionalComponent;
