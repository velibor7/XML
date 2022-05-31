import React from "react";
import {
  Routes,
  BrowserRouter as Router,
  Route,
  Navigate,
} from "react-router-dom";

import NewCocktail from "./cocktails/pages/NewCocktail";
import UserCocktails from "./cocktails/pages/UserCocktails";
import UpdateCocktail from "./cocktails/pages/UpdateCocktail";
import MainNavigation from "./shared/components/Navigation/MainNavigation";
import Home from "./cocktails/pages/Home";
import Auth from './users/pages/Auth'
import { AuthContext } from "./shared/context/auth-context";
import { useAuth } from "./shared/hooks/auth-hook";

import "./App.css";

const App = () => {
  const { token, login, logout, userId } = useAuth();

  let routes;

  if (token) {
    routes = (
      <Routes>
        <Route path="/" exact>
          <Home />
        </Route>
        <Route path="/:userId/cocktails" exact>
          <UserCocktails />
        </Route>
        <Route path="/cocktails/new" exact>
          <NewCocktail />
        </Route>
        <Route path="/cocktails/:cocktailId">
          <UpdateCocktail />
        </Route>
        <Navigate to="/" />
      </Routes>
    );
  } else {
    routes = (
      <Routes>
        <Route path="/" exact>
          <Home />
        </Route>
        <Route path="/:userId/cocktails" exact>
          <UserCocktails />
        </Route>
        <Route path="/auth">
          <Auth />
        </Route>
        <Navigate to="/auth" />
      </Routes>
    );
  }

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
        <main>{routes}</main>
      </Router>
    </AuthContext.Provider>
  );
};

export default App;
