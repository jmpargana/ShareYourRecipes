import React from 'react';
import {useAuth0} from './helpers/react-auth0-spa';
import Home from "./components/Home";
import LoggedIn from "./components/LoggedIn";

function App() {
  const { isAuthenticated, loading } = useAuth0();

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="App">
      {!isAuthenticated && (
        <Home />
      )}

      {isAuthenticated && <LoggedIn />}
    </div>
  );
}

export default App;
