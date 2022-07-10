import ConnectionRequestItem from "./ConnectionRequestItem";
import React from "react";

const ConnectionRequestsList = (props) => {
  return (
    <div className="connections">
      <h1>Connection Requests</h1>
      <div className="connection-list__container">
        <div className="connection__wrapper">
          <div className="connection-list__items">
            {props.items?.map((item) => (
              <ConnectionRequestItem item={item} key={item.id} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ConnectionRequestsList;
