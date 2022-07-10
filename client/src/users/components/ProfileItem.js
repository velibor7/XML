import React, { useContext } from "react";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";
import { useParams } from 'react-router-dom' 
import "./ProfileItem.css";
import WorkExperienceItem from "./WorkExperienceItem";
import EducationItem from "./EducationItem";
import Button from "../../shared/components/FormElements/Button";
import { useNavigate } from "react-router-dom";

const ProfileItem = (props) => {
  const id = useParams()['userId']
  const privacy = props.item.profile?.isPrivate
  const { sendRequest } = useHttpClient();
  const navigate = useNavigate();
  const auth = useContext(AuthContext);
  
  const UpdateProfile = async () => {
    try {
      navigate(`/profiles/${auth.userId}/update`);
    } catch (err) {
      navigate(`/profiles/${auth.userId}/update`);
      console.log(err);
    }
  };

  const SendConnectionRequest = async () => {
    try {
        var body = {
          IssuerId: auth.userId,
          subjectId: id,
          approved: !props.item.profile?.isPrivate,
          date: new Date(),  
        };
  console.log(String(new Date()));

    await sendRequest(
      "http://localhost:8000/connection",
      "POST",
      JSON.stringify(body),
      {
          "Content-Type": "application/json",
          Authorization: "token " + auth.token,
      }
    );
    console.log(JSON.stringify(body));
  } catch (err) {
    console.log(JSON.stringify(body));
    console.log(err);
  }
};

  const AddPost = async () => {
    try {
      navigate(`/posts/${auth.userId}/new`);
    } catch (err) {
      navigate(`/posts/${auth.userId}/new`);
      console.log(err);
    }
  };

  const ViewPosts = async () => {
    try {
      navigate(`/posts/${auth.userId}`);
    } catch (err) {
      navigate(`/posts/${auth.userId}`);
      console.log(err);
    }
  };

  return (
    <>
      <h1>Profile</h1>
      <div className="profile__item">
        <h4 className="profile__item__firstName">{props.item.profile?.username}</h4>
        <div className="profile__item__info">
          <p className="profile__item__firstName">
            First name: {props.item.profile?.firstName}
          </p>
          <p className="profile__item__lastName">
            Last name: {props.item.profile?.lastName}
          </p>
          <p className="profile__item__dateOfBirth">
            Date of birth: {props.item.profile?.dateOfBirth}
          </p>
          <p className="profile__item__gender">
            Gender: {props.item.profile?.gender}
          </p>
          <p className="profile__item__biography">
            Biography: {props.item.profile?.biography}
          </p>
          <p className="profile__item__phoneNumber">
            Phone number: {props.item.profile?.phoneNumber}
          </p>
          <p className="profile__item__email">
            Email: {props.item.profile?.email}
          </p>
          <div className="skills">
            <h4>Skills</h4>
            <div className="profile__item__skill">
              <div className="skill__wrapper">
                <div className="skill-list__items">
                  {props.item.profile?.skills.map((item) => (
                    <p> {item} </p>
                  ))}
                </div>
              </div>
            </div>
          </div>
          <div className="interests">
            <h4>Interests</h4>
            <div className="profile__item__interest">
              <div className="interest__wrapper">
                <div className="interest-list__items">
                  {props.item.profile?.interests.map((item) => (
                    <p> {item} </p>
                  ))}
                </div>
              </div>
            </div>
          </div>
          <div className="work_experience">
            <h4>Work experience</h4>
            <div className="profile__item__work_experience">
              <div className="work_experience__wrapper">
                <div className="work_experience-list__items">
                  {props.item.profile?.workExperience.map((item) => (
                    <WorkExperienceItem item={item} key={item.id} />
                  ))}
                </div>
              </div>
            </div>
          </div>
          <div className="education">
            <h4>Education</h4>
            <div className="profile__item__educations">
              <div className="profile__wrapper">
                <div className="education-list__items">
                  {props.item.profile?.education.map((item) => (
                    <EducationItem item={item} key={item.id} />
                  ))}
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="cocktail-item__actions">
          {
            (auth.userId == id) &&
            (<Button info onClick={UpdateProfile}>
              UPDATE
            </Button>)
          }
          {
            (auth.userId == id) &&
            (<Button info onClick={AddPost}>
              NEW POST
            </Button>
            )
          }
          {
            (auth.userId != id) && (auth.isLoggedIn) && 
            (<Button info onClick={SendConnectionRequest}>
              CONNECT
            </Button>
            )
          }
          <p></p>
            <Button info onClick={ViewPosts}>
              POSTS
            </Button>
        </div>
      </div>
    </>
  );
};

export default ProfileItem;
