import React, { useState, useEffect, _ } from "react";
import PostList from "../components/PostList";
import { useParams } from 'react-router-dom' 

const AllPosts = () => {
  const [loadedPosts, setLoadedPosts] = useState();
  var id = useParams()['id']

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        fetch(
          `http://localhost:8000/posts/profile/${id}`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            console.log(data["posts"])
            setLoadedPosts(data["posts"]);
          });
      } catch (err) {
        console.log("error happend")
        console.log(err);
      };
    };
    fetchPosts();
  }, []);

  return (
    <>
        <h1>Posts</h1>
        <PostList items={loadedPosts}></PostList>
    </>
  );
};

export default AllPosts;
