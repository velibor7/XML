import React, { useState, useEffect, useContext} from "react";
import { useNavigate, useParams } from 'react-router-dom';

import ConnectionList from "../components/ConnectionList";
import Button from "../../shared/components/FormElements/Button";
import { AuthContext } from "../../shared/context/auth-context";


const AllConnections = () => {
  const [loadedConnections, setLoadedConnections] = useState();
  const navigate = useNavigate();
  const auth = useContext(AuthContext);
  const {userId} =  useParams();

  useEffect(() => {
    const fetchConnections = async () => {

      try {
        fetch(
          `http://localhost:8000/connections/${userId}`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            console.log(data)
            var connections = [];
            data = data['connection']
            for (var i = 0; i < data.length; i++) {
              var connection = data[i];
              if (connection.approved == true) {
                connections.push(connection);
              }
            }
            setLoadedConnections(connections);
          });
      } catch (err) {
        console.log("error happend")
        console.log(err);
      };
    };
    fetchConnections();
  }, []);

  return (
    <>
      <ConnectionList items={loadedConnections}></ConnectionList>
      {auth.isLoggedIn && (
      <div className="connection-item__actions">
      <Button info /*onClick={CreateNewJob}*/>
        DELETE
      </Button>
      </div>
      )} 
    </>
  );
};

export default AllConnections;