import React, {useState} from 'react';
import {Box, Button} from 'grommet';
import {User, Add, Logout} from 'grommet-icons';

export default function Actions() {
  const [show, setShow] = useState(false);
  return (
    <Box>
      <User /> 
      <Logout />
      <Button 
        icon={<Add />} 
        onClick={() => setShow(!show)}
      />
      <Add />
    </Box>
  );
}
