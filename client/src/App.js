import React from "react";
import {Box,Grommet} from "grommet";
import Home from './components/Home'
import Actions from './components/Actions'
import theme from './assets/theme';

export default function App() {
  return (
    <Grommet theme={theme} full>
      <Box fill>
        <Box direction='row' flex overflow={{ horizontal: 'hidden' }}>
          <Box flex align='center' justify='center'>
            <Home />
          </Box>
          <Box
            width='medium'
            background='light-2'
            elevation='small'
            align='center'
            justify='center'
          >
            <Actions />
          </Box>
        </Box>
      </Box>
    </Grommet>
  );
}
