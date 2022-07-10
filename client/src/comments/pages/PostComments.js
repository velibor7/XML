import React, { useState, useEffect, _ } from "react";
import CommentList from "../components/CommentList";
import { useParams } from 'react-router-dom' 

const AllComments = () => {
  const [loadedComments, setLoadedComments] = useState();
  var id = useParams()['id']

  useEffect(() => {
    const fetchComments = async () => {
      try {
        
        fetch(
          `http://localhost:8000/post/${id}/details`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            console.log(data["comments"])
            setLoadedComments(data["comments"]);
          });
      } catch (err) {
        console.log("error happend")
        console.log(err);
      };
    };
    fetchComments();
  }, []);

  return (
    <>
        <h1>Comments</h1>
        <CommentList items={loadedComments}></CommentList>
        
    </>
  );
};

export default AllComments;
