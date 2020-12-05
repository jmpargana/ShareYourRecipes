import React from "react";
import {useAuth0} from "../helpers/react-auth0-spa";

export default function LoggedIn() {
  const { loading, user } = useAuth0();

  if (loading || !user) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Logged in</h1>
    </div>
  )
}
