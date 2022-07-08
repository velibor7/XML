import { Link } from "react-router-dom";
import "./ProfileItem.css";

import React, { useContext } from "react";


const ProfileItemForList = (props) => {

  return (
    <>
      <div className="profile__item">
        <Link to={`/profiles/${props.item.id}`}>
          <h4 className="profile__item__firstName">{props.item.username}</h4>
        </Link>
        <div className="profile__item__info">
          <p className="profile__item__firstName">
            First name: {props.item.firstName}
          </p>
          <p className="profile__item__lastName">
            Last name: {props.item.lastName}
          </p>
          <p className="profile__item__gender">
            Gender: {props.item.gender}
          </p>
          <p className="profile__item__biography">
            Biography: {props.item.biography}
          </p>
        </div>
      </div>
    </>
  );
};

export default ProfileItemForList;
