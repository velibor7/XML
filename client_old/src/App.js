import React from "react";
import {
  Routes,
  BrowserRouter as Router,
  Route,
  Navigate,
} from "react-router-dom";
import MainNavigation from "./shared/components/Navigation/MainNavigation";
import { AuthContext } from "./shared/context/auth-context";
import { useAuth } from "./shared/hooks/auth-hook";
import Users from './users/pages/Users'

import "./App.css";

const App = () => {
  const { token, login, logout, userId } = useAuth();

  let routes;

  return (
    <div>
      somethin
    </div>
  )

  return (
    <AuthContext.Provider
      value={{
        isLoggedIn: !!token,
        token: token,
        userId: userId,
        login: login,
        logout: logout,
      }}
    >
      <Router>
        <MainNavigation />
        <Routes>
          <Route path="/" exact="true" element={<Users />} />
        </Routes>
      </Router>
    </AuthContext.Provider>
  );
};

export default App;