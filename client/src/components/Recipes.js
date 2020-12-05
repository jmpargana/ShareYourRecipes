import React from 'react';
import {List} from 'grommet';

export default function Recipes() {
  return (
    <List
      primaryKey="name"
      data={[
        {name: "Olah"},
        {name: "Olah"},
        {name: "Olah"},
        {name: "Olah"},
      ]}
    />
  );
}
