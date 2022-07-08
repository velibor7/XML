import React from "react";
import {
  Routes,
  BrowserRouter as Router,
  Route,
} from "react-router-dom";

import UserProfile from "./users/pages/Profile"
import UserProfiles from "./users/pages/AllProfiles"
import UserPosts from "./posts/pages/AllPosts"

import MainNavigation from "./shared/components/Navigation/MainNavigation";
import NotFound from './shared/components/UIElements/NotFound'
import Auth from './users/pages/Auth'
import { AuthContext } from "./shared/context/auth-context";
import { useAuth } from "./shared/hooks/auth-hook";

import "./App.css";
import UpdateProfile from "./users/pages/UpdateProfile";

const App = () => {
  const { token, username, login, logout, userId } = useAuth();

  return(
    <AuthContext.Provider
      value={{
        isLoggedIn: !!token,
        token: token,
        userId: userId,
        username: username,
        login: login,
        logout: logout,
      }}
    >
      <Router>
        <MainNavigation />
        <Routes>
          <Route path="/profiles/:userId/" element={<UserProfile/>}/>
          <Route path="/profiles/:userId/update" element={<UpdateProfile/>}/>
          <Route path="/profiles" element={<UserProfiles/>}/>
          <Route path="posts/:id/" element={<UserPosts/>}/>
          <Route path="/auth" element={<Auth/>}/>
          <Route path="*" element={<NotFound />} />
          {(token === true) && (
            <>
              <Route path="*" element={<NotFound />} />
            </>
          )
          }
        </Routes>
      </Router>
    </AuthContext.Provider>
  )
};

export default App;
