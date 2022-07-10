import { Link } from "react-router-dom";
import "./JobItem.css";

import React, { useContext } from "react";


const JobItem = (props) => {

  return (
    <>
      <div className="job__item">
        <Link to={`/jobs/${props.item.id}`}>
          <h4 className="job__item__title">{props.item.title}</h4>
        </Link>
        <div className="job__item__info">
          <p className="job__item__title">
            Title: {props.item.title}
          </p>
          <p className="job__item__description">
            Description: {props.item.description}
          </p>
          <p className="job__item__skills">
            Skills: {props.item.skills.map((item) => (
                    <p> {item} </p>
                  ))}
          </p>
        </div>
      </div>
    </>
  );
};

export default JobItem;