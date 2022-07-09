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

  const history = useNavigate();

  const NavigateToJobs = async () => {
    try {
      navigate(`/jobs`);
    } catch (err) {
      navigate(`/jobs`);
      console.log(err);
    }
  }

  const jobSubmitHandler = async (event) => {
    event.preventDefault();
    // console.log(auth.token);

    try {
      const formData = new FormData();
      formData.append("title", formState.inputs.title.value);
      formData.append("description", formState.inputs.description.value);
      formData.append("skills", formState.inputs.skills.value);

      // console.log(formState.inputs.image.value);

      await sendRequest(
        "http://localhost:8000/jobs",
        "POST",
        formData,
        {
          Authorization: "Bearer " + auth.token,
        }
      );
      history.push("/");
    } catch (err) {}
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
          <Button type="submit" disabled={!formState.isValid} onClick={NavigateToJobs}>
            SUBMIT
          </Button>
        </form>
      </Card>
    </React.Fragment>
  );
};

export default NewJob;