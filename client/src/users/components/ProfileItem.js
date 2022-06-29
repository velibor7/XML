import { Link } from "react-router-dom";
import "./ProfileItem.css";

import React, { useContext } from "react";


const ProfileItem = (props) => {

  return (
    <>
      <h1>Profile</h1>
      <div className="profile__item">
        <h4 className="profile__item__firstName">{props.item.profile?.firstName}</h4>
        <div className="profile__item__info">
          <p className="profile__item__firstName">
            First name: {props.item.profile?.firstName}
          </p>
          <p className="profile__item__lastName">
            Last name: {props.item.profile?.lastName}
          </p>
          <p className="profile__item__gender">
            Gender: {props.item.profile?.gender}
          </p>
          <p className="profile__item__biography">
            Biography: {props.item.profile?.biography}
          </p>
        </div>
      </div>
    </>
  );
};

export default ProfileItem;
