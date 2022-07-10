
import React from "react";
import {
  Routes,
  BrowserRouter as Router,
  Route,
} from "react-router-dom";

import UserProfile from "./users/pages/Profile"
import UserProfiles from "./users/pages/AllProfiles"

import UserPosts from "./posts/pages/AllPosts"
import Post from "./posts/pages/Post"
import NewPost from "./posts/pages/NewPost"

import PostComments from "./comments/pages/PostComments"
import NewComment from "./comments/pages/NewComment"

import AllJobs from "./jobs/pages/AllJobs"
import RecommendedJobs from "./jobs/pages/RecommendedJobs"
import NewJob from "./jobs/pages/NewJob"


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
          <Route path="/jobs" element={<AllJobs/>}/>
          <Route path="/newjob" element={<NewJob/>}/>
          <Route path="/jobs/:id/recommended" element={<RecommendedJobs/>}/>
          <Route path="/profiles" element={<UserProfiles/>}/>
          <Route path="posts/:id/" element={<UserPosts/>}/>
          <Route path="posts/:userId/:postId" element={<Post/>}/>
          <Route path="posts/:id/new" element={<NewPost/>}/>
          <Route path="comments/:id" element={<PostComments/>}/>
          <Route path="comments/:id/new" element={<NewComment/>}/>
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