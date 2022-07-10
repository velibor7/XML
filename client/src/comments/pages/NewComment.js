import React, { useState, useEffect, useContext } from 'react'
import { useParams } from 'react-router-dom'  
import { useNavigate } from "react-router-dom";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { useForm } from "../../shared/hooks/form-hook";
import { AuthContext } from "../../shared/context/auth-context";

import Input from "../../shared/components/FormElements/Input";
import Button from "../../shared/components/FormElements/Button";
import Card from "../../shared/components/UIElements/Card";
import Spinner from "../../shared/components/UIElements/Spinner";
import NewCommentItem from '../components/NewCommentItem';



const NewComment = () => {
  const auth = useContext(AuthContext);
  const { isLoading, error, sendRequest} = useHttpClient();
  const navigate = useNavigate();
  const userId =  auth.id;
  const [formState, inputHandler, setFormData] = useForm({});

  const [loadedUser, setLoadedUser] = useState({});


  return (
    <NewCommentItem></NewCommentItem>
  );
}

export default NewComment;