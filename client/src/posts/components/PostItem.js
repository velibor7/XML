import "./PostItem.css";
import React, { useContext } from "react";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";
import { useParams } from 'react-router-dom' 
import Button from "../../shared/components/FormElements/Button";
import { useNavigate } from "react-router-dom";


const PostItem = (props) => {
  const id = useParams()['userId']
  const { sendRequest } = useHttpClient();
  const navigate = useNavigate();
  const auth = useContext(AuthContext);
  
  const LikePost = async () => {
    for(let i=0; i<props.item.post?.likes.length; i++){
      if(props.item.post?.likes[i].userId == auth.userId){
        props.item.post?.likes.splice(i,1);
      }
    }
    for(let i=0; i<props.item.post?.dislikes.length; i++){
      if(props.item.post?.dislikes[i].userId == auth.userId){
        props.item.post?.dislikes.splice(i,1);
      }
    }

    try {
      props.item.post?.likes.push({userId : auth.userId})
      let body = props.item.post;
      await sendRequest(
        `http://localhost:8000/posts/${props.item.post?.id}`,
        "PUT",
        JSON.stringify(body),
        {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
        }
      );
      navigate("/posts/" + auth.userId +"/" + props.item.post?.id);
    } catch (err) {
      console.log(err);
    }

  };

  const DislikePost = async () => {
    for(let i=0; i<props.item.post?.likes.length; i++){
      if(props.item.post?.likes[i].userId == auth.userId){
        props.item.post?.likes.splice(i,1);
      }
    }
    for(let i=0; i<props.item.post?.dislikes.length; i++){
      if(props.item.post?.dislikes[i].userId == auth.userId){
        props.item.post?.dislikes.splice(i,1);
      }
    }

    try {
      props.item.post?.dislikes.push({userId : auth.userId})
      let body = props.item.post;
      await sendRequest(
        `http://localhost:8000/posts/${props.item.post?.id}`,
        "PUT",
        JSON.stringify(body),
        {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
        }
      );
      navigate("/posts/" + auth.userId +"/" + props.item.post?.id);
    } catch (err) {
      console.log(err);
    }
  };

  const CommentsNavigate = async () => {
    navigate("/comments/" + props.item.post?.id);
  }

  return (
    <>
    <h1>A</h1>
    <h1>  </h1>
      <div className="post__item">
        <div className="post__item__info">
          <p className="post__item__text">
          text: {props.item.post?.text}
          </p>
          <p className="post__item__links">
            links: {props.item.post?.links}
          </p>
          <p className="post__item__created">
          created: {props.item.post?.created}
          </p>
          <p className="post__item__likes">
          likes: {props.item.post?.likes.length}
          </p>
          <p className="post__item__dislikes">
          dislikes: {props.item.post?.dislikes.length}
          </p>
          {
            (auth.userId == id) &&
            (<Button success onClick={LikePost}>
              Like
            </Button>
            )
          }
          {
            (auth.userId == id) &&
            (<Button danger onClick={DislikePost}>
              Dislike
            </Button>
            )
          }
          <Button danger onClick={CommentsNavigate}>
              Comments
          </Button>
            
        </div>
      </div>
    </>
  );
};

export default PostItem;
