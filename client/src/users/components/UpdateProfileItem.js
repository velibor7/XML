import React, { useState, useContext } from "react";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";
import { useParams } from 'react-router-dom'  
import { Link } from "react-router-dom";
import "./ProfileItem.css";
import WorkExperienceItem from "./WorkExperienceItem"
import EducationItem from "./EducationItem"
import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import Card from "../../shared/components/UIElements/Card";
import { useNavigate } from "react-router-dom";
import { useForm } from "../../shared/hooks/form-hook";

import {
  VALIDATOR_REQUIRE,
  VALIDATOR_MINLENGTH,
} from "../../shared/util/validators";

const UpdateProfileItem = (props) => {
  const { isLoading, error, sendRequest } = useHttpClient();
  const navigate = useNavigate();
  const auth = useContext(AuthContext);
  const [formState, inputHandler, setFormData] = useForm({});
  const [loadedUser, setLoadedUser] = useState({});
  //console.log(auth)
  const UpdateSubmitHandler = async (event) => {
    event.preventDefault()
    
    try {
      var body = {
        id: auth.userId,
        username: formState.inputs.uname.value,
        firstName: formState.inputs.fname.value,
        lastName: formState.inputs.lname.value,
        phoneNumber: formState.inputs.phone.value,
        email: formState.inputs.email.value,
        dateOfBirth: formState.inputs.dateOfBirth.value + ":00Z",
        gender: formState.inputs.gender.value,
        biography: formState.inputs.biography.value,
        isPrivate: formState.inputs.biography.value == "true",
      };
      console.log(auth.userId)
      await sendRequest(
        `http://localhost:8000/profiles/${auth.userId}`,
        "PUT",
        JSON.stringify(body),
        {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
        }
      );
      navigate("/profiles");
      console.log(JSON.stringify(body));
    } catch (err) {
      console.log(JSON.stringify(body));
      console.log(err);
    }
  };

  return (
    <div style={{ marginTop: "5em" }}>
      <Card>
        {!isLoading && loadedUser && (
          <form className="user-form" onSubmit={UpdateSubmitHandler}>
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
              type="datetime-local"
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
};

export default UpdateProfileItem;
