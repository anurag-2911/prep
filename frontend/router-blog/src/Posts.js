import {useParams,useLocation} from "react-router-dom";

function Posts(){
    const {postId}= useParams();
    const location= useLocation();
    const queryParams= new URLSearchParams(location.search);
    const author = queryParams.get('author');

    return (
        <div>
            <h2>Posts Page</h2>
            {postId && <p>Post ID: {postId}</p>}
            {author && <p>Author: {author}</p>}
        </div>
    );
}
export default Posts;
