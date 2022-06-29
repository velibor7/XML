import React, { useContext} from 'react'
import { useNavigate, useParams } from 'react-router-dom';
import Button from '../../shared/components/FormElements/Button';
import Input from '../../shared/components/FormElements/Input';
import Card from '../../shared/components/UIElements/Card';
import Spinner from '../../shared/components/UIElements/Spinner';
import { AuthContext } from '../../shared/context/auth-context';
import { useForm } from '../../shared/hooks/form-hook';
import { useHttpClient } from '../../shared/hooks/http-hook';
import { VALIDATOR_REQUIRE } from '../../shared/util/validators';

const DeleteUser = () => {
    const auth = useContext(AuthContext);
    const { isLoading, sendRequest} = useHttpClient();
    const [ formState, inputHandler ] = useForm();
    const {userId} = useParams();
    const navigate = useNavigate();
  

    const deleteSubmitHandler = async (event) => {
      event.preventDefault();
      const body = {
          is_processed : false,
          description: formState.inputs.desc.value,
          response : "",
          authenticated_user : userId
      };  
    console.log(body)
    try {
      await sendRequest(
        `http://localhost:8001/api/userDeleteRequest`,
        "POST",
        JSON.stringify(body),
        { 
            "Content-Type": "application/json",
            Authorization: "token " + auth.token
         }
      );
      navigate("/")
    } catch (err) {
      console.log(err);
    }
  };
  return (
    <div style={{marginTop : "5em"}}>
      <Card className="create-card">
        <form className="create-form" onSubmit={deleteSubmitHandler}>
          {isLoading && <Spinner asOverlay />}
          <Input
            id="desc"
            element="input"
            type="text"
            label="Why do you want to delete an account"
            validators={[VALIDATOR_REQUIRE()]}
            errorText="Please enter valid reason"
            onInput={inputHandler}
          />
          <Button type="submit" 
          // disabled={!formState.isValid}
          >
            Confirm
          </Button>
        </form>
      </Card>
    </div>
  )
}

export default DeleteUser;