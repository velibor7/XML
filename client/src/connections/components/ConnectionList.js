import ConnectionItem from "./ConnectionItem";
import React from "react";

const ConnectionList = (props) => {
  return (
    <div className="connections">
      <h1>My Connections</h1>
      <div className="connection-list__container">
        <div className="connection__wrapper">
          <div className="connection-list__items">
            {props.items?.map((item) => (
              <ConnectionItem item={item} key={item.id} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default ConnectionList;