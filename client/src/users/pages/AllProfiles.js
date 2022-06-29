import React, { useState, useEffect, _ } from "react";
import ProfileList from "../components/ProfileList";

const AllProfiles = () => {
  const [loadedProfiles, setLoadedProfiles] = useState();


  useEffect(() => {
    const fetchProfiles = () => {
      try {
        fetch(
          `http://localhost:8000/profiles`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            var profiles = [];
            for (var i = 0; i < data.length; i++) {
              var profile = data[i];
              console.log(profile);
              profiles.push(profile);
            }
            setLoadedProfiles(profiles);
          });
      } catch (err) {
        console.log(err);
      };
    };
    console.log('aaaaaaaaaaa')
    fetchProfiles();
    console.log('bbbbbbbbbbbbbb')
    return
  }, []);

  return (
    <>
        <ProfileList items={loadedProfiles}></ProfileList>
    </>
  );
};

export default AllProfiles;
