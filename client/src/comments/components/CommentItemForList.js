import "./CommentItem.css";
import { Link } from "react-router-dom";
import React, { useContext } from "react";


const CommentItemForList = (props) => {

  return (
    <>
      <div className="comment__item">
        <div className="comment__item__info">
          <h4 className="comment__item__text">
          {props.item.username}
          </h4>
          <p className="comment__item__links">
            {props.item.content}
          </p>
        </div>
      </div>
      <hr/>
    </>
  );
};

export default CommentItemForList;
