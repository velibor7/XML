import { Link } from "react-router-dom";
import "./ConnectionItem.css";

import React, { useContext } from "react";


const ConnectionItem = (props) => {

  return (
    <>
      <div className="connection__item">
        <Link to={`/profiles/${props.item.id}`}>
          <h4 className="connection__item__title">{props.item.subjectId}</h4>
        </Link>
        <div className="connection__item__info">
          <p className="connection__item__description">
            Date created: {props.item.date}
          </p>
        </div>
      </div>
    </>
  );
};

export default ConnectionItem;