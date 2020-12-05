import React, {useState} from 'react';
import {TextInput} from 'grommet';

export default function SearchBar() {
  const [value, setValue] = useState('')
  return (
    <TextInput
      placeholder='Search here'
      value={value}
      onChange={event => setValue(event.target.value)}
    />
  );
}
