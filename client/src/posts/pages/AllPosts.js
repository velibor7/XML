import React, { useState, useEffect, _ } from "react";
import PostList from "../components/PostList";
import { useParams } from 'react-router-dom' 

const AllPosts = () => {
  const [loadedPosts, setLoadedPosts] = useState();
  var id = useParams("id")

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        fetch(
          `http://localhost:8000/posts/profile/${id}`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            console.log(data)
            var posts = [];
            data = data['post']
            for (var i = 0; i < data.length; i++) {
              var post = data[i];
              posts.push(post);
            }
            setLoadedPosts(posts);
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
        <PostList items={loadedPosts}></PostList>
    </>
  );
};

export default AllPosts;
