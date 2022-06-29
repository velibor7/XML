import React, { useContext, useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";

import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import Card from "../../shared/components/UIElements/Card";
import Spinner from "../../shared/components/UIElements/Spinner";

import {
  VALIDATOR_REQUIRE,
  VALIDATOR_MINLENGTH,
} from "../../shared/util/validators";
import { useForm } from "../../shared/hooks/form-hook";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";


const UpdateUser = () => {
  const auth = useContext(AuthContext);
  const { isLoading, error, sendRequest} = useHttpClient();
  const [loadedUser, setLoadedUser] = useState();
  const navigate = useNavigate();
  const userId =  auth.userId;
  const [formState, inputHandler, setFormData] = useForm({});

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const resData = await sendRequest(
          `http://localhost:8001/api/authenticatedUser/${userId}`
        );
        setLoadedUser(resData);
        setFormData();
      } catch (err) {}
    };
    fetchUser();
  }, [sendRequest, userId, setFormData]);

  const userUpdateSubmitHandler = async (event) => {
    event.preventDefault()
    
    try {
      var body = {
        first_name: formState.inputs.fname.value,
        last_name: formState.inputs.lname.value,
        phone_number: formState.inputs.phone.value,
        country: formState.inputs.country.value,
        city: formState.inputs.city.value,
        street: formState.inputs.street.value,
        //password : formState.inputs.password.value
      };
      await sendRequest(
        `http://localhost:8001/api/authenticatedUser/${userId}`,
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
          <h2>Could not find a cottage</h2>
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
              id="fname"
              element="input"
              type="text"
              label="First name"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid first name"
              onInput={inputHandler}
              initialValue={loadedUser.first_name}
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
              initialValue={loadedUser.last_name}
              initialValid={true}
            />
            <Input
              id="city"
              element="input"
              type="text"
              label="City"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid city"
              onInput={inputHandler}
              initialValue={loadedUser.city}
              initialValid={true}
            />
            <Input
              id="country"
              element="input"
              type="text"
              label="Country"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid country"
              onInput={inputHandler}
              initialValue={loadedUser.country}
              initialValid={true}
            />
            <Input
              id="street"
              element="input"
              type="text"
              label="Street"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid street name"
              onInput={inputHandler}
              initialValue={loadedUser.street}
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
            {/* <Input
              id="password"
              element="input"
              type="password"
              label="Password"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid phone number"
              onInput={inputHandler}
              initialValid={true}
            /> */}
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

export default UpdateUser;
