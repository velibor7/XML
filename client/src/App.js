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

import "./App.css";
import HomeUnauth from "./home/pages/HomeUnauth";
import Jobs from "./jobs/pages/AllJobs";
import Profiles from "./profiles/pages/Profiles";
import Auth from "./users/pages/Auth";

const App = () => {
  const { token, login, logout, userId } = useAuth();

  let routes;

  // if (token) {
  //   routes = (
  //     <Routes>
  //       <Route path="/" exact>
  //         {/* <Home /> */}
  //         <div>home</div>
  //       </Route>
  //       {/*
  //       <Route path="/:userId/cocktails" exact>
  //         <UserCocktails />
  //       </Route>
  //       <Route path="/cocktails/new" exact>
  //         <NewCocktail />
  //       </Route>
  //       <Route path="/cocktails/:cocktailId">
  //         <UpdateCocktail />
  //       </Route> */}
  //       <Navigate to="/" />
  //     </Routes>
  //   );
  // } else {
  //   routes = (
  //     <Routes>
  //       <Route path="/auth" component={ Auth } />
  //       {/* <Navigate to="/auth" /> */}
  //   );
  // }

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
          <Route path="/" exact="true" element={<HomeUnauth />} />
          <Route path="/auth" exact="true" element={<Auth />} />
          <Route path="/jobs" exact="true" element={<Jobs />} />
          <Route path="/profiles" exact="true" element={<Profiles />} />
        </Routes>
      </Router>
    </AuthContext.Provider>
  );
};

export default App;