import { Link } from "react-router-dom";
import "./ConnectionItem.css";
import Button from "../../shared/components/FormElements/Button";
import { useParams } from 'react-router-dom' 
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";

import React, { useContext } from "react";

const ConnectionRequestItem = (props) => {
  const id = useParams()['userId']
  const { sendRequest } = useHttpClient();
  const auth = useContext(AuthContext);

  const DeclineConnection = async () => {
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

  const AcceptConnection = async () => {
    console.log(props);
    try {
      var body = {
        issuerId: props.item.issuerId,
        subjectId: props.item.subjectId,
        approved: true, 
        date: props.item.date, 

      };

    await sendRequest(
    "http://localhost:8000/connection",
    "POST",
    JSON.stringify(body),
    {
        "Content-Type": "application/json",
        Authorization: "token " + auth.token,
    }
  );

  console.log(JSON.stringify(body));
} catch (err) {
  console.log(err);
}
};

  return (
    <> { (props.item.approved == false) &&
      (
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
            (<Button onClick={AcceptConnection}>
              ACCEPT
            </Button>
            )
        }
                {
            (<Button danger onClick={DeclineConnection}>
              DECLINE
            </Button>
            )
        }
        </div>
      </div>
)}
    </>
  );
};

export default ConnectionRequestItem;