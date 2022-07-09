import "./PostItem.css";
import { Link } from "react-router-dom";
import React, { useContext } from "react";


const PostItemForList = (props) => {

  return (
    <>
      <div className="post__item">
        <Link to={`/posts/${props.item.userId}/${props.item.id}`}>
          <h4 className="profile__item__firstName">Post</h4>
        </Link>
        <div className="post__item__info">
          <p className="post__item__text">
          text: {props.item.text}
          </p>
          <p className="post__item__links">
            links: {props.item.links}
          </p>
          <p className="post__item__created">
          created: {props.item.created}
          </p>
          <p className="post__item__likes">
          likes: {props.item.likes.length}
          </p>
          <p className="post__item__dislikes">
          dislikes: {props.item.dislikes.length}
          </p>
        </div>
      </div>
    </>
  );
};

export default PostItemForList;
