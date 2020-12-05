import React, {Fragment} from "react";
import {useAuth0} from "../helpers/react-auth0-spa";

export default function Home() {
  const { isAuthenticated, loginWithRedirect} = useAuth0();
  return (
    <Fragment>
      <div>
        <div>
          <h1>We R VR</h1>
          <p>Provide valuable feedback</p>
          {!isAuthenticated && (
            <button onClick={() => loginWithRedirect({})}>Sign in</button>
          )}
        </div>
      </div>
    </Fragment>      
  );
}
