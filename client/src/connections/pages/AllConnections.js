import React, { useState, useEffect, _, useContext } from "react";
import ConnectionList from "../components/ConnectionList";
import { useParams } from 'react-router-dom'
import { AuthContext } from "../../shared/context/auth-context"; 

const AllConnections = () => {
  const [loadedConnections, setLoadedConnections] = useState();
  var id = useParams()['id']
  const auth = useContext(AuthContext);

  useEffect(() => {
    const fetchConnections = async () => {
      try {
        fetch(
          `http://localhost:8000/connection/${auth.userId}`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            console.log(data["connections"])
            setLoadedConnections(data["connections"]);
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
        <h1>Connections</h1>
        <ConnectionList items={loadedConnections}></ConnectionList>
    </>
  );
};

export default AllConnections;