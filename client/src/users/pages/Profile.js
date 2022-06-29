import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'  
import { useHttpClient } from "../../shared/hooks/http-hook";
import User from '../components/ProfileItem';

const UserProfile = () => {

  const {userId} = useParams();
  const [loadedUser, setLoadedUser] = useState({});

  const { sendRequest} = useHttpClient();

  
  useEffect(() => {
    const fetchCottage = async () => {
        const resData = await sendRequest(
          `http://localhost:8000/profiles/${userId}`
        )
        console.log(resData)

      setLoadedUser(resData)
    };
    fetchCottage();
  }, [userId, sendRequest]);

  return (
    <>
      <User item={loadedUser}></User>
    </>
  );
}

export default UserProfile;