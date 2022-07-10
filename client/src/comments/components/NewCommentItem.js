import React, { useState, useContext } from "react";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";
import { useParams } from 'react-router-dom'  
import "./CommentItem.css";
import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import Card from "../../shared/components/UIElements/Card";
import { useNavigate } from "react-router-dom";
import { useForm } from "../../shared/hooks/form-hook";

import {
  VALIDATOR_REQUIRE,
  VALIDATOR_MINLENGTH,
} from "../../shared/util/validators";

const NewCommentItem = (props) => {
  const { isLoading, error, sendRequest } = useHttpClient();
  const navigate = useNavigate();
  const auth = useContext(AuthContext);
  const [formState, inputHandler, setFormData] = useForm({});
  const [loadedUser, setLoadedUser] = useState({});
  var id = useParams()['id']
  //console.log(auth)
  const UpdateSubmitHandler = async (event) => {
    event.preventDefault()
    console.log(props)
    try {
      var body = {
        postId: id,
        username: auth.username,
        content: formState.inputs.content.value,
      };
      await sendRequest(
        `http://localhost:8000/post/${id}/comment`,
        "POST",
        JSON.stringify(body),
        {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
        }
        );
        console.log(JSON.stringify(body));
        navigate("/comments/" + id);
    } catch (err) {
      console.log(err);
      console.log(JSON.stringify(body))
    }
  };

  return (
    <div style={{ marginTop: "5em" }}>
      <Card>
        {!isLoading && loadedUser && (
          <form className="user-form" onSubmit={UpdateSubmitHandler}>
            <Input
              id="content"
              element="input"
              type="textarea"
              label="Content"
              validators={[VALIDATOR_REQUIRE()]}
              errorText="plase enter a valid content"
              onInput={inputHandler}
              initialValue={true}
              initialValid={true}
            />
            <Button
              type="submit"
            >
              SUBMIT
            </Button>
          </form>
        )}
      </Card>
    </div>
  );
};

export default NewCommentItem;
