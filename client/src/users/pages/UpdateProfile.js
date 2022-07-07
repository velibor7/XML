import React, { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'  
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




const UpdateProfile = () => {
  const auth = useContext(AuthContext);
  const { isLoading, error, sendRequest} = useHttpClient();
  const navigate = useNavigate();
  const userId =  auth.userId;
  const [formState, inputHandler, setFormData] = useForm({});

  const [loadedUser, setLoadedUser] = useState({});


  
  useEffect(() => {
    const fetchProfile = async () => {
        const resData = await sendRequest(
          `http://localhost:8000/profiles/${userId}`,"GET",
        );
      setLoadedUser(resData);
      setFormData();

      setLoadedUser(resData)
    };
    fetchProfile();
  }, [userId, sendRequest, setFormData]);

  const userUpdateSubmitHandler = async (event) => {
    event.preventDefault()
    
    try {
      var body = {
        username: formState.inputs.uname.value,
        firstName: formState.inputs.fname.value,
        lastName: formState.inputs.lname.value,
        phoneNumber: formState.inputs.phone.value,
        email: formState.inputs.email.value,
        dateOfBirth: formState.inputs.dateOfBirth.value,
        gender: formState.inputs.gender.value,
        biography: formState.inputs.biography.value,
        isPrivate: formState.inputs.biography.value,
      };
      await sendRequest(
        `http://localhost:8000/profiles/${userId}`,
        "PATCH",
        JSON.stringify(body),
        {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
        }
      );
      navigate("/");
    } catch (err) {console.log(err)}
  };

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
    <div style={{ marginTop: "5em" }}>
      <Card>
        {!isLoading && loadedUser && (
          <form className="user-form" onSubmit={userUpdateSubmitHandler}>
            <Input
              id="uname"
              element="input"
              type="text"
              label="Username"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid username"
              onInput={inputHandler}
              initialValue={loadedUser.username}
              initialValid={true}
            />
            <Input
              id="fname"
              element="input"
              type="text"
              label="First name"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid first name"
              onInput={inputHandler}
              initialValue={loadedUser.firstName}
              initialValid={true}
            />
            <Input
              id="lname"
              element="input"
              type="text"
              label="Last name"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid last name"
              onInput={inputHandler}
              initialValue={loadedUser.lastName}
              initialValid={true}
            />
            <Input
              id="phone"
              element="input"
              type="number"
              label="Phone number"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid phone number"
              onInput={inputHandler}
              initialValue={loadedUser.phone_number}
              initialValid={true}
            />
            <Input
              id="email"
              element="input"
              type="text"
              label="Email"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid email"
              onInput={inputHandler}
              initialValue={loadedUser.email}
              initialValid={true}
            />
            <Input
              id="dateOfBirth"
              element="input"
              type="text"
              label="Date Of Birth"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid date"
              onInput={inputHandler}
              initialValue={loadedUser.dateOfBirth}
              initialValid={true}
            />
            <Input
              id="gender"
              element="input"
              type="text"
              label="Gender"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid gender"
              onInput={inputHandler}
              initialValue={loadedUser.gender}
              initialValid={true}
            />
            <Input
              id="biography"
              element="input"
              type="text"
              label="Biography"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid biography"
              onInput={inputHandler}
              initialValue={loadedUser.biography}
              initialValid={true}
            />
            <Input
              id="isPrivate"
              element="input"
              type="text"
              label="Is Private"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid is private value"
              onInput={inputHandler}
              initialValue={loadedUser.isPrivate}
              initialValid={true}
            />
            <Button
              type="submit"
              // disabled={!formState.isValid}
            >
              SUBMIT
            </Button>
          </form>
        )}
      </Card>
    </div>
  );
}

export default UserProfile;