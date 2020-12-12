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

  return (
    <Box>
      <Fab 
        color='secondary'
        aria-label='user'
        className={classes.fab_user}
      >
        <Person onClick={handleUser} />
      </Fab>
      <Fab 
        color='primary'
        aria-label='help'
        className={classes.fab_help}
      >
        <Help onClick={() => setShowHelp(true)} />
      </Fab>
      {isAuthenticated ? (
        <Fab 
          color='secondary'
          aria-label='add'
          className={classes.fab_add}
        >
          <Add onClick={() => setShowAdd(true)} />
        </Fab>
      ) : ''}
      {showAdd && (<CreateRecipe toggler={setShowAdd} />)}
      {showHelp && (<HelpModal toggler={setShowHelp} />)}
    </Box>
  );
}
