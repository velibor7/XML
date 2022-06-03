import React, { useState, useContext } from "react";
import { useNavigate } from 'react-router-dom'

import Card from "../../shared/components/UIElements/Card";
import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import ErrorModal from "../../shared/components/UIElements/ErrorModal";
import Spinner from "../../shared/components/UIElements/Spinner";

import {
  VALIDATOR_EMAIL,
  VALIDATOR_MINLENGTH,
  VALIDATOR_REQUIRE,
} from "../../shared/util/validators";
import { useForm } from "../../shared/hooks/form-hook";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";
import "./Auth.css";

const Auth = (props) => {
  const auth = useContext(AuthContext);
  const [ isLoginMode, setIsLoginMode ] = useState(true);
  const { isLoading, error, sendRequest, clearError } = useHttpClient();

  const navigate = useNavigate()

  const [formState, inputHandler, setFormData] = useForm(
    {
      email: {
        value: "",
        isValid: false,
      },
      password: {
        value: "",
        isValid: false,
      },
    },
    false
  );

  const switchModeHandler = () => {
    if (!isLoginMode) {
      setFormData(
        {
          ...formState.inputs,
          name: undefined,
        },
        formState.inputs.email.isValid && formState.inputs.password.isValid
      );
    } else {
      setFormData(
        {
          ...formState.inputs,
          name: {
            value: "",
            isValid: false,
          },
        },
        false
      );
    }
    setIsLoginMode((prevMode) => !prevMode);
  };

  const authSubmitHandler = async (event) => {
    event.preventDefault();

    if (isLoginMode) {
      try {
        const responseData = await sendRequest(
          "http://localhost:8001/api/authToken",
          "POST",
          JSON.stringify({
            username: formState.inputs.email.value,
            password: formState.inputs.password.value,
          }),
          {
            "Content-Type": "application/json",
          }
        );
        console.log(responseData);
        auth.login(responseData.user_id, responseData.type, responseData.token);

        navigate.push('/')
        
      } catch (err) {}
    } else {
      try {
        const formData = new FormData();
        formData.append("first_name", formState.inputs.firstName.value);
        formData.append("last_name", formState.inputs.lastName.value);
        formData.append("phone_number", formState.inputs.phoneNumber.value);
        formData.append("type", formState.inputs.userType.value);
        formData.append("country", formState.inputs.country.value);
        formData.append("city", formState.inputs.city.value);
        formData.append("street", formState.inputs.street.value);
        formData.append("email", formState.inputs.email.value);
        formData.append("password", formState.inputs.password.value);
        const responseData = await sendRequest(
          "http://localhost:8001/api/authenticatedUser",
          "POST",
          formData
        );

        // auth.login(responseData.userId, responseData.token);
      } catch (err) {}
    }
  };

  return (
    <>
      <ErrorModal error={error} onClear={clearError} />
      <Card className="auth">
        <h2>Login Required</h2>
        <hr />
        <form onSubmit={authSubmitHandler}>
          {isLoading && <Spinner asOverlay />}
          {!isLoginMode && (
            <>
            <Input
              element="input"
              id="firstName"
              type="text"
              label="First Name"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a first name."
              onInput={inputHandler}
            />
            <Input
              element="input"
              id="lastName"
              type="text"
              label="Last Name"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a last name."
              onInput={inputHandler}
            />
            <Input
              element="input"
              id="phoneNumber"
              type="text"
              label="Phone number"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a phone number."
              onInput={inputHandler}
            />
            <Input
              element="input"
              id="country"
              type="text"
              label="Country"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a country."
              onInput={inputHandler}
            />
            <Input
              element="input"
              id="city"
              type="text"
              label="City"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a city."
              onInput={inputHandler}
            />
            <Input
              element="input"
              id="street"
              type="text"
              label="Street"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a street."
              onInput={inputHandler}
            />
            <Input
              element="input"
              id="userType"
              type="text"
              label="Type: CLIENT/BOAT_OWNER/COTTAGE_OWNER/INSTRUCTOR"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="Please enter a user type."
              onInput={inputHandler}
            />
          </>
          )}
          <Input
            element="input"
            id="email"
            type="email"
            label="E-Mail"
            validators={[VALIDATOR_EMAIL()]}
            errorText="Please enter a valid email address."
            onInput={inputHandler}
          />
          <Input
            element="input"
            id="password"
            type="password"
            label="Password"
            validators={[VALIDATOR_MINLENGTH(6)]}
            errorText="Please enter a valid password, at least 6 characters."
            onInput={inputHandler}
          />
          {!isLoginMode && (
            <Input
              element="input"
              id="password"
              type="password"
              label="Password, again"
              validators={[VALIDATOR_MINLENGTH(6)]}
              errorText="Please enter a valid password, at least 6 characters."
              onInput={inputHandler}
            />
          )}
          <Button
            type="submit"
            // disabled={!formState.isValid}
            className="center-btn"
          >
            {isLoginMode ? "LOGIN" : "SIGNUP"}
          </Button>
        </form>
        <Button inverse onClick={switchModeHandler} classNameName="center-btn">
          SWITCH TO {isLoginMode ? "SIGNUP" : "LOGIN"}
        </Button>
      </Card>
    </>
  );
};

export default Auth;
