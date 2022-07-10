import CommentItem from "./CommentItemForList";
import React, { useContext } from "react";
import { AuthContext } from "../../shared/context/auth-context";
import { useParams } from 'react-router-dom' 
import Button from "../../shared/components/FormElements/Button";
import { useNavigate } from "react-router-dom";

const CommentList = (props) => {

    const navigate = useNavigate();
    const auth = useContext(AuthContext);
    const id = useParams()["id"]

    const newCommentNavigate = async () => {
        navigate("/comments/" + id + "/new")
    }

  return (
    <div className="comments">
      <h1>List of Comments</h1>
      <div className="comment-list__container">
        <div className="comment__wrapper">
          <div className="comment-list__items">
            {props.items?.map((item) => (
              <CommentItem item={item} key={item.id} />
            ))}
          </div>
          <div>
          </div>
          {
              (auth.isLoggedIn) &&
              (<Button success onClick={newCommentNavigate}>
              New comment
            </Button>
            )
          }
        </div>
      </div>
    </div>
  );
};

export default CommentList;