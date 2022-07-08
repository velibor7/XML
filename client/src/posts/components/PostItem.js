import { Link } from "react-router-dom";
import "./PostItem.css";

import React, { useContext } from "react";


const PostItemForList = (props) => {

  return (
    <>
      <div className="post__item">
        <div className="post__item__info">
          <p className="post__item__firstName">
            First name: {props.item.firstName}
          </p>
          <p className="post__item__lastName">
            Last name: {props.item.lastName}
          </p>
          <p className="post__item__gender">
            Gender: {props.item.gender}
          </p>
          <p className="post__item__biography">
            Biography: {props.item.biography}
          </p>
        </div>
      </div>
    </>
  );
};

export default PostItemForList;
