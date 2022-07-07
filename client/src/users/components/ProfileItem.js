import { Link } from "react-router-dom";
import "./ProfileItem.css";
import WorkExperienceItem from "./WorkExperienceItem"
import EducationItem from "./EducationItem"
import Button from "../../shared/components/FormElements/Button";

import React, { useContext } from "react";


const ProfileItem = (props) => {
  

  return (
    <>
      <h1>Profile</h1>
      <div className="profile__item">
        <h4 className="profile__item__firstName">{props.item.profile?.firstName}</h4>
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
                  {props.items?.map((item) => (
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
                  {props.items?.map((item) => (
                    <EducationItem item={item} key={item.id} />
                  ))}
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="cocktail-item__actions">
          <Link to={`/profiles/${props.item.id}`}>
            <Button>
              UPDATE
            </Button>
          </Link>

        </div>
      </div>
    </>
  );
};

export default ProfileItem;
