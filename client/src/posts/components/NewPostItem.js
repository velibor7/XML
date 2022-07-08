import React, { useState, useContext } from "react";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";
import { useParams } from 'react-router-dom'  
import "./PostItem.css";
import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import Card from "../../shared/components/UIElements/Card";
import { useNavigate } from "react-router-dom";
import { useForm } from "../../shared/hooks/form-hook";

import {
  VALIDATOR_REQUIRE,
  VALIDATOR_MINLENGTH,
} from "../../shared/util/validators";

const NewPostItem = (props) => {
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
        user_id: auth.userId,
        text: formState.inputs.text.value,
        links: formState.inputs.links.value,
      };
      await sendRequest(
        `http://localhost:8000/posts`,
        "POST",
        JSON.stringify(body),
        {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
        }
      );
      navigate("/profiles/" + auth.userId);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div style={{ marginTop: "5em" }}>
      <Card>
        {!isLoading && loadedUser && (
          <form className="user-form" onSubmit={UpdateSubmitHandler}>
            <Input
              id="links"
              element="input"
              type="text"
              label="Links"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid first name"
              onInput={inputHandler}
              initialValue={loadedUser.firstName}
              initialValid={true}
            />
            <Input
              id="text"
              element="input"
              type="text"
              label="Text"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid last name"
              onInput={inputHandler}
              initialValue={loadedUser.lastName}
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

export default NewPostItem;
