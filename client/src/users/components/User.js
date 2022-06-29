import React from "react";
import "./UserItem.css";
import Button from "../../shared/components/FormElements/Button"
const User = (props) => {
  return (
    <>
      <div className="user__item">
        <div className="user__item__info">
          <p className="user__item__name">Firstname: <b>{props.item.first_name}</b></p>
          <p className="user__item__city">Lastname: <b>{props.item.last_name}</b></p>
          <p className="user__item__street">Phone number: <b>{props.item.phone_number}</b></p>
          <p className="user__item__city">Email: <b>{props.item.email}</b></p>
          <p className="user__item__street">Country: <b>{props.item.country}</b></p>
          <p className="user__item__city">City: <b>{props.item.city}</b></p>
          <p className="user__item__street">Street: <b>{props.item.street}</b></p>
          <p className="user__item__city">Loyalty points: <b>{props.item.loyalty_points}</b></p>
          <p className="user__item__street">UserType: <b>{props.item.type}</b></p>
        </div>
        <Button  to={`/users/update/${props.item?.id}`}>
          Update profile
        </Button>
        <Button  inverse to={`/users/delete/${props.item?.id}`}>
          Delete profile
        </Button>
      </div>
    </>
  );
};

export default User;
