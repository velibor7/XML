import "./PostItem.css";

import React, { useContext } from "react";


const PostItemForList = (props) => {

  return (
    <>
      <div className="post__item">
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
