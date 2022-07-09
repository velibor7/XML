import React, { useContext } from "react";
import { useNavigate } from "react-router-dom";

import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import ErrorModal from "../../shared/components/UIElements/ErrorModal";
import Spinner from "../../shared/components/UIElements/Spinner";
import Card from "../../shared/components/UIElements/Card";
import {
  VALIDATOR_REQUIRE,
  VALIDATOR_MINLENGTH,
} from "../../shared/util/validators";
import { useForm } from "../../shared/hooks/form-hook";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";

import "./NewJobForm.css";



const NewJob = () => {
  const auth = useContext(AuthContext);
  const { isLoading, error, sendRequest, clearError } = useHttpClient();
  const navigate = useNavigate();
  const [formState, inputHandler] = useForm(
    {
      title: {
        value: "",
        isValid: false,
      },
      description: {
        value: "",
        isValid: false,
      },
      skills: {
        value: null,
        isValid: false,
      },
    },
    false
  );

  const jobSubmitHandler = async (event) => {
    event.preventDefault();

    try {
        var body = {
          title: formState.inputs.title.value,
          description: formState.inputs.description.value,
          skills: formState.inputs.skills.value.split(" "),
          userId: auth.userId    
        };

      await sendRequest(
        "http://localhost:8000/jobs",
        "POST",
        JSON.stringify(body),
        {
            "Content-Type": "application/json",
            Authorization: "token " + auth.token,
        }
      );
      navigate("/jobs");
      console.log(JSON.stringify(body));
    } catch (err) {
      console.log(JSON.stringify(body));
      console.log(err);
    }
  };

  return (
    <React.Fragment>
      <ErrorModal error={error} onClear={clearError} />
      <Card className="job-card">
        <form className="job-form" onSubmit={jobSubmitHandler}>
          {isLoading && <Spinner asOverlay />}
          <Input
            id="title"
            element="input"
            type="text"
            label="Title"
            validators={[VALIDATOR_REQUIRE()]}
            errorText="Please enter valid title"
            onInput={inputHandler}
          />
          <Input
            id="description"
            element="textarea"
            label="Description"
            validators={[VALIDATOR_MINLENGTH(5)]}
            errorText="Please enter valid description"
            onInput={inputHandler}
          />
          <Input
            id="skills"
            element="textarea"
            label="Skills"
            validators={[VALIDATOR_MINLENGTH(5)]}
            errorText="Please enter valid skills"
            onInput={inputHandler}
          />
          <Button type="submit" disabled={!formState.isValid}>
            SUBMIT
          </Button>
        </form>
      </Card>
    </React.Fragment>
  );
};

export default NewJob;