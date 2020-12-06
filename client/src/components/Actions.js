import React, {useState} from 'react';
import {Box, Button} from 'grommet';
import {User, Add, Logout} from 'grommet-icons';

export default function Actions() {
  const [show, setShow] = useState(false);
  return (
    <Box>
      <Button 
        icon={<User />} 
        onClick={() => setShow(!show)}
      />
      <Button 
        icon={<Logout />} 
        onClick={() => setShow(!show)}
      />
      <Button 
        icon={<Add />} 
        onClick={() => setShow(!show)}
      />
    </Box>
  );
}
