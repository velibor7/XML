import Button from '../../shared/components/FormElements/Button'
import './UserItem.css'

import React from "react";

const UserItem = (props) => {


  return (
    <>
      <div className='user__item'>
        <h4 className="user__item__name">{props.item.first_name}</h4>
          <div className="user__item__info">
            <p className="user__item__last_name">{props.item.last_name}</p>
            <p className="user__item__phone_number">{props.item.phone_number}</p>
            <p className="user__item__country">{props.item.country}</p>
            <p className="user__item__city">{props.item.city}</p>
            <p className="user__item__street">{props.item.street}</p>
            <p className="user__item__loyalty_points">{props.item.loyalty_points}</p>
            <p className="user__item__type">{props.item.type}</p>
          </div>
          {/* Samo ako je admin vidi ovo */}
          <Button danger to={`/admin/make/${props.item.id}`}>
            Make admin
          </Button>
        </div>
    </>
  )
}

export default UserItem