import React, { useContext } from "react";
import { NavLink } from "react-router-dom";

import { AuthContext } from "../../context/auth-context";
import "./NavLinks.css";

const NavLinks = (props) => {
  const auth = useContext(AuthContext);

  return (
    <ul className="nav-links">
      <li>
        <NavLink to="/" exact>
          HOME
        </NavLink>
      </li>
      <li>
          <NavLink to="/profiles">PROFILES</NavLink>
      </li>
      {auth.isLoggedIn && (
      <li>
          <NavLink to="/jobs">JOBS</NavLink>
      </li>
      )}
      {auth.isLoggedIn && (
        <li>
          <NavLink to={"/jobs/" + auth.userId + "/recommended"} exact>RECOMMENDED JOBS</NavLink>
        </li>
      )}   
      {!auth.isLoggedIn && (
        <li>
          <NavLink to="/auth">AUTH</NavLink>
        </li>
      )}   
      {auth.isLoggedIn && (
        <li>
          <NavLink to={"/profiles/" + auth.userId}>MY PROFILE</NavLink>
        </li>
      )}
      {auth.isLoggedIn && (
        <li>
          <NavLink to={"/connections/" + auth.userId}>MY CONNECTIONS</NavLink>
        </li>
      )}
      {auth.isLoggedIn && (
        <li>
          <NavLink to={"/requests/" + auth.userId}>CONNECTION REQUESTS</NavLink>
        </li>
      )}
      {auth.isLoggedIn && (
        <li>
          <button onClick={auth.logout}>LOGOUT</button>
        </li>
      )}
    </ul>
  );
};

export default NavLinks;
