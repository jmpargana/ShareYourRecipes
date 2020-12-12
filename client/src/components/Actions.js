import React, {useState} from 'react';
import {Box} from 'grommet';
import {Fab} from '@material-ui/core';
import {Person, Add, Help} from '@material-ui/icons';
import CreateRecipe from './CreateRecipe';
import HelpModal from './HelpModal';
import useStyles from './styles';
import { useAuth0 } from '../helpers/react-auth0-spa'


export default function Actions() {
  const classes = useStyles()
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();
  const [showAdd, setShowAdd] = useState(false);
  const [showHelp, setShowHelp] = useState(false);

  const handleUser = () => isAuthenticated
      ? logout()
      : loginWithRedirect()

  const fabs = [
    {
      color: 'secondary',
      label: 'user',
      icon: <Person onClick={handleUser} />,
      style: classes.fab_user,
    },
    {
      color: 'secondary',
      label: 'add',
      icon: <Add onClick={() => setShowAdd(true)} />,
      style: classes.fab_add,
    },
    {
      color: 'primary',
      label: 'help',
      icon: <Help onClick={() => setShowHelp(true)} />,
      style: classes.fab_help,
    },
  ]

  return (
    <Box>
      {fabs.map((fab, index) => (
        <Fab 
          key={index} 
          color={fab.color} 
          aria-label={fab.label} 
          className={fab.style}
        >
          {fab.icon}
        </Fab>
      ))}
      {showAdd && (<CreateRecipe toggler={setShowAdd} />)}
      {showHelp && (<HelpModal toggler={setShowHelp} />)}
    </Box>
  );
}
