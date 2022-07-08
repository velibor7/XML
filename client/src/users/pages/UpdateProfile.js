import React, { useState, useEffect, useContext } from 'react'
import { useParams } from 'react-router-dom'  
import { useNavigate } from "react-router-dom";
import { useHttpClient } from "../../shared/hooks/http-hook";
import User from '../components/ProfileItem';
import { useForm } from "../../shared/hooks/form-hook";
import { AuthContext } from "../../shared/context/auth-context";
import {
  VALIDATOR_REQUIRE,
  VALIDATOR_MINLENGTH,
} from "../../shared/util/validators";

import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import Card from "../../shared/components/UIElements/Card";
import Spinner from "../../shared/components/UIElements/Spinner";
import UpdateProfileItem from '../components/UpdateProfileItem';



const UpdateProfile = () => {
  const auth = useContext(AuthContext);
  const { isLoading, error, sendRequest} = useHttpClient();
  const navigate = useNavigate();
  const userId =  auth.id;
  const [formState, inputHandler, setFormData] = useForm({});

  const [loadedUser, setLoadedUser] = useState({});

  
  useEffect(() => {
    const fetchProfile = async () => {
        const resData = await sendRequest(
          `http://localhost:8000/profiles?username=${auth.userId}`,"GET",
        );
      console.log(resData["profile"][0])

      
      setLoadedUser(resData["profile"]);
      console.log("loaded user:")
      console.log(loadedUser)
      setFormData();
    };
    fetchProfile();
  }, [userId, sendRequest, setFormData]);


  if (isLoading) {
    return (
      <div className="center">
        <Spinner />
      </div>
    );
  }

  if (!loadedUser && !error) {
    return (
      <div className="center">
        <Card>
          <h2>Could not find a profile</h2>
        </Card>
      </div>
    );
  }

  return (
    <UpdateProfileItem item={loadedUser}  key={userId}></UpdateProfileItem>
  );
}

export default UpdateProfile;