import { Link } from "react-router-dom";
import "./ConnectionItem.css";
import Button from "../../shared/components/FormElements/Button";
import { useParams } from 'react-router-dom' 
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";

import React, { useContext } from "react";


const ConnectionItem = (props) => {
  const id = useParams()['userId']
  const { sendRequest } = useHttpClient();
  const auth = useContext(AuthContext);

  const DeleteConnection = async () => {
    console.log(props);
    try {
      await sendRequest(
        `http://localhost:8000/connection/${props.item.id}`,
        "DELETE",
        null,
        { "Content-Type": "application/json",
        Authorization: "token " + auth.token, }
      );
    } catch (err) {
      console.log(err);
    }
  };

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
        <div>
        {
            (<Button danger onClick={DeleteConnection}>
              DELETE
            </Button>
            )
          }
        </div>
      </div>
    </>
  );
};

export default ConnectionItem;